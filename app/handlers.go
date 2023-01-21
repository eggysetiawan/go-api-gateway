package app

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	fmt.Println("Server starting")

	mux := mux.NewRouter()

	log.Panic(http.ListenAndServe("localhost:8000", mux))

}
