package main

import (
	"goinventorybook/app"
	"goinventorybook/auth"
	"goinventorybook/db"
	"goinventorybook/middleware"

	_ "database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine

func main()  {
	conn := db.InitDB()
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	handler := app.New(conn)

	// Home
	router.GET("/", auth.HomeHandler)

	// Login
	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", auth.LoginPostHandler)

	// Get all books
	router.GET("/books", middleware.AuthValid, handler.GetBooks)
	router.GET("/book/:id", middleware.AuthValid, handler.GetBookById)

	// Add book
	router.GET("/addBook", middleware.AuthValid, handler.AddBook)
	router.POST("/book", middleware.AuthValid, handler.PostBook)

	// Update book
	router.GET("/updateBook/:id", middleware.AuthValid, handler.UpdateBook)
	router.POST("/updateBook/:id", middleware.AuthValid, handler.PutBook)

	// Delete book
	router.POST("/deleteBook/:id", middleware.AuthValid, handler.DeleteBook)

	router.Run()
}
