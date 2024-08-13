package infrastructure

import (
	"context"
	"fmt"
	"os"
	"product-app/common/postgresql"
	"product-app/domain"
	"product-app/persistence"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()
	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: time.Second * 30,
	})
	productRepository = persistence.NewProductRepository(dbPool)
	fmt.Println("Befora all tests")
	exitCode := m.Run()
	fmt.Println("After all tests")
	os.Exit(exitCode)

}
func setup(ctx context.Context, dbpool *pgxpool.Pool) {
	fmt.Println("Initializing test data")
	TestDataInitialize(ctx, dbpool)
}
func clear(ctx context.Context, dbpool *pgxpool.Pool) {
	fmt.Println("Clearing test data")
	TruncateTestData(ctx, dbpool)
}

func TestGetAllProducts(t *testing.T) {
	fmt.Println(dbPool)
	setup(ctx, dbPool)

	expectedProducts := []domain.Product{
		{Id: 1, Name: "AirFryer", Price: 3000.0, Discount: 20.0, Store: "ABC TECH"},
		{Id: 2, Name: "Ütü", Price: 1500.0, Discount: 10.0, Store: "ABC TECH"},
		{Id: 3, Name: "Çamaşır Makinesi", Price: 15000.0, Discount: 24.0, Store: "ABC TECH"},
		{Id: 4, Name: "Lambader", Price: 2500.0, Discount: 8.0, Store: "ABC Dekorasyon"},
	}
	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestGetAllProductsByStore(t *testing.T) {
	fmt.Println(dbPool)
	setup(ctx, dbPool)

	expectedProducts := []domain.Product{
		{Id: 1, Name: "AirFryer", Price: 3000.0, Discount: 20.0, Store: "ABC TECH"},
		{Id: 2, Name: "Ütü", Price: 1500.0, Discount: 10.0, Store: "ABC TECH"},
		{Id: 3, Name: "Çamaşır Makinesi", Price: 15000.0, Discount: 24.0, Store: "ABC TECH"},
	}
	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProducts := productRepository.GetAllProductsByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})
	clear(ctx, dbPool)
}
func TestAddProduct(t *testing.T) {

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "Kalem",
			Price:    80.0,
			Discount: 10.0,
			Store:    "ABC KIRTASIYE",
		},
	}
	newProduct := domain.Product{
		Name:     "Kalem",
		Price:    80.0,
		Discount: 10.0,
		Store:    "ABC KIRTASIYE",
	}
	t.Run("GetAllProductsByStore", func(t *testing.T) {
		err := productRepository.AddProduct(newProduct)
		addedProduct := productRepository.GetAllProducts()
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, len(addedProduct))
		assert.Equal(t, expectedProducts, addedProduct)
	})
	clear(ctx, dbPool)
}
func TestGetById(t *testing.T) {
	setup(ctx, dbPool)

	expectedProduct := domain.Product{
		Id:       1,
		Name:     "AirFryer",
		Price:    3000.0,
		Discount: 20.0,
		Store:    "ABC TECH",
	}

	t.Run("GetProductById", func(t *testing.T) {
		actualProduct, err := productRepository.GetById(1)
		_, geterr := productRepository.GetById(5)
		assert.Equal(t, expectedProduct, actualProduct)
		assert.Equal(t, "Product not found with id 5", geterr.Error())
		assert.Equal(t, nil, err)
	})
	clear(ctx, dbPool)
}
func TestDeleteById(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("DeleteById", func(t *testing.T) {
		err := productRepository.DeleteById(1)
		products := productRepository.GetAllProducts()
		_, getByIdErr := productRepository.GetById(1)
		assert.Equal(t, 3, len(products))
		assert.Equal(t, nil, err)
		assert.Equal(t, "Product not found with id 1", getByIdErr.Error())
	})
	clear(ctx, dbPool)
}

func TestUpdatePrice(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("UpdatePrice", func(t *testing.T) {
		productBeforeUpdate, _ := productRepository.GetById(1)
		assert.Equal(t, float32(3000.0), productBeforeUpdate.Price)
		productRepository.UpdatePrice(1, 4000)
		productAfterUpdate, _ := productRepository.GetById(1)
		assert.Equal(t, float32(4000.0), productAfterUpdate.Price)
	})
	clear(ctx, dbPool)
}
