package server

import (
	"encoding/json"
	"net/http"
)

type packSizesResponseBody struct {
	PackSizes []int `json:"packSizes"`
}

func listPackSizes(w http.ResponseWriter, r *http.Request) {
	productID, valid := validatePidVar(w, r)
	if !valid {
		return
	}

	packSizes, err := productsDB.GetPackSizes(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(packSizesResponseBody{
		PackSizes: packSizes,
	})
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func addPackSize(w http.ResponseWriter, r *http.Request) {
	productID, valid := validatePidVar(w, r)
	if !valid {
		return
	}

	size, valid := validateSizeVar(w, r)
	if !valid {
		return
	}

	err := productsDB.AddPackSize(productID, size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deletePackSize(w http.ResponseWriter, r *http.Request) {
	productID, valid := validatePidVar(w, r)
	if !valid {
		return
	}

	size, valid := validateSizeVar(w, r)
	if !valid {
		return
	}

	err := productsDB.DeletePackSize(productID, size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
