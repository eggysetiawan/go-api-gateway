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

	api := route.PathPrefix("/api").Subrouter()

	am := AuthMiddleware{}
	api.Use(am.AuthorizationHandler())

	api.HandleFunc("/routings", rh.indexRouting).Methods(http.MethodGet)
	api.HandleFunc("/routings", rh.storeRouting).Methods(http.MethodPost)
	api.HandleFunc("/routings/{uuid}", rh.showRouting).Methods(http.MethodGet)
	api.HandleFunc("/routings/{uuid}/edit", rh.updateRouting).Methods(http.MethodPut)
	api.HandleFunc("/routings/{uuid}/delete", rh.deleteRouting).Methods(http.MethodDelete)

	//middleware

	//example middleware
	//route.Use(func(next http.Handler) http.Handler {
	//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		//before
	//		next.ServeHTTP(w, r)
	//		//	after
	//	})
	//})
	//
	//route.Use(am.AuthorizationHandler())

	log.Panic(http.ListenAndServe("localhost:9000", route))

}
