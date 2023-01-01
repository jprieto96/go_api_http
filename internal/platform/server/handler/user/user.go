package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUserHandler(name string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusCreated, "User with name: ", name, " has been created")
	}
}
