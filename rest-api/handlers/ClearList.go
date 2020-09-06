package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest/model"
)

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
