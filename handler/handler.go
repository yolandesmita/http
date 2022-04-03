package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Nomor 1
func BukuHandler(c *gin.Context) {
	judul := c.Query("judul")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"Message":   "Item Buku",
		"Item Name": judul,
		"Time":      time.Now(),
		"Price":     price,
	})
}

// Nomor 2
func TVHandler(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"id":           id,
		"Message":      "Item Name",
		"Product Name": name,
		"Time":         time.Now(),
		"Price":        price,
	})
}
