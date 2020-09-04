package routing

import (
	"net/http"

	"rest/handlers"
)

//Route Model
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes - collection of routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"addGarment",
		"GET",
		"/add-garment",
		handlers.AddGarment,
	},
	Route{
		"getAllGarment",
		"GET",
		"/all-garments",
		handlers.GetAllGarments,
	},
	Route{
		"clearList",
		"GET",
		"/clear",
		handlers.ClearList,
	},
}
