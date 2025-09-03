package db

import (
	"context"

	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/internal/entity"
	"github.com/samber/do"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func init() {
	do.Provide(bootstrap.Injector, NewSeeder)
}

type Seeder struct {
	sql *Sql
}

func NewSeeder(i *do.Injector) (*Seeder, error) {
	return &Seeder{
		sql: do.MustInvoke[*Sql](i),
	}, nil
}

func (s *Seeder) Seed() error {
	ctx := context.Background()
	defaultAdminEmail := "anonychun@gmail.com"
	defaultAdminPassword := "didbnyaada"
	hashedDefaultAdminPassword, err := bcrypt.GenerateFromPassword([]byte(defaultAdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	defaultAdmin := &entity.Admin{
		Name:           "Achmad Chun Chun",
		EmailAddress:   defaultAdminEmail,
		PasswordDigest: string(hashedDefaultAdminPassword),
	}
	err = s.sql.DB(ctx).First(defaultAdmin, "email_address = ?", defaultAdmin.EmailAddress).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	defaultAdmin.DeletedAt = nil
	err = s.sql.DB(ctx).Save(defaultAdmin).Error
	if err != nil {
		return err
	}

	return nil
}
