package controllers

import "jirani-api/api/middlewares"

func (server *Server) initializeRoutes() {
	//Home Route
	server.Router.HandleFunc("/", middlewares.setMiddlewareJSON(server.home)).Methods("GET")

	//login route
	server.Router.HandlerFunc("/users",middlewares.setMiddlewareJSON(server.login)).Methods("POST")

	//Users routes
	server.Router.HandleFunc("/users", middlewares.setMiddlewareJSON(server.createUser)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.setMiddlewareJSON(server.getUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.setMiddlewareJSON(server.getUser)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.setMiddlewareJSON(middlewares.setMiddlewareAuthentication(server.updateUser))).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middlewares.setMiddlewareAuthentication(server.deleteUser)).Methods("DELETE")
}