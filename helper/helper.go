package helper

import (
	"fmt"

	"github.com/projectdiscovery/gologger"
)

type Student struct {
	Name  string
	Id    int
	Marks map[string]int
}

var students []Student //for all students overall application

func AddNewStudent() {
	var student Student //for each student to be add in students
	fmt.Println("Enter Your Name:")
	fmt.Scan(&student.Name)

	fmt.Println("Enter Your ID:")
	fmt.Scan(&student.Id)

	student.Marks = make(map[string]int)
	Subjects := []string{"English", "Hindi", "Maths", "Science", "CS"}

	for _, subject := range Subjects {
		var marks int
		fmt.Println("Enter Your Marks Of ", subject)
		fmt.Scan(&marks)
		student.Marks[subject] = marks

	}
	students = append(students, student)
	gologger.Info().Msgf("The Student With Name %s and With ID %d and Marks %v Is Added Successfully", student.Name, student.Id, student.Marks)

}

func GetStudentById() {
	var idToFind int

	fmt.Println("Enter Student Id That You want To Find Out")
	fmt.Scan(&idToFind)

	for _, student := range students {
		if student.Id == idToFind {
			fmt.Println(student)
			return
		} else {
			gologger.Error().Msgf("The Mention Student is Not Found")
		}

	}

}

func GetAllStudents() {
	if len(students) == 0 {
		gologger.Error().Msgf("There Is No Student Present")
	}

	for _, student := range students {
		fmt.Println("The All Students Are\n :", student)
	}
}

func UpdateNameUsingId() {
	var idOfstudent int
	var newName string

	fmt.Println("Enter The Id Of Student Whose Name You Want To Change")
	fmt.Scan(&idOfstudent)

	for i, student := range students {
		if student.Id == idOfstudent {
			fmt.Println("Enter NewName To The Student")
			fmt.Scan(&newName)

			students[i].Name = newName
			gologger.Info().Msgf("Student Name Is Updated Succesfully")
			return
		} else {
			gologger.Error().Msgf("The Mention Student is Not Found")
		}
	}

}

func DeleteStudentById() {
	var idToDelete int

	fmt.Println("Enter The Student ID That You Want To Delete")
	fmt.Scan(&idToDelete)

	for i, student := range students {
		if student.Id == idToDelete {
			students = append(students[:i], students[i+1:]...)
			gologger.Info().Msgf("The Selected Student Deleted Successfully")
		} else {
			gologger.Error().Msgf("The Mention Student is Not Found")
		}
	}

}
