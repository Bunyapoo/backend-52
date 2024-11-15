package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET METHOD",
		})
	})
	//POST ME THOD
	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST METHOD",
		})
	})
	//PUT ME THOD
	r.PUT("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT METHOD",
		})
	}) //DELETE ME THOD
	r.DELETE("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE METHOD",
		})
	})
	//---------------------------------------------------------------------
	//---------------------------CUSTOMER METHOD---------------------------
	//GET METHOD
	r.GET("/CUSTOMER", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET CUSTOMER METHOD",
		})
	})
	//POST ME THOD
	r.POST("/CUSTOMER", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST CUSTOMER METHOD",
		})
	})
	//PUT ME THOD
	r.PUT("/CUSTOMER", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT CUSTOMER METHOD",
		})
	})	//DELETE ME THOD
	r.DELETE("/CUSTOMER", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE CUSTOMER METHOD",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
