package persistence

import (
	"context"
	"product-app/domain"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
}

type ProductRepository struct {
	dbPool pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool: *pool,
	}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	prodRows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products")
	if err != nil {
		log.Error("Error while getting all products %v", err)

		return []domain.Product{}
	}
	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string
	for prodRows.Next() {
		prodRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}
	return products
}
