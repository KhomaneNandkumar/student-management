package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/projectdiscovery/gologger"
)

type Student struct {
	Name  string         `json:"name"`
	Id    int            `json:"id"`
	Marks map[string]int `json:"marks"`
}

// var students []Student //for all students overall application
var (
	students []Student
	nextID   int = 1
	mu       sync.Mutex
)

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
	fmt.Println("The All Students Are\n :")
	for _, student := range students {
		fmt.Println(student)
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

func SaveToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		gologger.Error().Msgf("Failed to create file: %v", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(students); err != nil {
		gologger.Error().Msgf("Failed to save data: %v", err)
		return
	}
	gologger.Info().Msg("Student Information Saved Successfully")
}

func LoadFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		gologger.Error().Msgf("Failed to Open File: %v", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&students); err != nil {
		gologger.Error().Msgf("Failed to Load Data: %v", err)
		return
	}
	gologger.Info().Msg("Data Loaded Successfully")
	if len(students) > 0 {
		mu.Lock()
		nextID = students[len(students)-1].Id + 1
		mu.Unlock()
	}
}

func SortStudentsBySubject(subject string) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Marks[subject] > students[j].Marks[subject]
	})
	gologger.Info().Msg("Students sorted by subject marks")
	GetAllStudents()
}

func AverageMarks(subject string) {
	total := 0
	count := 0
	for _, student := range students {
		if mark, exists := student.Marks[subject]; exists {
			total += mark
			count++
		}
	}
	if count == 0 {
		gologger.Warning().Msg("No marks available for the subject")
		return
	}
	fmt.Printf("Average marks in %s: %.2f\n", subject, float64(total)/float64(count))
}

func HighestMarksInSubject(subject string) {
	var topStudent *Student
	maxMarks := -1
	for _, student := range students {
		if mark, exists := student.Marks[subject]; exists && mark > maxMarks {
			topStudent = &student
			maxMarks = mark
		}
	}
	if topStudent == nil {
		gologger.Warning().Msg("No marks available for the subject")
		return
	}
	fmt.Printf("Top student in %s: %s with %d marks\n", subject, topStudent.Name, maxMarks)
}

func HighestTotalMarks() {
	var topStudent *Student
	maxTotal := -1
	for _, student := range students {
		total := 0
		for _, marks := range student.Marks {
			total += marks
		}
		if total > maxTotal {
			topStudent = &student
			maxTotal = total
		}
	}
	if topStudent == nil {
		gologger.Warning().Msg("No students available")
		return
	}
	fmt.Printf("Top student in batch: %s with %d total marks\n", topStudent.Name, maxTotal)
}
