package controllers

func (server *Server) initializeRoutes() {

	api := server.Router.Group("/api/v1")
	{
		api.GET("/")
		// 		api.POST("/login", server.login)

		userRoute := api.Group("/users")
		{
			userRoute.POST("/create", server.createUser)
			userRoute.GET("/show", server.getUsers)
			userRoute.GET("/{id}", server.getUser)
			userRoute.PUT("/{id}", server.updateUser)
			//userRoute.DELETE("/{id}", server.deleteUser)
		}
	}
}
