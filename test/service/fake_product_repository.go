package service

import (
	"product-app/domain"
	"product-app/persistence"
)

type FakeProductReposityory struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductReposityory{
		products: initialProducts,
	}
}
