/*
##### Problem Statement:
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
	"student-management/helper"
)

func main() {

	// helper.AddNewStudent()
	// helper.GetStudentById()
	// helper.GetAllStudents()
	// helper.UpdateNameUsingId()
	// helper.DeleteStudentById()
	// helper.SaveToFile("studentInfo.json")
	// helper.LoadFromFile("studentInfo.json")
	//helper.SortStudentsBySubject()
	//helper.AverageMarks()

	var choice int
	var subjectName string

	for {
		// Display menu
		fmt.Println("Choose an option to perform the operation:")
		fmt.Println("1. Add New Student")
		fmt.Println("2. Get Student by ID")
		fmt.Println("3. Get All Students")
		fmt.Println("4. Update Student Name by ID")
		fmt.Println("5. Delete Student by ID")
		fmt.Println("6. Save to File")
		fmt.Println("7. Load from File")
		fmt.Println("8. Sort Students by Subject")
		fmt.Println("9. Calculate Average Marks")
		fmt.Println("10. Find Maximum Marks In Perticular Subject")
		fmt.Println("11. Find Maximum Total Marks Obtained Student From Batch")
		fmt.Println("0. Exit")

		// Get user input
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Adding New Student")
			helper.AddNewStudent()
		case 2:
			fmt.Println("Enter Student ID: ")
			helper.GetStudentById()
		case 3:
			fmt.Println("GetTing all students")
			helper.GetAllStudents()
		case 4:
			//fmt.Println("Enter Student ID to update: ")
			helper.UpdateNameUsingId()
		case 5:
			//fmt.Println("Enter Student ID to delete: ")
			helper.DeleteStudentById()
		case 6:
			fmt.Println("Saving to file")
			helper.SaveToFile("studentInfo.json")
		case 7:
			fmt.Println("Loading from file")
			helper.LoadFromFile("studentInfo.json")
		case 8:
			fmt.Println("Enter subject name to sort by: ")
			fmt.Scan(&subjectName)
			helper.SortStudentsBySubject(subjectName)
		case 9:
			fmt.Println("Enter Subject Name To Calculating Average marks")
			fmt.Scan(&subjectName)
			helper.AverageMarks(subjectName)
		case 10:
			fmt.Println("Enter Subject Name To Find Out Heighest marks Obtained Student In that Subject")
			fmt.Scan(&subjectName)
			helper.HighestMarksInSubject(subjectName)
		case 11:
			fmt.Println("The Student Having Heighest Total Marks In Batch")
			helper.HighestTotalMarks()
		case 0:
			fmt.Println("Exiting From The Code")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
