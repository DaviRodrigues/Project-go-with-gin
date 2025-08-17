package handler

import (
	"net/http"

	"github.com/DaviRodrigues/Project-go-with-gin/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description Get all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningReponse
// @Failure 500 {object} ErrorReponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError,
			"error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
