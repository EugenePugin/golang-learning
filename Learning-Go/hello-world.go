package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Define the structure that matches the input JSON data.
type GroupData struct {
	ID       int       `json:"ID"`
	Number   string    `json:"Number"`
	Year     int       `json:"Year"`
	Students []Student `json:"Students"`
}

type Student struct {
	LastName   string `json:"LastName"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	Birthday   string `json:"Birthday"`
	Address    string `json:"Address"`
	Phone      string `json:"Phone"`
	Rating     []int  `json:"Rating"` // This is the slice of grades we need to count
}

// Define the structure for the output JSON data.
type AverageResult struct {
	Average float64 `json:"Average"`
}

func main() {
	// 1. Read the input JSON data from standard input (os.Stdin)
	jsonData, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// 2. Unmarshal the JSON data into the Go struct
	var group GroupData
	err = json.Unmarshal(jsonData, &group)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// 3. Calculate the total number of grades and total number of students
	totalGrades := 0
	numStudents := len(group.Students)

	if numStudents == 0 {
		// Handle the case where there are no students to avoid division by zero
		outputResult(0.0)
		return
	}

	for _, student := range group.Students {
		totalGrades += len(student.Rating) // Add the count of ratings for each student
	}

	// 4. Calculate the average number of grades
	averageGrades := float64(totalGrades) / float64(numStudents)

	// 5. Format and write the output JSON
	outputResult(averageGrades)
}

// outputResult marshals the average value into the required JSON format and prints it to standard output.
func outputResult(average float64) {
	result := AverageResult{
		Average: average,
	}

	outputJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating JSON output: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(outputJSON))
}
