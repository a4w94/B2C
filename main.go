package main

import (
	"b2c/dbAPI"
)

func main() {
	var tmp dbAPI.DBinfo
	tmp.DB = dbAPI.Connect()
	tmp.TableName = "User"
	tmp.Table = dbAPI.User{}
	//tmp.CreatTable()
	tmp.CheckTableExist()
}
