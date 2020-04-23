package rest

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	//Get gin's default engine
	r := gin.Default()
	//Define a handler
	h, _ := NewHandler()
	//load homepage
	r.GET("/", h.GetMainPage)
	//get products

	return r.Run(address)
}

func RunServer() {

	log.Println("Main log....")
	RunAPI(":9090")
}
