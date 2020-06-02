package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-gin-smallproject/pkg/setting"
	"log"
)

// models 的初始化使用

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database':%v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}
	// 重写方法  // DefaultTableNameHandler default table name handler 返回表头 + 表明
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	// SingularTable默认使用单数表
	db.SingularTable(true)
	// LogMode设置日志模式，“true”表示详细日志，“false”表示没有日志，默认情况下，只会打印错误日志
	db.LogMode(true)
	// SetMaxIdleConns设置空闲连接池中的最大连接数。
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns设置到数据库的最大打开连接数。
	db.DB().SetMaxOpenConns(100)
}
func CloseDB() {
	// 延迟关闭数据库连接
	defer db.Close()
}
