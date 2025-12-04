package handler

import (
	"fmt"
	"net/http"

	"github.com/felipematheus1337/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id); err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(ctx, "delete-opening", opening)
}
