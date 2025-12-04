package handler

import (
	"net/http"

	"github.com/felipematheus1337/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "no id provided")
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sendSucess(ctx, "show-opening", opening)
}
