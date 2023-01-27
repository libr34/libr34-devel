package main

import (
	"example.com/libr34/services"
	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	r := gin.Default()

	services.AttachServices(r)

	r.Run()
}
