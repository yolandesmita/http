package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BookString(c *gin.Context) {
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"Message": "Buku ini seharga",
		"Harga":   price,
		"Time":    time.Now(),
	})
}

func CarString(c *gin.Context) {
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"Message": "Mobil ini seharga",
		"Harga":   price,
		"Time":    time.Now(),
	})
}

func DollString(c *gin.Context) {
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"Message": "Boneka ini seharga",
		"Harga":   price,
		"Time":    time.Now(),
	})
}

func HoneyString(c *gin.Context) {
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"Message": "Madu ini seharga",
		"Harga":   price,
		"Time":    time.Now(),
	})
}

func PaperString(c *gin.Context) {
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"Message": "Kertas ini seharga",
		"Harga":   price,
		"Time":    time.Now(),
	})
}
