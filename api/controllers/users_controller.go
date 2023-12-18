package controllers

import (
	"encoding/json"
	"io"
	"jirani-api/api/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (server *Server) createUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		ctx.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("registration")
	if err != nil {
		ctx.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"data": userCreated})
}

func (server *Server) getUsers(ctx *gin.Context) {
	user := models.User{}
	users, err := user.FindAllUsers(server.DB)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, users)
}

func (server *Server) getUser(ctx *gin.Context) {

}

func (server *Server) updateUser(ctx *gin.Context) {

}
