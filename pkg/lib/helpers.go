package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"schneider.vip/problem"
)

var Validate = validator.New()

func WriteNoContent(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, nil)
}

func NoRoute(ctx *gin.Context) {
	problem.New(
		problem.Title("Not Found"),
		problem.Type("errors:http/not-found"),
		problem.Status(http.StatusNotFound),
	).WriteTo(ctx.Writer)
}
