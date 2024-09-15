package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onurkantarci/shopwise/config"
	"github.com/onurkantarci/shopwise/models"
)

func ListDeals(ctx *gin.Context) {
	storeID := ctx.Query("storeID")

	var deals []models.Deal

	if storeID != "" {
		config.DB.Where("store_id = ?", storeID).Find(&deals)
	} else {
		config.DB.Find(&deals)
	}

	ctx.JSON(http.StatusOK, deals)
}

func ListStores(ctx *gin.Context) {
	stores := []models.Store{}
	config.DB.Find(&stores)
	ctx.JSON(http.StatusOK, &stores)
}
