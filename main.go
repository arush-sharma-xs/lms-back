package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xs-arush-0856/lms/models"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Headers", "*")
        /*
            c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
            c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
            c.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
            c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
        */
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}


func main() {
	models.ConnectDB()
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.LoadHTMLGlob("pages/*")

	// output html pages
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/adminpage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "adminpage.html", gin.H{})
	})

	r.GET("/allbooks", func(c *gin.Context) {
		c.HTML(http.StatusOK, "books.html", gin.H{})
	})

	r.GET("/readerpage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "readerpage.html", gin.H{})
	})

	r.GET("/createaccount", func(c *gin.Context) {
		c.HTML(http.StatusOK, "createaccount.html", gin.H{})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.GET("/addbook", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addbook.html", gin.H{})
	})

	// endPoint
	r.POST("/library", handleCreateLibrary) // new library creation
	r.POST("/users", addUser)               // Creating new User

	// Admin Routes
	r.POST("/auth", auth);
	r.GET("/book", getBook)
	r.POST("/book", addBook)                // Add books
	r.DELETE("/book/:isbn", removeBook)     // Delete books
	r.PATCH("/book/:isbn", updateBook)      // Update Book
	r.GET("/issuerequest", getIssueRequest) // Get Issue List
	r.GET("/search/:key/:value", searchBook)
	r.POST("/raiseissue", raiseIssue)

	r.Run(":5101")
}
