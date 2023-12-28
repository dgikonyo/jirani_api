package controllers

import (
	"github.com/gin-gonic/gin"
	"jirani-api/api/models"
	"net/http"
)

func (server *Server) createUser(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Prepare()
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"data": userCreated})
}

func (server *Server) getUsers(ctx *gin.Context) {
	user := models.User{}
	users, err := user.FindAllUsers(server.DB)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, users)
}

func (server *Server) getUser(ctx *gin.Context) {

}

func (server *Server) updateUser(ctx *gin.Context) {

}
