package main

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		ApiHandler{Index},
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		ApiHandler{TodoIndex},
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		ApiHandler{TodoShow},
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		ApiHandler{TodoCreate},
	},
}
