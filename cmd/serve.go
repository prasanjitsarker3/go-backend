package cmd

import (
	"backend-golang/config"
	"backend-golang/db"
	"backend-golang/repo"
	"backend-golang/rest"
	"backend-golang/rest/handler/product"
	userHandlers "backend-golang/rest/handler/users"
	userservice "backend-golang/users"
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
	
	// Domains 
	userService := userservice.NewService(repo.NewUserRepo(dbConn))

    // Initialize repositories and handlers
	productHandler := product.NewHandler(repo.NewProductRepo())
	userHandler := userHandlers.NewHandler(userService)
    

	server := rest.NewServer(productHandler, userHandler)
	server.StartServer(cnf)
}
