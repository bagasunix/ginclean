package usecases

import (
	"context"
	"time"

	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/bagasunix/ginclean/pkg/helpers"
	"github.com/bagasunix/ginclean/pkg/jwt"
	"github.com/bagasunix/ginclean/server/domains/data/models"
	"github.com/bagasunix/ginclean/server/domains/data/repositories"
	"github.com/bagasunix/ginclean/server/domains/entities"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/bagasunix/ginclean/server/endpoints/responses"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type AccountService interface {
	CreateAccount(ctx context.Context, req *requests.CreateAccount) (res *responses.EntityId, err error)
	ListAccount(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Account], err error)
	ViewAccountByID(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Account], err error)
	DisableAccount(ctx context.Context, req *requests.DisableAccount) (res *responses.Empty, err error)
	DeleteAccount(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error)
	LoginAccount(ctx context.Context, req *requests.SignInWithEmailPassword) (res *responses.SignIn, err error)
	RefreshToken(ctx context.Context, req *requests.Token) (res *responses.RefreshToken, err error)
}

type accountUseCase struct {
	jwtKey        string
	jwtKeyRefresh string
	logs          *zap.Logger
	repo          repositories.Repositories
}

// RefreshToken implements AccountService
func (a *accountUseCase) RefreshToken(ctx context.Context, req *requests.Token) (res *responses.RefreshToken, err error) {
	resBuilder := responses.NewRefreshTokenBuilder()
	if req.Validate() != nil {
		return resBuilder.Build(), req.Validate()
	}
	claims, err := jwt.ValidateRefreshToken(a.logs, req.Token)
	if err != nil {
		return resBuilder.Build(), err
	}
	uUid := uuid.FromStringOrNil(claims.User.Id.(string))
	if claims.ExpiresAt < time.Now().Local().Unix() {
		if err = a.repo.GetRefreshToken().Delete(ctx, uUid); err != nil {
			return resBuilder.Build(), err
		}
		return resBuilder.Build(), errors.CustomError("token expired")
	}
	resMail := a.repo.GetAccount().GetByEmail(ctx, claims.User.Email)
	if resMail.Error != nil || *resMail.Value.IsActive != true {
		return resBuilder.Build(), errors.ErrSomethingWrong(a.logs, err)
	}

	clm := jwt.NewClaimsBuilder()
	clm.User(claims.User)
	clm.Client(claims.Client)

	clm.ExpiresAt(time.Now().Add(5 * time.Minute))
	token, err := jwt.GenerateToken(a.jwtKey, *clm.Build())
	if err != nil {
		return resBuilder.Build(), errors.ErrSomethingWrong(a.logs, err)
	}

	resBuilder.SetToken(token)
	return resBuilder.Build(), nil
}

// LoginAccount implements AccountService
func (a *accountUseCase) LoginAccount(ctx context.Context, req *requests.SignInWithEmailPassword) (res *responses.SignIn, err error) {
	resBuild := responses.NewSignInBuilder()
	if req.Validate() != nil {
		return resBuild.Build(), req.Validate()
	}
	dataCleint := ctx.Value("clients").(*entities.Client)

	if helpers.IsEmailValid(req.Email) == false {
		return resBuild.Build(), errors.ErrValidEmail(a.logs, string(req.Email))
	}

	userResult := a.repo.GetAccount().GetByEmail(ctx, req.Email)
	if userResult.Error != nil {
		return nil, userResult.Error
	}

	if !helpers.ComparePasswords(userResult.Value.Password, []byte(req.Password)) {
		return nil, errors.ErrInvalidAttributes("username and password")
	}

	roleResult := a.repo.GetRole().GetById(ctx, userResult.Value.RoleId)
	if roleResult.Error != nil {
		return nil, roleResult.Error
	}

	roleBuild := entities.NewRoleBuilder()
	roleBuild.SetId(userResult.Value.RoleId)
	roleBuild.SetName(roleResult.Value.Name)

	userBuild := entities.NewAccountBuilder()
	userBuild.SetId(userResult.Value.Id)
	userBuild.SetEmail(userResult.Value.Email)
	userBuild.SetRole(*roleBuild.Build())
	userBuild.SetIsActive(*userResult.Value.IsActive)
	userBuild.SetCreatedAt(userResult.Value.CreatedAt)
	userBuild.SetCreatedBy(userResult.Value.CreatedBy)

	cooBuild := entities.NewClientBuilder()
	cooBuild.SetIpClient(dataCleint.IpClient)
	cooBuild.SetUserAgent(dataCleint.UserAgent)

	clm := jwt.NewClaimsBuilder()
	clm.User(userBuild.Build())
	clm.Client(cooBuild.Build())

	clm.ExpiresAt(time.Now().Add(5 * time.Minute))
	token, err := jwt.GenerateToken(a.jwtKey, *clm.Build())
	if err != nil {
		return resBuild.Build(), errors.ErrSomethingWrong(a.logs, err)
	}

	clm.ExpiresAt(time.Now().Add(168 * time.Hour))
	refreshToken, err := jwt.GenerateToken(a.jwtKeyRefresh, *clm.Build())
	if err != nil {
		return resBuild.Build(), errors.ErrSomethingWrong(a.logs, err)
	}

	mRefreshToken := models.NewRefershTokenBuilder()
	mRefreshToken.SetId(helpers.GenerateUUIDV4(a.logs))
	mRefreshToken.SetUserId(userResult.Value.Id)
	mRefreshToken.SetToken(refreshToken)
	mRefreshToken.SetCreatedAt(time.Now().Local().UTC())
	mRefreshToken.SetCreatedBy(userBuild.Build().Id.(uuid.UUID))

	if err = a.repo.GetRefreshToken().Create(ctx, mRefreshToken.Build()); err != nil {
		return resBuild.Build(), err
	}

	resBuild.SetToken(token)
	resBuild.SetRefreshToken(refreshToken)
	return resBuild.Build(), nil
}

// CreateAccount implements AccountService
func (a *accountUseCase) CreateAccount(ctx context.Context, req *requests.CreateAccount) (res *responses.EntityId, err error) {
	resBuilder := responses.NewEntityIdBuilder()
	defaultStat := true

	if err = req.Validate(); err != nil {
		return resBuilder.Build(), err
	}

	if helpers.IsEmailValid(req.Email) != true {
		return nil, errors.ErrValidEmail(a.logs, req.Email)
	}

	mUser := models.NewAccountBuilder()
	mUser.SetId(helpers.GenerateUUIDV1(a.logs))
	mUser.SetEmail(req.Email)
	mUser.SetPassword(helpers.HashAndSalt([]byte(req.Password)))
	mUser.SetRoleId(req.Role)
	mUser.SetIsActive(&defaultStat)

	if err = a.repo.GetAccount().Create(ctx, mUser.Build()); err != nil {
		return resBuilder.Build(), err
	}

	return resBuilder.SetId(mUser.Build().Id).Build(), nil
}

// DeleteAccount implements AccountService
func (a *accountUseCase) DeleteAccount(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	result := a.repo.GetAccount().GetById(ctx, req.Id.(uuid.UUID))
	if result.Error != nil {
		return nil, result.Error
	}
	return new(responses.Empty), a.repo.GetAccount().Delete(ctx, req.Id.(uuid.UUID))
}

// DisableAccount implements AccountService
func (a *accountUseCase) DisableAccount(ctx context.Context, req *requests.DisableAccount) (res *responses.Empty, err error) {
	uUid := uuid.FromStringOrNil(req.Id.(string))
	result := a.repo.GetAccount().GetById(ctx, uUid)
	if result.Error != nil {
		return nil, result.Error
	}
	mBuild := models.NewAccountBuilder()
	mBuild.SetId(uUid)
	mBuild.SetIsActive(&req.IsActive)
	mBuild.SetUpdatedAt(time.Now())
	return new(responses.Empty), a.repo.GetAccount().UpdateStatus(ctx, mBuild.Build())
}

// ListAccount implements AccountService
func (a *accountUseCase) ListAccount(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Account], err error) {
	var (
		accounteData []entities.Account
		result       models.SliceResult[models.Account]
	)
	if req.Limit == 0 {
		req.Limit = 25
	}
	resBuilder := responses.NewListEntityBuilder[entities.Account]()
	if validation.IsEmpty(req.Keyword) {
		result = a.repo.GetAccount().GetAll(ctx, req.Limit)
		for _, v := range result.Value {
			resRole := a.repo.GetRole().GetById(ctx, v.RoleId)
			roleBuild := entities.NewRoleBuilder()
			roleBuild.SetId(resRole.Value.Id)
			roleBuild.SetName(resRole.Value.Name)

			accountBuild := entities.NewAccountBuilder()
			accountBuild.SetId(v.Id)
			accountBuild.SetEmail(v.Email)
			accountBuild.SetIsActive(*v.IsActive)
			accountBuild.SetRole(*roleBuild.Build())
			accountBuild.SetCreatedAt(v.CreatedAt)
			accountBuild.SetCreatedBy(v.CreatedBy)
			accounteData = append(accounteData, *accountBuild.Build())
		}
		resBuilder.SetData(accounteData)
		return resBuilder.Build(), result.Error
	}
	result = a.repo.GetAccount().GetByKeywords(ctx, req.Keyword, req.Limit)
	for _, v := range result.Value {
		resRole := a.repo.GetRole().GetById(ctx, v.RoleId)
		roleBuild := entities.NewRoleBuilder()
		roleBuild.SetId(resRole.Value.Id)
		roleBuild.SetName(resRole.Value.Name)

		accountBuild := entities.NewAccountBuilder()
		accountBuild.SetId(v.Id)
		accountBuild.SetEmail(v.Email)
		accountBuild.SetIsActive(*v.IsActive)
		accountBuild.SetRole(*roleBuild.Build())
		accountBuild.SetCreatedAt(v.CreatedAt)
		accountBuild.SetCreatedBy(v.CreatedBy)
		accounteData = append(accounteData, *accountBuild.Build())
	}
	resBuilder.SetData(accounteData)
	return resBuilder.Build(), result.Error
}

// ViewAccountByID implements AccountService
func (a *accountUseCase) ViewAccountByID(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Account], err error) {
	if err = req.Validate(); err != nil {
		return nil, err
	}

	result := a.repo.GetAccount().GetById(ctx, req.Id.(uuid.UUID))
	if result.Error != nil {
		return nil, result.Error
	}

	resRole := a.repo.GetRole().GetById(ctx, result.Value.RoleId)
	roleBuild := entities.NewRoleBuilder()
	roleBuild.SetId(resRole.Value.Id)
	roleBuild.SetName(resRole.Value.Name)

	mBuild := entities.NewAccountBuilder()
	mBuild.SetId(result.Value.Id)
	mBuild.SetEmail(result.Value.Email)
	mBuild.SetRole(*roleBuild.Build())
	mBuild.SetCreatedAt(result.Value.CreatedAt)
	mBuild.SetCreatedBy(result.Value.CreatedBy)

	resBuild := responses.NewViewEntityBuilder[*entities.Account]()
	return resBuild.SetData(mBuild.Build()).Build(), nil
}

func NewAccount(logs *zap.Logger, jwtKey string, jwtKeyRefresh string, r repositories.Repositories) AccountService {
	a := new(accountUseCase)
	a.logs = logs
	a.jwtKey = jwtKey
	a.jwtKeyRefresh = jwtKeyRefresh
	a.repo = r
	return a
}
