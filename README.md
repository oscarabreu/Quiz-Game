# Quiz Game 

This Go program is a quiz game that allows users to test their knowledge by answering a series of questions. It accepts a CSV file containing questions and answers, a time limit for each question, and presents the user with questions one by one. Users must answer within the specified time limit.

## Prerequisites

You need Go installed on your system to run this program.

## Usage

1. Clone or download the program's source code.
2. Open a terminal and navigate to the directory containing the program.
3. Compile the program using the following command:```go build quiz_game.go```
4. Run the compiled program with optional flags:
   1. -csv (string): Specifies the CSV file containing questions and answers. Default is "problems.csv".
   2. -limit (int): Sets the time limit for each question in seconds. Default is 30 seconds.
5. The program will display each question one by one. Enter your answer within the specified time limit for each question.
6. 6After completing the quiz, the program will display your score.
7. To exit the program at any time, press Ctrl+C.

## CSV File Format

The CSV file should be in the following format:

```
question,answer
What is the capital of France?,Paris
What is 2 + 2?,4
```

Each line consists of a question and its corresponding answer separated by a comma.
