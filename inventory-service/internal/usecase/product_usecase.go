package usecase

import (
	"context"
	"inventory-service/internal/models"
	"inventory-service/internal/repository"
)

type ProductUseCase interface {
	Create(ctx context.Context, p *models.Product) error
	Get(ctx context.Context, id string) (*models.Product, error)
	List(ctx context.Context) ([]*models.Product, error)
	Update(ctx context.Context, id string, p *models.Product) error
	Delete(ctx context.Context, id string) error
	ListFiltered(ctx context.Context, category, limit, offset string) ([]*models.Product, error)
}

type productUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}

func (u *productUseCase) Create(ctx context.Context, p *models.Product) error {
	return u.repo.Create(ctx, p)
}

func (u *productUseCase) Get(ctx context.Context, id string) (*models.Product, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *productUseCase) List(ctx context.Context) ([]*models.Product, error) {
	return u.repo.List(ctx)
}

func (u *productUseCase) Update(ctx context.Context, id string, p *models.Product) error {
	return u.repo.Update(ctx, id, p)
}

func (u *productUseCase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

func (u *productUseCase) ListFiltered(ctx context.Context, category, limit, offset string) ([]*models.Product, error) {
	return u.repo.ListFiltered(ctx, category, limit, offset)
}
