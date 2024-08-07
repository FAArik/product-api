package infrastructure

import (
	"context"
	"fmt"
	"os"
	"product-app/common/postgresql"
	"product-app/persistence"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()
	dbPool := postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: time.Second * 30,
	})
	productRepository = persistence.NewProductRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)

}
func TestGetAllProducts(t *testing.T) {

	fmt.Println(productRepository)
	fmt.Println(dbPool)
}
