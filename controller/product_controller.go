package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"product-app/controller/requests"
	"product-app/controller/response"
	"product-app/service"
	"strconv"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (ProductController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products/:id", ProductController.GetProductById)
	e.GET("/api/v1/products", ProductController.GetAllProducts)
	e.POST("/api/v1/products", ProductController.AddProduct)
	e.PUT("/api/v1/products/:id", ProductController.UpdatePrice)
	e.DELETE("/api/v1/products/:id", ProductController.DeleteProductById)
}

func (ProductController *ProductController) GetProductById(c echo.Context) error {
	param := c.Param("id")
	productId, errParse := strconv.Atoi(param)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: errParse.Error()})
	}
	prod, err := ProductController.productService.GetById(int64(productId))

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{ErrorDescription: err.Error()})
	}
	return c.JSON(http.StatusOK, response.ToResponse(prod))
}

func (ProductController *ProductController) GetAllProducts(c echo.Context) error {
	store := c.QueryParam("Store")
	if len(store) == 0 {
		allProducts := ProductController.productService.GetAllProducts()
		return c.JSON(http.StatusOK, allProducts)
	}
	productsByStore := ProductController.productService.GetAllProductsByStore(store)
	return c.JSON(http.StatusOK, response.ToResponseList(productsByStore))
}

func (ProductController *ProductController) AddProduct(c echo.Context) error {
	var addProductRequest requests.AddProductRequest
	err := c.Bind(&addProductRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}
	errAdd := ProductController.productService.Add(addProductRequest.ToModel())
	if errAdd != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{ErrorDescription: errAdd.Error()})
	}
	return c.NoContent(http.StatusCreated)
}

func (ProductController *ProductController) UpdatePrice(c echo.Context) error {
	param := c.Param("id")
	productId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{ErrorDescription: err.Error()})
	}
	queryParam := c.QueryParam("newPrice")
	if len(queryParam) == 0 {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{ErrorDescription: "Newprice parameter is mandatory!"})
	}
	newPrice, parseErr := strconv.ParseFloat(queryParam, 32)
	if parseErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{"New Price Distrupted"})
	}
	updateErr := ProductController.productService.UpdatePrice(int64(productId), float32(newPrice))
	if updateErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{updateErr.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (ProductController *ProductController) DeleteProductById(c echo.Context) error {
	param := c.Param("id")
	productId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{ErrorDescription: err.Error()})
	}
	deleteErr := ProductController.productService.DeleteById(int64(productId))
	if deleteErr != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{deleteErr.Error()})
	}
	return c.NoContent(http.StatusOK)
}
