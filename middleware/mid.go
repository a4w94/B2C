package middleware

import (
	"b2c/dbAPI"
)

func NewCommodity(info dbAPI.Commodity) dbAPI.Commodity {
	var tmp = dbAPI.DBinfo{
		DB:        dbAPI.Connect(),
		TableName: "Shop",
	}

	check := tmp.NewCommodity(info)
	if check == "Success" {
		info.Code = 200
		info.Status = "商品新增成功"
	} else {
		info.Code = 400
		info.Status = "商品新增失敗"

	}
	return info
}

func UpdateCommodity_byID(info dbAPI.Commodity) {

}
