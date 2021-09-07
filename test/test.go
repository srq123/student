package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"log"
)
type Userr struct {
	Id          int    ` xorm:"pk"`
	UserName    string ` xorm:"DEFAULT '' varchar(45)"`
	SysCreated  string `xorm:"not null default 0 comment('创建时间') varchar(45)"`
	SysUpdated  string `xorm:"not null default 0 comment('修改时间') varchar(45)"`
}
func NewXorm() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接错误！", err)
	}
	err = engine.Sync(new(Userr))
	if err != nil {
		log.Fatal("数据表同步失败:", err)
	}
	return engine
}

//Add 增加
func AddStudent(ret Userr) {
	engine := NewXorm()
	fmt.Printf("请输入需要所需插入的用户信息：(依次输入id，姓名，密码，创建时间，修改时间)\n")
	var (
		id       int64
		username string
		pwd      string
		syCreate int
		syUpdate int
	)
	_, err := fmt.Scanln(&id, &username, &pwd, &syCreate, &syUpdate)
	if err != nil {
		return
	}

	ret.Id = id
	ret.Username = username
	ret.Pwd = pwd
	ret.SysCreated = syCreate
	ret.SysUpdated = syUpdate
	insert, err := engine.Insert(&ret)
	if err != nil {
		return
	}
	log.Println(insert)
}

//Dele 删除
func DeleteStudent(ret base.Userr) {
	//engine,_ := xorm.NewEngine("mysql", "root:123456@/testrun?charset=utf8")
	engine := NewXorm()
	var id int64
	fmt.Printf("请输入需要删除的用户id：")
	_, err := fmt.Scanln(&id)
	if err != nil {
		return
	}
	ret.Id = id
	del, err := engine.Delete(&ret)
	if err != nil {
		return
	}
	log.Println(del)
	fmt.Printf("该用户已删除成功！\n")
}

//Update 更新
func UpdateStudent(ret base.Userr) bool {
	//engine,_ := xorm.NewEngine("mysql", "root:123456@/testrun?charset=utf8")
	engine := NewXorm()
	var id int64
	fmt.Printf("请输入需要删除对象的id：")
	_, err := fmt.Scanln(&id)
	if err != nil {
		return false
	}
	var (
		username string
		pwd      string
		syCreate int
		syUpdate int
	)
	fmt.Println("请输入修改信息：（依次输入姓名，密码，创建时间，修改时间）")
	scan, err := fmt.Scanln(&username, &pwd, &syCreate, &syUpdate)
	if err != nil {
		return false
	}
	log.Println(scan)
	ret.Username = username
	ret.Pwd = pwd
	ret.SysCreated = syCreate
	ret.SysUpdated = syUpdate
	str, err := engine.ID(id).Update(&ret)
	if err != nil {
		log.Fatal("update error")
	}
	if str == 0 {
		return false
	}
	return true
}

//查询
func SearchStudent(ret base.Userr) base.Userr {
	engine := NewXorm()
	str, err := engine.Count(ret)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("总人数：", str)
	app := iris.New()
	users := make(map[int64]base.Userr)
	err = engine.Find(&users)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("id\t姓名\t\t密码\t\t创建时间\t更新时间\n")
		for _, v := range users {
			fmt.Println(v.Id, "\t", v.Username, "\t", v.Pwd, "\t", v.SysCreated, "\t", v.SysUpdated)
		}
	}
	app.Get("student", func(ctx iris.Context) {
		for _, v := range users {
			ctx.Writef("%d\t%s\t%s\t%d\t%d\n", v.Id, v.Username, v.Pwd, v.SysCreated, v.SysUpdated)
		}
	})
	app.Get("/student/{id}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		data := &ret
		_, err := engine.Where("id=?", id).Get(data)
		if err != nil {
			fmt.Println(err)
		}
		ctx.Writef("%d\t%s\t%s\t%d\t%d\n", data.Id, data.Username, data.Pwd, data.SysCreated, data.SysUpdated)
	})
	err = app.Listen(":8000")
	if err != nil {
		return base.Userr{}
	}

	return ret
}