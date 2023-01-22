package app

import (
	"github.com/eggysetiawan/go-api-gateway/config"
	"github.com/eggysetiawan/go-api-gateway/domain"
	"github.com/eggysetiawan/go-api-gateway/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	route := mux.NewRouter()

	db := config.NewDatabaseConnection()

	ah := AuthHandler{service.NewAuthService(domain.NewAuthRepositoryDb(db))}

	route.HandleFunc("/api/login", ah.Login).Methods(http.MethodPost).Name("login")
	route.HandleFunc("/api/register", ah.Register)

	log.Panic(http.ListenAndServe("localhost:9000", route))

}
