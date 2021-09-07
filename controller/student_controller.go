package controller

import (
	"fmt"
	"student/datamodels"
	"student/service"
	"time"
)

type StudentController struct {
	Service service.StudentService
}

func (s *StudentController) Insert(data datamodels.Student) error {
	_, err := fmt.Scanln(&data.Id, &data.UserName, &data.StudentName, &data.Sex)
	if err != nil {
		return err
	}
	data.SysCreated = time.Now().String()
	data.SysUpdated = time.Now().String()
	err = s.Service.InsertStudent(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *StudentController) ListStudent(data *datamodels.Student) datamodels.Student {

	stu := s.Service.ListStudent(*data)
	return stu
}

func (s *StudentController) UpdateStudent(student datamodels.Student) {
	_, err := fmt.Scanln(&student.Id, &student.UserName, &student.StudentName, &student.Sex, &student.SysUpdated, student.SysCreated)
	if err != nil {
		return
	}
	s.Service.UpdateStudentById(student)
}

