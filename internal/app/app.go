package app

import (
	"cart/internal/confi"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func StartApplication(app *confi.AppConfig) {
	RunURL(app)
	log.Println("It is Listening to port", app.Server.Port)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", app.Server.Port),
		Handler:      router,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  1 * time.Minute,
	}
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Can't able to start the server", err)

	}
}
