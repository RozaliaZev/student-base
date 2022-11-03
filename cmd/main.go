package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
	"test/pkg/student"
)


func main() {
	studentMap := make(map[string]*student.Student)
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
		fmt.Println(value.Name, value.Age, value.Grade)
		}

		fmt.Println("--------------------")
		fmt.Println("Если хотите изменить данные, напишите имя студента")
		var studentName string
		input.Scan()
    	studentName = input.Text()
		if studentName != "нет" {
			fmt.Println("Существующие данные о студенте", studentName)

      		if studentMap[studentName].GetInf(studentMap, studentName) == nil {
        		fmt.Println("--------------------")
				fmt.Println("Введите новые данные о возрасте и курсе студента")
				input.Scan()
    			inputInfNewAgeGrade := input.Text()
			
				newInf := studentName + " " + inputInfNewAgeGrade
				_, studentInfNew := stringToStruct(newInf)
				studentMap[studentName].PutInf(studentInfNew.Age, studentInfNew.Grade)

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

func stringToStruct(inputInf string) (studentName string, newStudent *student.Student)  {
	var inputMasInf []string = strings.Fields(inputInf)

	studentName = inputMasInf[0]

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
	
	newStudent = &student.Student{Name: inputMasInf[0], Age: studentAge, Grade: studentGrade}
	return 
}

	