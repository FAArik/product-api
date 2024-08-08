package infrastructure

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

var INSERT_PRODUCTS = `INSERT INTO products (name,price,discount,store)
VALUES('AirFryer',3000.0,20.0,'ABC TECH'),
('Ütü',1500.0,10.0,'ABC TECH'),
('Çamaşır Makinesi',15000.0,24.0,'ABC TECH'),
('Lambader',2500.0,8.0,'ABC Dekorasyon')`

func TestDataInitialize(ctx context.Context, dbpool *pgxpool.Pool) {
	insertProductResult, insertProductError := dbpool.Exec(ctx, INSERT_PRODUCTS)
	if insertProductError != nil {
		log.Error(insertProductError)
	} else {
		log.Info(fmt.Sprintf("Products data created with %d rows", insertProductResult.RowsAffected()))
	}
}
