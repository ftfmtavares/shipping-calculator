package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const maxOrder = 100000000

func validatePidVar(w http.ResponseWriter, r *http.Request) (int, bool) {
	pidVar := mux.Vars(r)["pid"]
	convertedPid, err := strconv.Atoi(pidVar)
	if err != nil || convertedPid <= 0 {
		http.Error(w, "product id not valid", http.StatusBadRequest)
		return 0, false
	}

	return convertedPid, true
}

func validateSizeVar(w http.ResponseWriter, r *http.Request) (int, bool) {
	sizeVar := mux.Vars(r)["size"]
	convertedSize, err := strconv.Atoi(sizeVar)
	if err != nil || convertedSize <= 0 {
		http.Error(w, "size not valid", http.StatusBadRequest)
		return 0, false
	}

	return convertedSize, true
}

func validateOrderQuery(w http.ResponseWriter, r *http.Request) (int, bool) {
	orders := r.URL.Query()["order"]
	if len(orders) == 0 {
		http.Error(w, "order query parameter must be specified", http.StatusBadRequest)
		return 0, false
	}

	convertedOrder, err := strconv.Atoi(orders[0])
	if err != nil || convertedOrder < 0 {
		http.Error(w, "order query parameter not valid", http.StatusBadRequest)
		return 0, false
	}
	if convertedOrder > maxOrder {
		http.Error(w, fmt.Sprintf("order too large: maximum %d", maxOrder), http.StatusBadRequest)
		return 0, false
	}

	return convertedOrder, true
}
