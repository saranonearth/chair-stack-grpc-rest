package handlers

import (
	"encoding/json"
	"net/http"
	"rest/model"
)

// GetAllGarments - [Gets all garment present in thel ist]
func GetAllGarments(w http.ResponseWriter, r *http.Request) {
	list := model.GetAllItems()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	enableCors(&w)
	json.NewEncoder(w).Encode(list)

}
