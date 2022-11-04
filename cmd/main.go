package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"test/pkg/storage"
	"test/pkg/student"
)


func main() {
	fmt.Println("Создание базы студентов")
	fmt.Println("Введите данные о студенте, его имя, возраст и курс обучения.")
	studentMap := storage.NewStudentStorage()
	var input = bufio.NewScanner(os.Stdin)
	i := 0
	for  {
		input.Scan()
    	inputInf := input.Text()
    	
		if inputInf == "" { 
			break
		}

		studentName, studentInf, inputErr := stringToStruct(inputInf)
    	
		if inputErr != nil {
      		fmt.Println(inputErr)
    	} else {
      		//studentName, studentInf, _ := stringToStruct(inputInf)
			studentMap[studentName] = studentInf
			i++
    	}
	}
	
	
	for {
		storage.PrintStudentStorage(studentMap)

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
				_, studentInfNew, inputErr := stringToStruct(newInf)

				if inputErr != nil {
					fmt.Println(inputErr)
				} else {
					studentMap[studentName].PutInf(studentInfNew.Age, studentInfNew.Grade)
					fmt.Println("--------------------")
					fmt.Println("Обновленный список студентов:")
				}	
            } else {
        		continue
      		}	
		} else {
			return
		}
	}
}

func stringToStruct(inputInf string) (studentName string, newStudent *student.Student, err error)  {
	var inputMasInf []string = strings.Fields(inputInf)

	if len(inputMasInf) != 3 {
		err = errors.New("Введено неверное колическо параметров. Введите имя, возраст и курс студента.")
		return
	} 

	studentName = inputMasInf[0]

	studentAge, err := strconv.Atoi(inputMasInf[1])
	
	if err != nil || studentAge <= 0 {
		err = errors.New("ошибка ввода возраста студента " +inputMasInf[0])
		return
	}

	studentGrade, err := strconv.Atoi(inputMasInf[2])

	if err != nil || studentGrade <= 0 {
		err = errors.New("ошибка ввода курса студента " + inputMasInf[0])
	}

	if err == nil {
		newStudent = &student.Student{Name: inputMasInf[0], Age: studentAge, Grade: studentGrade}
	}
	return 
}

	