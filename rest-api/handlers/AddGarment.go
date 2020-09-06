package handlers

import (
	"encoding/json"
	"net/http"
	"rest/model"
)

// AddGarment - [Adds garment to list]
func AddGarment(w http.ResponseWriter, r *http.Request) {

	garment := r.URL.Query().Get("garment")

	list := model.AddToList(garment)

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(list)

}
