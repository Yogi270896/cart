package testing

import (
	"cart/internal/confi"
	"cart/internal/models"
	"cart/internal/service"
	"fmt"
	"testing"
)

func TestGetrate(t *testing.T) {
	app := confi.NewConfig()

	req := models.AddToDatas{
		Product: models.Product{
			Name:     "Shoe",
			Quantity: 1,
			Price:    10,
			Height:   2,
			Length:   2,
			Weight:   2,
			Width:    2,
		},
		Toaddress: models.Toaddress{
			Pincode: "90005",
			City:    "California",
			State:   "CA",
			Country: "US",
		},
	}

	s := service.NewService(app)
	res, err := s.GetShippingRate(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}
