package main

import (
	"b2c/server"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	port      = ":5000"
	htmlroute = "template/html/*"
)

// @title Gin Swagger Demo
// @version 1.0
// @description Swagger API.
// @host localhost:5000
func main() {
	APIServerRun()

}

func APIServerRun() {
	router := InitRouter()
	url := ginSwagger.URL("http://localhost:5000/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	fmt.Println("Route:", "localhost", port)
	s := &http.Server{
		Addr:           port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.RedirectFixedPath = true

	// r.LoadHTMLGlob(htmlroute)
	// r.Static("/js", "./template/script")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//新建标签
		apiv1.POST("/newcommodity", server.NewCommodity)

	}
	return r
}
