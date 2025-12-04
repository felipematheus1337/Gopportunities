package handler

import (
	"net/http"

	"github.com/felipematheus1337/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error listening openings")
		return
	}

	sendSucess(ctx, "list-openings", openings)
}
