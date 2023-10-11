package main

import (
	"encoding/csv" // Package that allows us to read/write CSV.
	"flag"         // This package provides support for defining and parsing command-line flags.
	"fmt"          // This package provides formatted I/O (iostream for C++, or stdlib.h for C)
	"os"           // This package provides OS functionality such as env variables and file operations.
	"strings"      // This package provides string manipulation functions.
	"time"         // This package provides the use of time-based functions such as "timer" and "tickers"
)

// Usage of flags in GoLang?
// Programs can define flags using the flag.XXX functions, where XXX represents the data type of the flag.
// Users can then provide values for these flags when running the program from the command line.
// The flag package takes care of parsing and assigning the flag values to variables within the program.

/*
	Ex: Program: {
	 	name := flag.String("name", "Guest", "Specify your name")
    	age := flag.Int("age", 0, "Specify your age")
    	flag.Parse()
		fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
	}

		Executing ./program -name Alice -age 25 would output: "Hello, Alice! You are 25 years old."
*/

func main() {
	/*
		Overall, these lines of code allow the program to accept command-line flags
		for specifying a CSV file and a time limit for a quiz. The csvFilename and
		timeLimit variables will hold the values provided through the flags, which
		can be used in the program's logic.
	*/

	// func flag.String(name string, value string, usage string) *string // Usage of flag.String().
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	// func flag.Int(name string, value int, usage string) *int
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	// flag.Parse() examines the command-line arguments and assigns values
	// to the corresponding variables defined by the flags.
	flag.Parse()

	// When the user supplies the csvfilename, os (the package) . Open attempts to
	// open the file => func os.Open(name string) (*os.File, error). The output being
	// a pointer os.File and error. If error is nil, then the file had been opened successfully.
	// We instantiate this 'file' instance into the var file.
	file, err := os.Open(*csvFilename)
	// if error is not nil, then we call the exit function (last function)
	// The exit function does two things -> Outputs the error in the terminal,
	// and exits the program with a "1" exit code.

	/*

		What does sPrintf do?
		Ex:
		name := "Oscar"
		age := 25
		formattedString := fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)

		What is actually stored in formattedString, would be "Hello, my name is Oscar and I am 25 years old."
		What is really passed into exit then, is "Failed to open the CSV file *csvFilename", (or whatever the filename was
		given by the user) and is printed to the terminal.

	*/
	// Check if there was an error opening the CSV file. If so, exit the program with an error message.
if err != nil {
	exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
}

// Create a new CSV reader from the opened file.
r := csv.NewReader(file)

// Read all the lines from the CSV file.
lines, err := r.ReadAll()
if err != nil {
	exit("Failed to parse the provided CSV file.")
}

// Parse the read lines into a list of problems.
problems := parseLines(lines)

// Create a new timer for the quiz with the given time limit.
timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

// Variable to track the number of correct answers.
correct := 0

// Label for the outer loop to break out of it from inner select statement.
problemloop:
for i, p := range problems {
	// Prompt the user with the problem.
	fmt.Printf("Problem #%d: %s = ", i+1, p.q)

	// Channel to receive the answer from the user.
	answerCh := make(chan string)

	// Goroutine to read user's answer.
	go func() {
		var answer string
		fmt.Scanf("%s\n", &answer)
		answerCh <- answer
	}()

	// Select statement to wait either for the timer to finish or to get an answer from the user.
	select {
	case <-timer.C:
		// Timer has finished.
		fmt.Println()
		break problemloop
	case answer := <-answerCh:
		// Received an answer from the user.
		if answer == p.a {
			correct++
		}
	}
}

// Print the user's score at the end of the quiz.
fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

// Function to parse lines from the CSV into a list of problems.
func parseLines(lines [][]string) []problem {
ret := make([]problem, len(lines))
for i, line := range lines {
	ret[i] = problem{
		q: line[0],
		a: strings.TrimSpace(line[1]),
	}
}
return ret
}

// Struct to define a problem with a question and an answer.
type problem struct {
q string
a string
}

// Function to print an error message and then exit the program.
func exit(msg string) {
fmt.Println(msg)
os.Exit(1)
}