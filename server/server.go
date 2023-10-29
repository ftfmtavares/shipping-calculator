package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ftfmtavares/shipping-calculator/database"
	"github.com/gorilla/mux"
)

var productsDB database.Products

func NewServer(port int, products database.Products) http.Server {
	productsDB = products

	router := mux.NewRouter()

	router.HandleFunc("/product/{pid}/calculation", calculateShipping).Methods(http.MethodOptions, http.MethodGet)
	router.HandleFunc("/product/{pid}/packsizes", listPackSizes).Methods(http.MethodOptions, http.MethodGet)
	router.HandleFunc("/product/{pid}/packsizes/{size}", addPackSize).Methods(http.MethodOptions, http.MethodPost)
	router.HandleFunc("/product/{pid}/packsizes/{size}", deletePackSize).Methods(http.MethodOptions, http.MethodDelete)

	return http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
}
