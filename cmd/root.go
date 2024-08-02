package cmd

import (
	"database/sql"
	dbInfra "github.com/Sans-arch/fc-arquitetura-hexagonal/adapters/db"
	"github.com/Sans-arch/fc-arquitetura-hexagonal/application"
	"github.com/spf13/cobra"
)

var db, _ = sql.Open("sqlite3", "db.sqlite")
var productDb = dbInfra.NewProductDb(db)
var productService = application.ProductService{Persistence: productDb}

var rootCmd = &cobra.Command{
	Use:   "fc-arquitetura-hexagonal",
	Short: "A brief description of your application",
	Long:  `A longer description that spans ....`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
