package usecases

import (
	"context"
	"time"

	"github.com/bagasunix/ginclean/domains/data/models"
	"github.com/bagasunix/ginclean/domains/data/repositories"
	"github.com/bagasunix/ginclean/domains/entities"
	"github.com/bagasunix/ginclean/endpoints/requests"
	"github.com/bagasunix/ginclean/endpoints/responses"
	"github.com/bagasunix/ginclean/pkg/helpers"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofrs/uuid"
)

type RoleService interface {
	CreateRole(ctx context.Context, req *requests.CreateRole) (res *responses.EntityId, err error)
	ListRole(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Role], err error)
	ViewRole(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Role], err error)
	UpdateRole(ctx context.Context, req *requests.UpdateRole) (res *responses.Empty, err error)
	DeleteRole(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error)
}

type RoleUseCase struct {
	repo repositories.Repositories
}

// DeleteRole implements RoleService
func (r *RoleUseCase) DeleteRole(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return new(responses.Empty), r.repo.GetRole().Delete(ctx, req.Id.(uuid.UUID))
}

// UpdateRole implements RoleService
func (r *RoleUseCase) UpdateRole(ctx context.Context, req *requests.UpdateRole) (res *responses.Empty, err error) {
	if err = req.Validate(); err != nil {
		return nil, err
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
