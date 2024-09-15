package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onurkantarci/shopwise/controller"
)

func DealRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/deals", controller.ListDeals)
}

func StoreRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/stores", controller.ListStores)
}
