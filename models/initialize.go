package models

//初始化数据库配置
//http://jinzhu.me/gorm/ gorm 文档
import (
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//这个是查血的初始
var Gorm *gorm.DB
var PageSize int

func init() {
	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	port, err := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")
	if nil != err {
		port = 3306
	}
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, passwd, host, port, dbname))
	db.DB().SetConnMaxLifetime(600) //TODO::这个地方有可能产生bug
	db.DB().SetMaxOpenConns(128)
	db.DB().SetMaxIdleConns(16)
	//https://github.com/jinzhu/gorm/issues/1053
	Gorm = db

	PageSize, _ = beego.AppConfig.Int("pagesize")

}