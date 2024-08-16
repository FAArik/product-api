package service

import (
	"os"
	"product-app/domain"
	"product-app/service"
	"product-app/service/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productService service.IProductService

func TestMain(m *testing.M) {

	initialProducts := []domain.Product{
		{
			Id:    1,
			Name:  "AirFryer",
			Price: 1000.0,
			Store: "ABC Tech",
		},
		{
			Id:    2,
			Name:  "Ütü",
			Price: 4000.0,
			Store: "ABC Tech",
		},
		{
			Id:    3,
			Name:  "Telefon",
			Price: 4000.0,
			Store: "CBA TECH",
		},
	}
	fakeProductRepository := NewFakeProductRepository(initialProducts)

	productService = service.NewProductService(fakeProductRepository)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_ShouldGetAllProducts(t *testing.T) {
	t.Run("ShouldGetAllProducts", func(t *testing.T) {
		products := productService.GetAllProducts()
		assert.Equal(t, 3, len(products))
	})
}

func Test_ShouldGetAllProductsByStoreName(t *testing.T) {
	t.Run("ShouldGetAllProductsByStoreName", func(t *testing.T) {
		products := productService.GetAllProductsByStore("ABC Tech")
		assert.Equal(t, 2, len(products))
	})
}
func Test_ShouldDeleteProductById(t *testing.T) {
	t.Run("ShouldDeleteProductById", func(t *testing.T) {
		products := productService.DeleteById(1)
		_, err := productService.GetById(1)
		assert.Equal(t, nil, products)
		assert.Equal(t, "Product with 1 not found", err.Error())
	})
}
func Test_ShouldAddNewProduct(t *testing.T) {
	t.Run("ShouldAddNewProduct", func(t *testing.T) {
		adderr := productService.Add(dto.ProductCreate{
			Name:     "Laptop",
			Price:    15000.0,
			Discount: 12.0,
			Store:    "Foldable Electro",
		})
		products := productService.GetAllProducts()
		prd, _ := productService.GetById(int64(len(products)))
		assert.Equal(t, nil, adderr)
		assert.Equal(t, domain.Product{
			Id:       4,
			Name:     "Laptop",
			Price:    15000.0,
			Discount: 12.0,
			Store:    "Foldable Electro",
		}, prd)
	})
}
