package storage

import (
	"test/pkg/student"
	"fmt"
)

type StudentStorage map[string]*student.Student

func NewStudentStorage() StudentStorage {
	return make(map[string]*student.Student)
}

func PrintStudentStorage(ss StudentStorage) {
	for _,value := range ss {
		fmt.Println(value.Name, value.Age, value.Grade)
		}
}