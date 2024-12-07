/*##### Problem Statement:
- Create a program that manages a list of students.
- Each student should be represented by a struct containing:
    - Name (string)
    - ID (int)
    - Marks (map[string]int)
- Assume every student has 5 subjects English, Hindi and Maths, Science and CS.
- The application should be a simple CLI program that allows different options using which user can interact (refer calculator program from previous assignment)
- The program should do error handling and show appropriate errors if something goes wrong.
- The program should output logs. Use https://github.com/projectdiscovery/gologger pkg for logging, set appropriate log levels.

##### Tasks:
###### Level 1:
- Write a function to add a new student to the list.
- Write a function to output student with matcing id from the list.
- Write a function to output all students from the list.
- Write a function to update name using id.
- Write a function to delete a student by id.

###### Level 2:
- Write a function to save students info into a JSON file.
- Write a function to load students info from a JSON file.

###### Level 3:
- Write a function to sort student based on marks in particular subject.
- Write a functtion to find the avg marks in a paritcular subject.
- Write a function to print the student name with highest marks in a particular subject.
- Write a function to print the student name with highest total marks in the batch.
*/

package main

import (
	"fmt"
)

type Student struct {
	name  string
	id    int
	marks map[string]int
}

var students []Student

func main() {
	fmt.Println("Hello")
	for {
		printMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanf("%d\n", &choice)
		if err != nil {
			continue
		}
		if choice == 5 {
			break
		}
		handleChoice(choice)
	}
}

func printMenu() {
	fmt.Println("1. Add Student")
	fmt.Println("2. Get Student by ID")
	fmt.Println("3. Get All Students")
	fmt.Println("4. Update Student Name by ID")
	fmt.Println("5. For Exit")
}

func handleChoice(choice int) {
	switch choice {
	case 1:
		addNewStudent()
	case 2:
		getStudentByID()
	case 3:
		getAllStudents()
	case 4:
		updateNameUsingId()
	case 5:
		break
	default:
		fmt.Println("Please Enter Valid Choice")
	}
}

func addNewStudent() {
	var student Student
	fmt.Print("Enter student ID: ")
	fmt.Scan(&student.id)
	fmt.Print("Enter student Name: ")
	fmt.Scan(&student.name)

	student.marks = make(map[string]int)
	subjects := []string{"English", "Hindi", "Maths", "Science", "CS"}
	for _, subject := range subjects {
		fmt.Printf("Enter Marks for %s: ", subject)
		var marks int
		fmt.Scan(&marks)
		student.marks[subject] = marks

	}

	students = append(students, student)
	fmt.Printf("Student Id=%d, Name=%s, Marks=%v\n", student.id, student.name, student.marks)
}

func getStudentByID() {
	var id int
	fmt.Print("Enter student Id: ")
	fmt.Scan(&id)

	for _, student := range students {
		if student.id == id {
			fmt.Printf("Student: %+v\n", student)
			return
		}
	}

}

func getAllStudents() {
	if len(students) == 0 {
		fmt.Println("No Student Found")
		return
	}
	for _, student := range students {
		fmt.Printf("%+v\n", student)
	}
}

func updateNameUsingId() {
	var id int
	var newName string
	fmt.Print("Enter student Id: ")
	fmt.Scan(&id)
	fmt.Print("Enter New Name For Student:")
	fmt.Scan(&newName)

	for i, student := range students {
		if student.id == id {
			students[i].name = newName
		}
		fmt.Printf("%+v\n", student)
	}

}
