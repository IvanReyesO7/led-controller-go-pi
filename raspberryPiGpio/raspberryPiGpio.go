package raspberryPiGpio

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio"
)

var pin = rpio.Pin(10)

func HandleRequest(c *gin.Context) {
	action := c.PostForm("action")
	switch action {
	case "ON":
		if err := turnOn(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Status": "Ok"})
		return
	case "OFF":
		if err := turnOff(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Status": "Ok"})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Bad Request"})
		return
	}
}

func turnOn() error {
	fmt.Println("Its ON")
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return err
	}
	defer rpio.Close()
	pin.Output()
	pin.High()
	return nil
}

func turnOff() error {
	fmt.Println("Its OFF")
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return err
	}
	defer rpio.Close()
	pin.Output()
	pin.Low()
	return nil
}
