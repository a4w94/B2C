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

// @Summary      新增
// @Description  新增商品
// @Tags         New
// @Accept  	 application/x-www-form-urlencoded
// @Produce      json
// @Param        id  formData  int  true  "id"
// @Param        cate_id  formData  int  true  "Cate_ID"
// @Param        amount  formData  int  true  "數量"
// @Param        price formData  int  true  "Price"
// @Param        title formData  string  false  "title"
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
