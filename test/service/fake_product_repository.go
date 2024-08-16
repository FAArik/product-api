package service

import (
	"errors"
	"fmt"
	"product-app/domain"
	"product-app/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return fakeRepository.products
}
func (fakeRepository *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	products := fakeRepository.products
	var getProductsByStore []domain.Product

	for _, s := range products {
		if s.Store == storeName {
			getProductsByStore = append(getProductsByStore, s)
		}
	}
	return getProductsByStore
}
func (fakeRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeRepository.products = append(fakeRepository.products, domain.Product{
		Id:       int64(len(fakeRepository.products)) + 1,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}
func (fakeRepository *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	products := fakeRepository.products
	var findByIdVal *domain.Product
	for _, s := range products {
		if s.Id == productId {
			findByIdVal = &s
		}
	}
	if findByIdVal == nil {
		return domain.Product{}, errors.New(fmt.Sprintf("Product with %d not found", productId))
	}
	return *findByIdVal, nil
}
func (fakeRepository *FakeProductRepository) DeleteById(productId int64) error {
	_, err := fakeRepository.GetById(productId)
	if err != nil {
		return errors.New(fmt.Sprintf(err.Error()))
	}
	products := fakeRepository.products
	var notDeletedProducts []domain.Product

	for _, s := range products {
		if s.Id != productId {
			notDeletedProducts = append(notDeletedProducts, s)
		}
	}
	fakeRepository.products = notDeletedProducts
	return nil
}
func (fakeRepository *FakeProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	_, err := fakeRepository.GetById(productId)
	if err != nil {
		return errors.New(fmt.Sprintf(err.Error()))
	}
	products := fakeRepository.products
	for _, s := range products {
		if s.Id == productId {
			s.Price = newPrice
		}
	}
	return nil
}
