package models

type AddToDatas struct {
	Product   Product   `json:"product"`
	Toaddress Toaddress `json:"to_address"`
}
type Product struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	SKU      string  `json:"sku"`
	Weight   float64 `json:"weight"`
	Length   float64 `json:"length"`
	Height   float64 `json:"height"`
	Width    float64 `json:"width"`
}

type CartData struct {
	Product  Product          `json:"product"`
	Shipment ShipmentResponse `json:"shipment"`
}
