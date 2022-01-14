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
	} else {
		info.Code = 400
	}
	return info
}

func UpdateCommodity_byID(info dbAPI.Commodity) {

}
