package app

import (
	"cart/internal/confi"
	"cart/internal/controller"
	"cart/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func RunURL(app *confi.AppConfig) {
	log.Println("Initialise all handlers")
	cartController := controller.NewController(app)

	router.GET("/ping", middleware.BasicAuth(app), func(ctx *gin.Context) {
		ctx.JSON(200, "ping")
	})

	api := router.Group("/v1")
	api.Use(middleware.BasicAuth(app))
	api.POST("/addtocart", cartController.AddtoCart)
	api.GET("/getallcart", cartController.GetAllCart)
	api.GET("/getbyid/:id", cartController.GetProductById)

}
