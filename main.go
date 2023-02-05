package main

import (
	"fmt"
	raspberryPiGpio "led-controller-go-pi/raspberryPiGpio"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world!")
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"status": "OK"})
	})

	r.POST("/action", raspberryPiGpio.HandleRequest)
	r.Run()
}
