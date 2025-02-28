package main

import "github.com/julienschmidt/httprouter"

/*
Define all the routes here.
A new Route entry passed to the routes slice will be automatically
translated to a handler with the NewRouter() function
*/
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{"Index", "GET", "/", Index},
		Route{"FindJournal", "GET", "/api/journals", FindJournal},
		Route{"ShowJournal", "GET", "/api/journals/:id", ShowJournal},
		Route{"CreateJournal", "POST", "/api/journals/create", CreateJournal},
		Route{"UpdateJournal", "PUT", "/api/journals/:id/update", UpdateJournal},
		Route{"DeleteJournal", "DELETE", "/api/journals/:id/delete", DeleteJournal},
	}
	return routes
}
