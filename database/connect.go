package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"student/datamodels"
)

func GetEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost)/student?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接错误！", err)
	}
	err = engine.Sync(new(datamodels.Student))
	if err != nil {
		log.Fatal("数据表同步失败:", err)
	}
	return engine
}

