package server

import (
	"encoding/json"
	"net/http"

	"github.com/ftfmtavares/shipping-calculator/shipping"
)

type pack struct {
	PackSize int `json:"packSize"`
	Quantity int `json:"quantity"`
}

type calculationResponseBody struct {
	Order      int    `json:"order"`
	Packs      []pack `json:"packs"`
	PacksCount int    `json:"packsCount"`
	Total      int    `json:"total"`
	Excess     int    `json:"excess"`
}

func calculateShipping(w http.ResponseWriter, r *http.Request) {
	productID, valid := validatePidVar(w, r)
	if !valid {
		return
	}

	packSizes, err := productsDB.GetPackSizes(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	order, valid := validateOrderQuery(w, r)
	if !valid {
		return
	}

	packsQty, excess, packsCount := shipping.CalculateShipping(packSizes, order)

	packs := make([]pack, 0, len(packsQty))
	for i, qty := range packsQty {
		if qty > 0 {
			packs = append(packs, pack{
				PackSize: packSizes[i],
				Quantity: packsQty[i],
			})
		}
	}

	err = json.NewEncoder(w).Encode(calculationResponseBody{
		Order:      order,
		Packs:      packs,
		PacksCount: packsCount,
		Total:      order + excess,
		Excess:     excess,
	})
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
