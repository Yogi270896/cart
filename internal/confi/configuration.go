package confi

type AppConfig struct {
	Server      server
	ShippingUrl string
}

type server struct {
	Port string
}

func NewConfig() *AppConfig {
	return &AppConfig{
		Server: server{
			Port: "8081",
		},

		ShippingUrl: "http://localhost:8080",
	}
}
