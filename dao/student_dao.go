package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"student/database"
	"student/datamodels"
)

type StudentDao struct {
	engine *xorm.Engine
}

func NewApplicationDao() *StudentDao {
	return &StudentDao{engine: database.GetEngine()}
}
func (stu *StudentDao) AddStudent(id int, username, studentName, sex, createTime, updateTime string) error {
	student := new(datamodels.Student)
	student.Id = id
	student.UserName = username
	student.StudentName = studentName
	student.Sex = sex
	student.SysCreated = createTime
	student.SysUpdated = updateTime
	_, err := stu.engine.Insert(student)
	if err != nil {
		log.Println("insert message fail!")
		return err
	}
	return nil
}

func (stu *StudentDao) DeleteStudent(id int) error {
	var student datamodels.Student
	student.Id = id
	_, err := stu.engine.Delete(student)
	if err != nil {
		return err
	}
	return nil
}
func (stu *StudentDao) UpdateStudent(username, studentName, sex, createTime, updateTime string) {
	var student datamodels.Student
	student.UserName = username
	student.StudentName = studentName
	student.Sex = sex
	student.SysCreated = createTime
	student.SysUpdated = updateTime
	_, err := stu.engine.Update(&student)
	if err != nil {
		return
	}
}

func (stu *StudentDao) GetStudent(student datamodels.Student) datamodels.Student {
	err := stu.engine.Find(&student)
	stu.engine.ShowSQL()
	if err != nil {
		return student
	}
	return student
}
func (stu *StudentDao) GetStudentById(id int) datamodels.Student {
	var student datamodels.Student
	_, err := stu.engine.ID(id).Get(&student)
	if err != nil {
		return datamodels.Student{}
	}
	return student
}

func (stu *StudentDao) GetStudentByName(name string) datamodels.Student {
	var student datamodels.Student
	_, err := stu.engine.Where("studentName=?", name).Get(&student)
	if err != nil {
		return datamodels.Student{}
	}
	return student
}
