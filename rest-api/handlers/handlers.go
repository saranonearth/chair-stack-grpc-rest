package handlers

import (
	"encoding/json"
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
