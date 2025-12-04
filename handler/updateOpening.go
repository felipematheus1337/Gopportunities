package handler

import (
	"net/http"

	"github.com/felipematheus1337/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	openings := schemas.Opening{}

	if err := db.First(&openings, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	openings = *updateProperties(&request, &openings)

	if err := db.Save(&openings).Error; err != nil {
		logger.Errorf("failed to save openings: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(ctx, "update-opening", openings)

}

func updateProperties(request *UpdateOpeningRequest, opening *schemas.Opening) *schemas.Opening {
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Link != "" {
		opening.Link = request.Link
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

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	return opening
}
