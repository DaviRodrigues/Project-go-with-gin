package handler

import (
	"net/http"

	"github.com/DaviRodrigues/Project-go-with-gin/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} OpeningResponse
// @Failure 400 {object} ErrorReponse
// @Failure 500 {object} ErrorReponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	var request CreateOpeningRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("bad formmated json: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var openingReq OpeningRequest = &request

	if err := openingReq.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
