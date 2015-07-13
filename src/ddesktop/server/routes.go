package server

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetIndex",
		"GET",
		"/",
		GetIndex,
	},
	Route{
		"GetFavicon",
		"GET",
		"/favicon.ico",
		GetFavicon,
	},
}
