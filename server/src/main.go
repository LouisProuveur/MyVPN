package main

import (
	"MyVPN/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	err := routes.CreateRoutes(router)
	if err != nil {
		log.Fatal(err)
	}
}
