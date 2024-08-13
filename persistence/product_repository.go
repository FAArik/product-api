package persistence

import (
	"context"
	"errors"
	"fmt"
	"product-app/domain"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
	AddProduct(product domain.Product) error
	GetById(productId int64) (domain.Product, error)
	DeleteById(productId int64) error
	UpdatePrice(productId int64, newPrice float32) error
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
	return extractProductsFromRows(prodRows)
}
func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()
	getProductsByStoreSql := `SELECT * FROM products where store = $1`
	prodRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreSql, storeName)
	if err != nil {
		log.Error("Error while getting all products %v", err)

		return []domain.Product{}
	}
	return extractProductsFromRows(prodRows)
}
func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()
	insert_sql := `INSERT INTO products(name,price,discount,store) VALUES ($1,$2,$3,$4)`
	addedProduct, insertErr := productRepository.dbPool.Exec(ctx, insert_sql, product.Name, product.Price, product.Discount, product.Store)
	if insertErr != nil {
		log.Error("Failed to add new product ", insertErr)
		return insertErr
	}
	log.Info(fmt.Printf("Product addet with %v", addedProduct))
	return nil
}
func (productRepository *ProductRepository) GetById(productId int64) (domain.Product, error) {
	ctx := context.Background()
	getByIdSql := `select * from products where id = $1`
	queryRow := productRepository.dbPool.QueryRow(ctx, getByIdSql, productId)
	var id int64
	var name string
	var price float32
	var discount float32
	var store string
	err := queryRow.Scan(&id, &name, &price, &discount, &store)

	if err != nil && err.Error() == "no rows in result set" {

		return domain.Product{}, errors.New(fmt.Sprintf("Product not found with id %d", productId))
	}

	if err != nil {
		return domain.Product{}, errors.New(fmt.Sprintf("Error while getting product with id %d", productId))
	}

	return domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}, nil
}
func (productRepository *ProductRepository) DeleteById(productId int64) error {
	ctx := context.Background()
	_, err := productRepository.GetById(productId)

	if err != nil {
		return errors.New("Product not found")
	}
	deleteSql := "DELETE FROM products where id=$1"
	_, deleteErr := productRepository.dbPool.Exec(ctx, deleteSql, productId)
	if deleteErr != nil {
		return errors.New(fmt.Sprintf("Error while deleting product with id %d", productId))
	}
	log.Info("Product deleted")

	return nil
}
func extractProductsFromRows(prodRows pgx.Rows) []domain.Product {
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
func (productRepository *ProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	ctx := context.Background()

	_, getByIdErr := productRepository.GetById(productId)
	if getByIdErr != nil {
		return errors.New(fmt.Sprintf("product with id %d not found ", productId))
	}

	var updateSql = `Update products set price =$1 where id =$2`
	_, err := productRepository.dbPool.Exec(ctx, updateSql, newPrice, productId)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while updating the product with id %d", productId))
	}
	log.Info("Product %d price updated with new price %v", productId, newPrice)
	return nil
}
