package app

import (
	"github.com/eggysetiawan/go-api-gateway/config"
	"github.com/eggysetiawan/go-api-gateway/domain"
	"github.com/eggysetiawan/go-api-gateway/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Start() {
	route := mux.NewRouter()

	db := config.NewDatabaseConnection()

	ah := AuthHandler{service.NewAuthService(domain.NewAuthRepositoryDb(db))}

	route.HandleFunc("/api/login", ah.Login).Methods(http.MethodPost).Name("login")
	route.HandleFunc("/api/register", ah.Register)

	client := &http.Client{Timeout: 10 * time.Second}
	clientAddress := "http://localhost:8000"

	rh := RoutingHandler{service.NewApiRoutingService(domain.NewRoutingRepositoryApi(client, clientAddress))}

	route.HandleFunc("/api/routings", rh.indexRouting).Methods(http.MethodGet)
	route.HandleFunc("/api/routings", rh.storeRouting).Methods(http.MethodPost)
	route.HandleFunc("/api/routings/{uuid}", rh.showRouting).Methods(http.MethodGet)
	route.HandleFunc("/api/routings/{uuid}/edit", rh.updateRouting).Methods(http.MethodPut)
	route.HandleFunc("/api/routings/{uuid}/delete", rh.deleteRouting).Methods(http.MethodDelete)

	log.Panic(http.ListenAndServe("localhost:9000", route))

}
