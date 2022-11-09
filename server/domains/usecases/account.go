package usecases

import (
	"context"
	"time"

	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/bagasunix/ginclean/pkg/helpers"
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
	ViewAccountByEmail(ctx context.Context, email string) (res *responses.ViewEntity[*entities.Account], err error)
	DisableAccount(ctx context.Context, req *requests.DisableAccount) (res *responses.Empty, err error)
	DeleteAccount(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error)
}

type AccountUseCase struct {
	logs zap.Logger
	repo repositories.Repositories
}

// CreateAccount implements AccountService
func (a *AccountUseCase) CreateAccount(ctx context.Context, req *requests.CreateAccount) (res *responses.EntityId, err error) {
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
func (a *AccountUseCase) DeleteAccount(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error) {
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
func (a *AccountUseCase) DisableAccount(ctx context.Context, req *requests.DisableAccount) (res *responses.Empty, err error) {
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
func (a *AccountUseCase) ListAccount(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Account], err error) {
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

// ViewAccountByEmail implements AccountService
func (*AccountUseCase) ViewAccountByEmail(ctx context.Context, email string) (res *responses.ViewEntity[*entities.Account], err error) {
	panic("unimplemented")
}

// ViewAccountByID implements AccountService
func (a *AccountUseCase) ViewAccountByID(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Account], err error) {
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

func NewAccount(logs zap.Logger, r repositories.Repositories) AccountService {
	a := new(AccountUseCase)
	a.logs = logs
	a.repo = r
	return a
}
