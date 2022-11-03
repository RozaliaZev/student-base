package student

import(
	"fmt"
)

type Student struct {
	Name string
	Age int
	Grade int
}

func (s *Student) PutInf (studentAge, studentGrade int) {
	s.Age =  studentAge
	s.Grade = studentGrade
	return
}

func (s *Student) GetInf (studentMap map[string]*Student, studentName string) (err error) {
	for _, value := range studentMap {
			if value.Name == studentName {
			fmt.Println(s.Age, s.Grade)
      		return
			} else {
			err = fmt.Errorf("студента с именем %v нет в базе \n", studentName)
		}
	}
	fmt.Print(err)
	return err
}