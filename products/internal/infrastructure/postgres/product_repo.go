package postgres

import (
	"context"
	"github.com/juancanchi/jujuy-market/products/internal/domain"
)

type MockProductRepository struct {
	data []*domain.Product
}

func NewProductRepository(_ interface{}) *MockProductRepository {
	return &MockProductRepository{data: []*domain.Product{}}
}

func (r *MockProductRepository) Save(ctx context.Context, p *domain.Product) error {
	r.data = append(r.data, p)
	return nil
}

func (r *MockProductRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	return r.data, nil
}

// Placeholder para evitar error en main.go al llamar a postgres.NewDB()
func NewDB(_ string) (interface{}, error) {
	return nil, nil
}
