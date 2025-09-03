package repository

import (
	"context"

	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/internal/db"
	"github.com/anonychun/ecorp/internal/repository/admin"
	"github.com/anonychun/ecorp/internal/repository/adminsession"
	"github.com/anonychun/ecorp/internal/repository/article"
	"github.com/samber/do"
)

func init() {
	do.Provide(bootstrap.Injector, NewRepository)
}

type Repository struct {
	sql *db.Sql

	Admin        *admin.Repository
	AdminSession *adminsession.Repository
	Article      *article.Repository
}

func NewRepository(i *do.Injector) (*Repository, error) {
	return &Repository{
		sql: do.MustInvoke[*db.Sql](i),

		Admin:        do.MustInvokeNamed[*admin.Repository](i, admin.RepositoryInjectorName),
		AdminSession: do.MustInvokeNamed[*adminsession.Repository](i, adminsession.RepositoryInjectorName),
		Article:      do.MustInvokeNamed[*article.Repository](i, article.RepositoryInjectorName),
	}, nil
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.sql.Transaction(ctx, fn)
}
