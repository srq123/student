package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"student/controller"
	"student/database"
	"student/datamodels"
)

func main() {
	database.GetEngine()
	var (
		student  datamodels.Student
		con     controller.StudentController
	)
	//err := con.Insert(student)
	//if err != nil {
	//	return
	//}
	fmt.Println(student)
	stu := con.ListStudent(&student)
	fmt.Println(stu)
}
