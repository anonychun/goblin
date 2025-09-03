package admin

import (
	"context"
	"net/http"

	"github.com/anonychun/ecorp/internal/api"
	"github.com/anonychun/ecorp/internal/entity"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func (u *Usecase) FindAll(ctx context.Context) ([]*AdminResponse, error) {
	admins, err := u.repository.Admin.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	res := lo.Map(admins, func(admin *entity.Admin, _ int) *AdminResponse {
		return NewAdminResponse(admin)
	})

	return res, nil
}

func (u *Usecase) FindById(ctx context.Context, req FindByIdRequest) (*AdminResponse, error) {
	admin, err := u.repository.Admin.FindById(ctx, req.Id)
	if err == gorm.ErrRecordNotFound {
		return nil, &api.Error{Status: http.StatusNotFound, Errors: "Admin not found"}
	} else if err != nil {
		return nil, err
	}

	return NewAdminResponse(admin), nil
}

func (u *Usecase) Create(ctx context.Context, req CreateRequest) (*AdminResponse, error) {
	admin := &entity.Admin{
		Name: req.Name,
	}

	if err := u.repository.Admin.Create(ctx, admin); err != nil {
		return nil, err
	}

	return NewAdminResponse(admin), nil
}

func (u *Usecase) Update(ctx context.Context, req UpdateRequest) (*AdminResponse, error) {
	admin, err := u.repository.Admin.FindById(ctx, req.Id)
	if err == gorm.ErrRecordNotFound {
		return nil, &api.Error{Status: http.StatusNotFound, Errors: "Admin not found"}
	} else if err != nil {
		return nil, err
	}

	admin.Name = req.Name

	if err := u.repository.Admin.Update(ctx, admin); err != nil {
		return nil, err
	}

	return NewAdminResponse(admin), nil
}

func (u *Usecase) Delete(ctx context.Context, req DeleteRequest) error {
	return u.repository.Admin.DeleteById(ctx, req.Id)
}
