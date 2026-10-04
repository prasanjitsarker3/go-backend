package cmd

import (
	"backend-golang/config"
	"backend-golang/db"
	"backend-golang/repo"
	"backend-golang/rest"
	"backend-golang/rest/handler/product"
	"backend-golang/rest/handler/users"
	"context"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
	ctx := context.Background()

	dbConn, err := db.NewConnection(ctx, cnf.DB)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	fmt.Println("PostgreSQL connected successfully!")

	productHandler := product.NewHandler(repo.NewProductRepo())
	userHandler := users.NewHandler(repo.NewUserRepo(dbConn))

	server := rest.NewServer(productHandler, userHandler)
	server.StartServer(cnf)
}
