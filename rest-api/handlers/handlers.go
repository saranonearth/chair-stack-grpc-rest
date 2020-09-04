package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"rest/model"
)

type jsonmsg struct {
	Msg string `json:"response"`
}

type data struct {
	Data model.ItemList `json:"data"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Index function - [root /]
func Index(w http.ResponseWriter, r *http.Request) {
	nowTime := time.Now()
	msg := jsonmsg{
		Msg: "Chair Stack v1" + " " + nowTime.String(),
	}
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

// AddGarment - [Adds garment to list]
func AddGarment(w http.ResponseWriter, r *http.Request) {

	garment := r.URL.Query().Get("garment")

	list := model.AddToList(garment)

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(list)

}

// GetAllGarments - [Gets all garment present in thel ist]
func GetAllGarments(w http.ResponseWriter, r *http.Request) {
	list := model.GetAllItems()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	enableCors(&w)
	json.NewEncoder(w).Encode(list)

}

// ClearList - [clear the list]
func ClearList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CLEAR")
	model.Clear()

	data := jsonmsg{
		Msg: "Cleared",
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)

}
