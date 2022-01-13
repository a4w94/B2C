package dbAPI

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName     string = "root"
	Password     string = "a4w941207"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "B2C"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func Connect() *gorm.DB {
	//組合sql連線字串
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)
	//連接MySQL
	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)

	}
	//設定ConnMaxLifetime/MaxIdleConns/MaxOpenConns
	db, err1 := conn.DB()
	if err1 != nil {
		fmt.Println("get db failed:", err)

	}
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)

	return conn
}

type User struct {
	ID          int       `gorm:"type:int"`
	Password    string    `gorm:"type:string"`
	CreatTime   time.Time `gorm:"type:time"`
	LatestLogin time.Time `gorm:"type:time"`
}

type DBinfo struct {
	DB        *gorm.DB
	TableName string
	Table     interface{}
}

func (d *DBinfo) CreatTable() {
	d.DB.Table(d.TableName).Migrator().CreateTable(&d.Table)
	fmt.Printf("DataBase:%s Create Table %s\n", Database, d.TableName)

}

func (d *DBinfo) DropTable() {
	d.DB.Migrator().DropTable(d.TableName)
	fmt.Printf("DataBase:%s Delete Table %s\n", Database, d.TableName)

}

func (d *DBinfo) CheckTableExist() {
	isExist := d.DB.Migrator().HasTable(d.TableName)

	if isExist {
		var check string
		fmt.Printf("table %s is already exist", d.TableName)
		fmt.Printf("是否刪除並創建？ %s (Y/N)", d.TableName)
		fmt.Scan(&check)
		if check == "Y" {
			d.DropTable()
			d.CreatTable()
		}
	} else {
		d.CreatTable()
	}

}
