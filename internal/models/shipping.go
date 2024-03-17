package models

type Ship struct {
	Shipment Shipment `json:"shipment"`
}
type Shipment struct {
	Fromaddr Fromaddress `json:"from_address"`
	Toaddr   Toaddress   `json:"to_address"`
	Parcel   Parcel      `json:"parcel"`
	Message  string      `json:"message"`
}

type Fromaddress struct {
	Pincode string `json:"zip"`
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type Toaddress struct {
	Pincode string `json:"zip"`
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
}

type Parcel struct {
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
	Length float64 `json:"length"`
}

type ShipmentResponse struct {
	Name string
	Rate float64
}
