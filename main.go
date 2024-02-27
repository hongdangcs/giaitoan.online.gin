package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var router *gin.Engine

func main() {

	// dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir, _ := os.Getwd()
	router = gin.Default()

	router.Static("/assets", dir+"/assets")

	router.LoadHTMLGlob(dir + "/templates/*")

	router.GET("/", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)

	})

	router.Run()

}
