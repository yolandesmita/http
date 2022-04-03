package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductInput struct {
	Toko   string
	Produk string
	Harga  int
	Kota   string
	Ukuran string
}

func PostProductHandler(c *gin.Context) {
	var produkInput ProductInput

	err := c.ShouldBindJSON(&produkInput)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"Toko":   produkInput.Toko,
		"Produk": produkInput.Produk,
		"Harga":  produkInput.Harga,
		"Kota":   produkInput.Kota,
		"Ukuran": produkInput.Ukuran,
	})
}
