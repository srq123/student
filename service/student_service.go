package service

import (
	"student/dao"
	"student/datamodels"
)

type StudentService interface {
	InsertStudent(stu datamodels.Student) error
	ListStudent(stu datamodels.Student) datamodels.Student
	DeleteStudent(id int) error
	UpdateStudentById(student datamodels.Student)
	//GetStudentById(id int)
	//GetStudentByName(name string)
}
type studentService struct {
	studentDao *dao.StudentDao
}

func NewApplicationService() *studentService {
	return &studentService{
		studentDao: dao.NewApplicationDao(),
	}
}
func (s *studentService) InsertStudent(stu datamodels.Student) error {
	err := s.studentDao.AddStudent(stu.Id, stu.UserName, stu.StudentName, stu.Sex, stu.SysCreated, stu.SysUpdated)
	if err != nil {
		return err
	}
	return nil
}
func (s *studentService) DeleteStudent(id int) error {
	err := s.studentDao.DeleteStudent(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *studentService) UpdateStudentById(student datamodels.Student) {
	s.studentDao.UpdateStudent(student.UserName, student.StudentName, student.Sex, student.SysCreated, student.SysUpdated)
}

//func GetStudentById(id int) {
//
//}
//func GetStudentByName(name string) {
//
//}
func (s *studentService) ListStudent(stu datamodels.Student) datamodels.Student {
	stuData := s.studentDao.GetStudent(stu)
	return stuData
}



