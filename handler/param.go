package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func WelcomeHandler(c *gin.Context) {
	account := c.Param("account")

	c.JSON(http.StatusOK, gin.H{
		"Message": "Selamat Datang",
		"Account": account,
		"Time":    time.Now(),
	})
}

func BooksHandler(c *gin.Context) {
	price := c.Param("price")

	c.JSON(http.StatusOK, gin.H{
		"Item":  "Buku Matematika ini seharga",
		"Price": price,
		"Time":  time.Now(),
	})
}

func BookmarkHandler(c *gin.Context) {
	link := c.Param("link")

	c.JSON(http.StatusOK, gin.H{
		"Message": "File ini telah ditandai",
		"Link":    link,
		"Time":    time.Now(),
	})
}

func RetweetHandler(c *gin.Context) {
	akun := c.Param("akun")

	c.JSON(http.StatusOK, gin.H{
		"Message": "Postingan ini telah diretweet",
		"Akun":    akun,
		"Time":    time.Now(),
	})
}

func DownloadHandler(c *gin.Context) {
	size := c.Param("size")

	c.JSON(http.StatusOK, gin.H{
		"Message":   "File ini telah diunduh",
		"File Size": size,
		"Time":      time.Now(),
	})
}
