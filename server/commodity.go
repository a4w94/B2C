package server

import (
	"b2c/dbAPI"
	"b2c/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct {
	User string `form:"user"`
}

func NewCommodity(c *gin.Context) {
	var tmp dbAPI.Commodity

	c.Bind(&tmp)
	fmt.Println(&tmp)
	tmp = middleware.NewCommodity(tmp)
	c.JSON(http.StatusOK, tmp)
}
