package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Question struct {
	question string
	answer   string
}

func readCsvProblem() []Question {
	f, err := os.Open("./problems.csv")
	if err != nil {
		log.Fatal("Unable to read input file: "+"problems.csv", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+" problems.csv", err)
	}
	var questions []Question
	for _, record := range records {
		newQuestion := Question{
			question: record[0],
			answer:   record[1],
		}
		questions = append(questions, newQuestion)
	}

	return questions
}

func startQuiz() {
	questions := readCsvProblem()
	var answer string
	totalScore := 0
	for index, question := range questions {
		showQuestion(index+1, question.question)
		answer = getAnswer()
		totalScore += getTheScore(answer, question.answer)
	}
	fmt.Printf("Total Score: %d", totalScore)

}
func showQuestion(numOfQuestion int, question string) {
	fmt.Printf("Problem #%d:  %s = ", numOfQuestion, question)
}

func getAnswer() string {
	var ans string
	fmt.Scanln(&ans)
	return ans
}

func getTheScore(userAnswer string, actualAnswer string) int {
	if strings.Compare(strings.Trim(strings.ToLower(userAnswer), "\n"), actualAnswer) == 0 {
		fmt.Println("Correct answer!")
		return 1
	}
	fmt.Println("Wrong answer!")
	return 0
}

func main() {
	// records := readCsvProblem()
	// fmt.Println(records)
	startQuiz()
}
