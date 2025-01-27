package router

import (
	"create-billing-details-service/dbcontext/billingdetails"
	"create-billing-details-service/service"
	"create-billing-details-service/controller"
	"create-billing-details-service/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(config.GetCORSConfig())
	//[ginshot-routes]
	router.POST("/create", controller.NewCreateBillingDetailsController(service.NewCreateBillingDetailsServiceImpl(dbcontext.GetDBClient())).CreateBillingDetails)
	return router
}
