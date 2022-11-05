package usecases

import (
	"context"
	"time"

	"github.com/bagasunix/ginclean/pkg/helpers"
	"github.com/bagasunix/ginclean/server/domains/data/models"
	"github.com/bagasunix/ginclean/server/domains/data/repositories"
	"github.com/bagasunix/ginclean/server/domains/entities"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/bagasunix/ginclean/server/endpoints/responses"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofrs/uuid"
)

type RoleService interface {
	CreateRole(ctx context.Context, req *requests.CreateRole) (res *responses.EntityId, err error)
	ListRole(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Role], err error)
	ViewRole(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Role], err error)
	UpdateRole(ctx context.Context, req *requests.UpdateRole) (res *responses.Empty, err error)
	UpdateMultipleRole(ctx context.Context, req *[]requests.UpdateRole) (res *responses.ListMultiple[requests.UpdateRole, requests.UpdateRole], err error)
	DeleteRole(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error)
}

type RoleUseCase struct {
	repo repositories.Repositories
}

// UpdateMultipleRole implements RoleService
func (r *RoleUseCase) UpdateMultipleRole(ctx context.Context, req *[]requests.UpdateRole) (res *responses.ListMultiple[requests.UpdateRole, requests.UpdateRole], err error) {
	var (
		dataFailed, dataSuccess []requests.UpdateRole
		counterFail             int = 0
	)

	resBuilder := responses.NewListMultipleBuilder[requests.UpdateRole, requests.UpdateRole]()
	for _, v := range *req {
		if err = v.Validate(); err != nil {
			// return nil, err
			continue

		}
		uuID := uuid.Must(uuid.FromString(v.Id.(string)))

		result := r.repo.GetRole().GetById(ctx, uuID)
		if result.Error != nil {
			counterFail += 1
			dataFailed = append(dataFailed, v)
			// return nil, result.Error
			continue

		}

		mBuild := models.NewRoleBuilder()
		mBuild.SetId(uuID)
		mBuild.SetName(v.Name)
		mBuild.SetUpdatedAt(time.Now())

		if err := r.repo.GetRole().Update(ctx, mBuild.Build()); err != nil {
			counterFail += 1
			dataFailed = append(dataFailed, v)
			// return nil, result.Error
			continue

		}

		dataSuccess = append(dataSuccess, v)
	}
	resBuilder.SetDataMulti(dataSuccess, dataFailed)
	return resBuilder.Build(), err

}

// DeleteRole implements RoleService
func (r *RoleUseCase) DeleteRole(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	result := r.repo.GetRole().GetById(ctx, req.Id.(uuid.UUID))
	if result.Error != nil {
		return nil, result.Error
	}
	return new(responses.Empty), r.repo.GetRole().Delete(ctx, req.Id.(uuid.UUID))
}

// UpdateRole implements RoleService
func (r *RoleUseCase) UpdateRole(ctx context.Context, req *requests.UpdateRole) (res *responses.Empty, err error) {
	if err = req.Validate(); err != nil {
		return nil, err
	}
	result := r.repo.GetRole().GetById(ctx, req.Id.(uuid.UUID))
	if result.Error != nil {
		return nil, result.Error
	}
	mBuild := models.NewRoleBuilder()
	mBuild.SetId(req.Id.(uuid.UUID))
	mBuild.SetName(req.Name)
	mBuild.SetUpdatedAt(time.Now())

	return new(responses.Empty), r.repo.GetRole().Update(ctx, mBuild.Build())
}

// ViewRole implements RoleService
func (r *RoleUseCase) ViewRole(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Role], err error) {
	if err = req.Validate(); err != nil {
		return nil, err
	}

	result := r.repo.GetRole().GetById(ctx, req.Id.(uuid.UUID))
	if result.Error != nil {
		return nil, result.Error
	}

	mBuild := entities.NewRoleBuilder()
	mBuild.SetId(result.Value.Id)
	mBuild.SetName(result.Value.Name)
	mBuild.SetCreatedAt(result.Value.CreatedAt)
	mBuild.SetCreatedBy(result.Value.CreatedBy)

	resBuild := responses.NewViewEntityBuilder[*entities.Role]()
	return resBuild.SetData(mBuild.Build()).Build(), nil
}

// ListRole implements RoleService
func (r *RoleUseCase) ListRole(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Role], err error) {
	if req.Limit == 0 {
		req.Limit = 25
	}
	resBuilder := responses.NewListEntityBuilder[entities.Role]()
	var roleData []entities.Role
	if validation.IsEmpty(req.Keyword) {
		result := r.repo.GetRole().GetAll(ctx, req.Limit)
		for _, v := range result.Value {
			roleBuild := entities.NewRoleBuilder()
			roleBuild.SetId(v.Id)
			roleBuild.SetName(v.Name)
			roleBuild.SetCreatedAt(v.CreatedAt)
			roleBuild.SetCreatedBy(v.CreatedBy)
			roleData = append(roleData, *roleBuild.Build())
		}
		resBuilder.SetData(roleData)
		return resBuilder.Build(), result.Error
	}
	result := r.repo.GetRole().GetByKeywords(ctx, req.Keyword, req.Limit)
	for _, i := range result.Value {
		roleBuild := entities.NewRoleBuilder()
		roleBuild.SetId(i.Id)
		roleBuild.SetName(i.Name)
		roleBuild.SetCreatedAt(i.CreatedAt)
		roleBuild.SetCreatedBy(i.CreatedBy)
		roleData = append(roleData, *roleBuild.Build())
	}
	resBuilder.SetData(roleData)
	return resBuilder.Build(), result.Error
}

// CreateRole implements RoleService
func (r *RoleUseCase) CreateRole(ctx context.Context, req *requests.CreateRole) (res *responses.EntityId, err error) {
	resBuild := responses.NewEntityIdBuilder()

	if err = req.Validate(); err != nil {
		return resBuild.Build(), err
	}
	mRole := models.NewRoleBuilder()
	mRole.SetId(helpers.GenerateUUIDV1())
	mRole.SetName(req.Name)
	mRole.SetCreatedAt(time.Now().UTC().Local())

	if err = r.repo.GetRole().Create(ctx, mRole.Build()); err != nil {
		return resBuild.Build(), err
	}

	return resBuild.SetId(mRole.Build().Id).Build(), nil
}

func NewRole(r repositories.Repositories) RoleService {
	u := new(RoleUseCase)
	u.repo = r
	return u
}
