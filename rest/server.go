package rest

import (
	"backend-golang/config"
	"backend-golang/middleware"
	"backend-golang/rest/handler/product"
	"backend-golang/rest/handler/users"
	"fmt"
	"net/http"
)

type Server struct {
	productHandler * product.Handler
    userHandler * users.Handler

}


func NewServer (
	productHandler * product.Handler,
	userHandler * users.Handler ,
	) * Server{
	return & Server {
		productHandler: productHandler,
		userHandler: userHandler,
	}
}



func ( server * Server) StartServer(cnf *config.Config) {
	mux := http.NewServeMux()
	routers(mux)

	handler := middleware.Chain(
		mux,
		middleware.CORS,
		middleware.Logger,
	)
	server.productHandler.ProductRegisterRoutes(mux)
	server.userHandler.UserRegisterRoutes(mux)

	err := http.ListenAndServe(fmt.Sprintf(":%d", cnf.HttpPort), handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

