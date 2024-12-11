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
	"student-management/helper"
)

func main() {

	helper.AddNewStudent()
	helper.GetStudentById()
	helper.GetAllStudents()
	helper.UpdateNameUsingId()
	helper.DeleteStudentById()
}
