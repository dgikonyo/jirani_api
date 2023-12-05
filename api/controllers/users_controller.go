package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"jirani-api/api/models"
)

func (server *Server) createUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("registration")
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusCreated, userCreated)
}

func (server *Server) getUsers(ctx *gin.Context) {
	user := models.User{}
	users, err := user.FindAllUsers(server.DB)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, users)
}

func (server *Server) getUser(ctx *gin.Context) {

}

func (server *Server) updateUser(ctx *gin.Context) {
	
}
