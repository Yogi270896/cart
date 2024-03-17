package service

import (
	"cart/internal/confi"
	"cart/internal/helper"
	"cart/internal/models"
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

type CartService struct {
	AppConfig *confi.AppConfig
	CartData  []models.CartData
}

func NewService(app *confi.AppConfig) *CartService {
	return &CartService{
		AppConfig: app,
	}
}

func (c *CartService) AddToCart(productData models.AddToDatas) (models.CartData, error) {
	m := "Service"
	f := "AddToCart"
	var cartData models.CartData
	log.Println(m, f, "Start")
	shipRate, err := c.GetShippingRate(productData)
	if err != nil {
		return cartData, err
	}
	cartData.Product = productData.Product
	cartData.Shipment = shipRate
	c.CartData = append(c.CartData, cartData)

	log.Println(m, f, "End")
	return cartData, nil

}

func (c *CartService) GetAllCart() ([]models.CartData, error) {
	m := "Service"
	f := "GetAllCart"
	log.Println(m, f, "Start")

	log.Println(m, f, "End")
	return c.CartData, nil

}

func (c *CartService) GetProductById(id string) (models.CartData, error) {
	m := "Service"
	f := "GetProductById"
	log.Println(m, f, "Start")
	var product models.CartData
	if len(c.CartData) == 0 {
		return product, nil
	}
	for _, data := range c.CartData {
		if val := data.Product.SKU; val == id {
			return data, nil
		}
	}
	log.Println(m, f, "End")

	return product, errors.New("Product Not Found")

}

func (c *CartService) GetShippingRate(productData models.AddToDatas) (models.ShipmentResponse, error) {
	m := "Service"
	log.Println(m, "GetRates", "Start")
	var shipReq models.Shipment
	var shipmentResponse models.ShipmentResponse
	var response map[string]interface{}
	Fromaddr := models.Fromaddress{
		Pincode: "90001",
		City:    "California",
		State:   "CA",
		Country: "US",
	}

	shipReq.Fromaddr = Fromaddr
	shipReq.Toaddr = productData.Toaddress
	shipReq.Parcel.Height = productData.Product.Height
	shipReq.Parcel.Width = productData.Product.Height
	shipReq.Parcel.Length = productData.Product.Height
	shipReq.Parcel.Weight = productData.Product.Weight

	url := c.AppConfig.ShippingUrl + "/v1/getrates"
	reqBody, _ := json.Marshal(shipReq)

	log.Println(m, "GetRates", "URL", url)
	log.Println(m, "GetRates", "BODY", string(reqBody))

	res, errs := helper.Send("GET", url, shipReq, "yogi", "GetRates")
	if errs != nil {
		log.Println(m, "GetRates", errs)
		return shipmentResponse, errs
	}
	json.Unmarshal(res, &response)
	//log.Println(m, "GetRates", "Response", response)

	shippingCarrier := response["rates"].([]interface{})

	for _, r := range shippingCarrier {
		carrierData := r.(map[string]interface{})
		if carrier := carrierData["carrier"].(string); carrier == "USPS" {
			if service := carrierData["service"].(string); service == "Priority" {
				rates := carrierData["rate"].(string) // We are receving the value in string so we need to convert inteface into string
				rate, _ := strconv.ParseFloat(rates, 64)
				shipmentResponse.Name = service
				shipmentResponse.Rate = rate
			}

		}
	}
	log.Println(m, "GetRates", "End")
	return shipmentResponse, nil
}
