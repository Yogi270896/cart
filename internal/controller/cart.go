package controller

import (
	"cart/internal/confi"
	"cart/internal/models"
	"cart/internal/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	AppConfig *confi.AppConfig
	Service   service.CartService
}

func NewController(app *confi.AppConfig) *CartController {
	return &CartController{
		AppConfig: app,
		Service:   *service.NewService(app),
	}
}
func (cc *CartController) GetAllCart(ctx *gin.Context) {
	m := "Contorller"
	log.Println(m, "GetAllCart", "Start")
	res, err := cc.Service.GetAllCart()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		log.Println(m, "GetAllCart", "End")
		return
	}
	ctx.JSON(http.StatusOK, res)
	log.Println(m, "GetAllCart", "End")
}
func (cc *CartController) GetProductById(ctx *gin.Context) {
	m := "Contorller"
	f := "GetProductById"
	log.Println(m, f, "Start")
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, "product id cannot be nil")
	}
	res, err := cc.Service.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusNoContent, err)
		log.Println(m, f, "End")
		return
	}
	ctx.JSON(http.StatusFound, res)
	log.Println(m, f, "End")
}
func (cc *CartController) AddtoCart(ctx *gin.Context) {
	m := "Contorller"
	log.Println(m, "Add To Cart", "Start")

	var reqModel models.AddToDatas
	if err := ctx.ShouldBindJSON(&reqModel); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	fmt.Println(reqModel)
	res, errs := cc.Service.AddToCart(reqModel)
	if errs != nil {
		ctx.JSON(http.StatusBadRequest, errs.Error())
		log.Println(m, "Add To Cart", "End")
		return
	}
	ctx.JSON(http.StatusAccepted, res)
	log.Println(m, "Add To Cart", "End")
}
