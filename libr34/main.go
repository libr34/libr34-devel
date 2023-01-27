package main

import (
	"example.com/libr34/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: false,
		AllowMethods:     []string{"GET"},
	}))

	services.AttachServices(r)

	r.Run()
}
