package main

import (
	"webapigolang/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Query Param
	router.GET("/homepage/:account", handler.WelcomeHandler)
	router.GET("/items/:price", handler.BooksHandler)
	router.GET("/bookmark/:link", handler.BookmarkHandler)
	router.GET("/retweet/:akun", handler.RetweetHandler)
	router.GET("/download/:size", handler.DownloadHandler)

	// Query String
	router.GET("/book", handler.BookString)
	router.GET("/car", handler.CarString)
	router.GET("/doll", handler.DollString)
	router.GET("/honey", handler.HoneyString)
	router.GET("/paper", handler.PaperString)

	// Nomor 1 dan 2
	router.GET("/item", handler.BukuHandler)
	router.GET("/id/:id/produk", handler.TVHandler)

	// POST
	router.POST("/dress", handler.PostProductHandler)

	router.Run()
}
