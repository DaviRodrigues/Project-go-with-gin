package handler

import (
	"fmt"
	"net/http"

	"github.com/DaviRodrigues/Project-go-with-gin/schemas"
	"github.com/gin-gonic/gin"
)

func validateAndUpdateFields(opening *schemas.Opening, request UpdateOpeningRequest) {
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
}

// @BasePath /api/v1

// @Summary Update opening
// @Description Update a job opening by ID
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Param request body UpdateOpeningRequest true "Request body"
// @Success 200 {object} OpeningResponse
// @Failure 400 {object} ErrorReponse
// @Failure 404 {object} ErrorReponse
// @Failure 500 {object} ErrorReponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	var request UpdateOpeningRequest

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest,
			errParamIsRequired("id", "queryParameter").Error())
		return
	}

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

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	validateAndUpdateFields(&opening, request)

	// Save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
