package handler

import (
	"fmt"
	"net/http"

	"github.com/DaviRodrigues/Project-go-with-gin/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorReponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type OpeningResponse struct {
	Message string                 `json:"message"`
	Data    schemas.OpeningReponse `json:"data"`
}

type ListOpeningReponse struct {
	Message string                   `json:"message"`
	Data    []schemas.OpeningReponse `json:"data"`
}
