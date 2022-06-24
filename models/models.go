package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/iamzhiyudong/go-gin-example/pkg/setting"
)

var db *gorm.DB

// 定义json和gorm的类型映射
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

	// 获取数据库配置
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		// TODO fatal(2) 是什么意思
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	// 连接数据库
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	// 实现表名的添加前缀方法
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	// gorm默认使用复数映射，go代码的单数、复数struct形式都匹配到复数表中：创建表、添加数据时都是如此。
	// 指定了db.SingularTable(true)之后，进行严格匹配。
	db.SingularTable(true)
	db.LogMode(true)             // 开启详细日志模式
	db.DB().SetMaxIdleConns(10)  // 设置空闲连接数
	db.DB().SetMaxOpenConns(100) // 对打开的连接数设置最大值
}

func CloseDB() {
	defer db.Close()
}
