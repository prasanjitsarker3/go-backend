package cmd

import (
	"backend-golang/config"
	"backend-golang/repo"
	"backend-golang/rest"
	"backend-golang/rest/handler/product"
	"backend-golang/rest/handler/users"
)

func Serve() {
	cnf := config.GetConfig()

	productHandler := product.NewHandler(repo.NewProductRepo())
	userHandler := users.NewHandler(repo.NewUserRepo())

	server := rest.NewServer(productHandler, userHandler)
	server.StartServer(cnf)
}
