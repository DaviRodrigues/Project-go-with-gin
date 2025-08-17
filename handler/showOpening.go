package handler

import (
	"fmt"
	"net/http"

	"github.com/DaviRodrigues/Project-go-with-gin/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Get a job opening by ID
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} OpeningResponse
// @Failure 400 {object} ErrorReponse
// @Failure 404 {object} ErrorReponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest,
			errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("opening not found: %s", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
