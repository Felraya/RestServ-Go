package persistence

import (
	. "internal/entities"
)

type StudentDAO interface {
	FindAll() []Student
	Find(id int) *Student
	Exists(id int) bool
	Create(entity *Student) bool
	Delete(id int) bool
	Update(entity *Student) bool
	Save(entity *Student)
}

func GetStudentDAO() StudentDAO {
	return NewStudentDAOBolt()
}
