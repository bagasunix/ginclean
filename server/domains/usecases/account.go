package usecases

// type AccountService interface {
// 	CreateAccount(ctx context.Context, req *requests.CreateAccount) (res *responses.EntityId, err error)
// 	ListAccount(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Account], err error)
// 	ViewAccountByID(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Account], err error)
// 	ViewAccountByEmail(ctx context.Context, email string) (res *responses.ViewEntity[*entities.Account], err error)
// 	DisableAccount(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error)
// 	DisableMultipleAccount(ctx context.Context, req []string) (res *responses.ListMultiple[*entities.Account, *entities.Account], err error)
// 	DeleteAccount(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error)
// }

// type AccountUseCase struct {
// 	logs zap.Logger
// 	repo repositories.Repositories
// }

// // CreateAccount implements AccountService
// func (a *AccountUseCase) CreateAccount(ctx context.Context, req *requests.CreateAccount) (res *responses.EntityId, err error) {
// 	resBuilder := responses.NewEntityIdBuilder()

// 	if err = req.Validate(); err != nil {
// 		return resBuilder.Build(), err
// 	}

// 	if helpers.IsEmailValid(req.Email) != true {
// 		return nil, errors.ErrValidEmail(a.logs, req.Email)
// 	}

// 	mUser := models.NewAccountBuilder()
// 	mUser.SetId(helpers.GenerateUUIDV1(a.logs))
// 	mUser.SetEmail(req.Email)
// 	mUser.SetPassword(helpers.HashAndSalt([]byte(req.Password)))
// 	mUser.SetRoleId(req.Role)
// 	mUser.SetIsActive(true)

// 	if err = a.repo.GetAccount().Create(ctx, mUser.Build()); err != nil {
// 		return resBuilder.Build(), err
// 	}

// 	return resBuilder.SetId(mUser.Build().Id).Build(), nil
// }

// // DeleteAccount implements AccountService
// func (*AccountUseCase) DeleteAccount(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error) {
// 	panic("unimplemented")
// }

// // DisableAccount implements AccountService
// func (*AccountUseCase) DisableAccount(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
// 	panic("unimplemented")
// }

// // DisableMultipleAccount implements AccountService
// func (*AccountUseCase) DisableMultipleAccount(ctx context.Context, req []string) (res *responses.ListMultiple[*entities.Account, *entities.Account], err error) {
// 	panic("unimplemented")
// }

// // ListAccount implements AccountService
// func (*AccountUseCase) ListAccount(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Account], err error) {
// 	panic("unimplemented")
// }

// // ViewAccountByEmail implements AccountService
// func (*AccountUseCase) ViewAccountByEmail(ctx context.Context, email string) (res *responses.ViewEntity[*entities.Account], err error) {
// 	panic("unimplemented")
// }

// // ViewAccountByID implements AccountService
// func (*AccountUseCase) ViewAccountByID(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Account], err error) {
// 	panic("unimplemented")
// }

// func NewAccount(logs zap.Logger, r repositories.Repositories) AccountService {
// 	a := new(AccountUseCase)
// 	a.logs = logs
// 	a.repo = r
// 	return a
// }
