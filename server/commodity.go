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

// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   int  int  true  "ID"
// @Success      200  string string 成功後返回的值
// @Router       /newcommodity [post]
func NewCommodity(c *gin.Context) {
	var tmp dbAPI.Commodity
	fmt.Println(c.PostForm("id"))

	c.Bind(&tmp)
	fmt.Println(&tmp)
	tmp = middleware.NewCommodity(tmp)
	c.JSON(http.StatusOK, tmp)
}
