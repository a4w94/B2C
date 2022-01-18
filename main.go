package main

import (
	"fmt"
	"net/http"

	docs "b2c/docs"
	"b2c/server"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const (
	port      = ":5000"
	htmlroute = "template/html/*"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /hello [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func main() {

	APIServerRun()

}

func APIServerRun() {
	router := InitRouter()

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
	docs.SwaggerInfo.BasePath = "/api/v1"

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
		apiv1.GET("/hello", Helloworld)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
