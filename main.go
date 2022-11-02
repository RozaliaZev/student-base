package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
)

type Student struct {
	name string
	age int
	grade int
}

func main() {
	studentMap := make(map[string]*Student)
	var input = bufio.NewScanner(os.Stdin)
	i := 0
	for  {
		input.Scan()
    	inputInf := input.Text()
    	
		if inputInf == "" { 
			break
		}

    	var inputMasInf []string = strings.Fields(inputInf)

    	if len(inputMasInf) != 3 {
      		fmt.Println("Введите имя, возраст и курс студента.")
    	} else {
      		studentName, studentInf := stringToStruct(inputInf)
			studentMap[studentName] = studentInf
			i++
    	}
	}
	
	
	for {
		for _,value := range studentMap {
		fmt.Println(value.name, value.age, value.grade)
		}

		fmt.Println("--------------------")
		fmt.Println("Если хотите изменить данные, напишите имя студента")
		var studentName string
		input.Scan()
    	studentName = input.Text()
		if studentName != "нет" {
			fmt.Println("Существующие данные о студенте", studentName)

      		if studentMap[studentName].getInf(studentMap, studentName) != nil {
        		fmt.Println("--------------------")
				fmt.Println("Введите новые данные о возрасте и курсе студента")
				input.Scan()
    			inputInfNewAgeGrade := input.Text()
			
				newInf := studentName + " " + inputInfNewAgeGrade
				fmt.Println("проверка", newInf)
				fmt.Println(newInf)
				_, studentInfNew := stringToStruct(newInf)
				studentMap[studentName].putInf(studentInfNew.age, studentInfNew.grade)

				fmt.Println("--------------------")
				fmt.Println("Обновленный список студентов:")
            } else {
        		return
      		}	
		} else {
			return
		}
	}
}

func stringToStruct(inputInf string) (studentName string, student *Student)  {
	var inputMasInf []string = strings.Fields(inputInf)

	studentName = inputMasInf[0]

	fmt.Println(strconv.Atoi(inputMasInf[1]))

	studentAge, err := strconv.Atoi(inputMasInf[1])
	
	if err != nil {
		log.Fatal("ошибка ввода возраста студента: ", inputMasInf[0])
		
		return
	}

	studentGrade, err := strconv.Atoi(inputMasInf[2])

	if err != nil {
		log.Fatal("ошибка ввода курса студента", inputMasInf[0])
		return
	}
	
	student = &Student{inputMasInf[0], studentAge, studentGrade}
	return 
}

func (s *Student) putInf (studentAge, studentGrade int) {
	s.age =  studentAge
	s.grade = studentGrade
	return
}

func (s *Student) getInf (studentMap map[string]*Student, studentName string) (err error) {
	err = fmt.Errorf("пук пук %v нет в базе \n", studentName) //nil
	for _, value := range studentMap {
			if value.name == studentName {
			fmt.Println(s.age, s.grade)
      		return
			} else {
			err = fmt.Errorf("студента с именем %v нет в базе \n", studentName)
		}
	}
	fmt.Print(err)
	return err
}	