package dbAPI

import (
	"fmt"
	"reflect"
)

type Commodity struct {
	Code    int    `gorm:"-" `
	Status  string `gorm:"-" `
	ID      int    `gorm:"type(11):Int;not null" form:"id"`
	Title   string `gorm:"type:String;comment:標題;not null " form:"title"`
	Cate_id int    `gorm:"type:Int(11)" form:"cate_id"`
	Amount  int    `gorm:"type:Int(11)" form:"amount"`
	Price   int    `gorm:"type:Int(10)" form:"price"`
}

func (p DBinfo) NewCommodity(c Commodity) string {
	if !p.CheckCommodity_IDisExist(c) {
		p.DB.Table(p.TableName).Create(c)

		fmt.Println("新增成功")
		t := reflect.TypeOf(c)
		v := reflect.ValueOf(c)
		for i := 0; i < t.NumField(); i++ {
			fmt.Println(t.Field(i).Name, ":", v.Field(i))
		}
		return "Success"
	}

	return "Failed"
}

func (p DBinfo) CheckCommodity_IDisExist(c Commodity) bool {
	var m Commodity
	p.DB.Table(p.TableName).Where("id = ?", c.ID).Find(&m)
	if m.ID == c.ID {
		fmt.Println("此商品ID已存在")
		return true
	} else if c.ID == 0 {
		fmt.Println("商品ID不得為零")
	}
	return false
}

func (p DBinfo) UpdateCommodity_byID(c Commodity) {
	p.DB = p.DB.Table(p.TableName).Where("id = ?", c.ID)
	var m Commodity
	p.DB.Find(&m)
	if m == c {
		fmt.Println("資訊相同")
	} else if !p.CheckCommodity_IDisExist(c) {
		fmt.Println("更新資料")
		p.DB.Save(&c)

	} else {
		fmt.Println("無法更新")
	}
}
