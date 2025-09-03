package article

import (
	"context"

	"github.com/anonychun/ecorp/internal/entity"
)

func (r *Repository) FindAll(ctx context.Context) ([]*entity.Article, error) {
	articles := make([]*entity.Article, 0)
	err := r.sql.DB(ctx).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *Repository) FindById(ctx context.Context, id string) (*entity.Article, error) {
	article := &entity.Article{}
	err := r.sql.DB(ctx).First(article, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (r *Repository) Create(ctx context.Context, article *entity.Article) error {
	return r.sql.DB(ctx).Create(article).Error
}

func (r *Repository) Update(ctx context.Context, article *entity.Article) error {
	return r.sql.DB(ctx).Save(article).Error
}

func (r *Repository) DeleteById(ctx context.Context, id string) error {
	return r.sql.DB(ctx).Delete(&entity.Article{}, "id = ?", id).Error
}

func (r *Repository) ExistsById(ctx context.Context, id string) (bool, error) {
	var exists bool
	err := r.sql.DB(ctx).Raw("SELECT EXISTS(SELECT 1 FROM articles WHERE id = ?)", id).Scan(&exists).Error
	return exists, err
}
