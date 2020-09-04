package handlers

import (
	"fmt"
	"net/http"
)

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/")
}
