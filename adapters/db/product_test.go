package db_test

import (
	"database/sql"
	"github.com/Sans-arch/fc-arquitetura-hexagonal/adapters/db"
	"github.com/Sans-arch/fc-arquitetura-hexagonal/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
    		"id" string,
    		"name" string,
    		"price" float,
    		"status" string
	);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Product Test", 0, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "Product Test", productResult.GetName())
	require.Equal(t, 25.0, productResult.GetPrice())
	require.Equal(t, "disabled", productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDb.Save(product)
	require.Equal(t, "Product Test", productResult.GetName())
	require.Equal(t, 25.0, productResult.GetPrice())
	require.Equal(t, "enabled", productResult.GetStatus())
}
