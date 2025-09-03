package auth

import (
	"context"

	"github.com/anonychun/ecorp/internal/entity"
)

func (u *Usecase) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	admin, err := u.repository.Admin.FindByEmailAddress(ctx, req.EmailAddress)
	if err != nil {
		return nil, err
	}

	adminSession := &entity.AdminSession{
		AdminId:   admin.Id,
		IpAddress: req.IpAddress,
		UserAgent: req.UserAgent,
	}
	adminSession.GenerateToken()

	err = u.repository.AdminSession.Create(ctx, adminSession)
	if err != nil {
		return nil, err
	}

	res := &LoginResponse{Token: adminSession.Token}
	res.Admin.Id = admin.Id.String()
	res.Admin.EmailAddress = admin.EmailAddress

	return res, nil
}

func (u *Usecase) Logout(ctx context.Context, req LogoutRequest) error {
	adminSession, err := u.repository.AdminSession.FindByToken(ctx, req.Token)
	if err != nil {
		return err
	}

	err = u.repository.AdminSession.DeleteById(ctx, adminSession.Id.String())
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) Me(ctx context.Context) (*MeResponse, error) {
	return &MeResponse{}, nil
}
