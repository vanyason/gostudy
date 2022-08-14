/**
* A simple program, where you can find examples of the flag ,
* csv, goroutines usages
*/

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type quizData struct {
	question string
	answer   string
}

func parseCmdFlags() (csvFilename string, timeLimit int, randomize bool) {
	flag.StringVar(&csvFilename, "csv", "problems.csv", "a csv file in the format of 'question.answer'")
	flag.IntVar(&timeLimit, "limit", 30, "the time limit for the quiz in seconds")
	flag.BoolVar(&randomize, "rand", false, "randomize csv content")
	flag.Parse()
	return csvFilename, timeLimit, randomize
}

func readCsv(filename string) (lines [][]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return lines, err
	}

	reader := csv.NewReader(file)

	lines, err = reader.ReadAll()
	if err != nil {
		return lines, fmt.Errorf("failed to parse the provided csv file (%s) : %s", filename, err.Error())
	}

	if len(lines) == 0 {
		return lines, fmt.Errorf("csv file (%s) has no lines", filename)
	}

	return lines, err
}

func parseCsv(lines [][]string, filename string) (qestionsAnswers []quizData, err error) {
	qestionsAnswers = make([]quizData, len(lines))

	for i, line := range lines {
		if len(line) != 2 {
			return qestionsAnswers, fmt.Errorf("problem CSV parsing (%s): line %d is not question-answer format", filename, i+1)
		}

		line[1] = strings.TrimSpace(line[1])

		if _, err := strconv.Atoi(line[1]); err != nil {
			return qestionsAnswers, fmt.Errorf("problem CSV parsing (%s): line %d - answer is not numeric", filename, i+1)
		}

		qestionsAnswers[i] = quizData{
			question: line[0],
			answer:   line[1],
		}
	}

	return qestionsAnswers, err
}

func mixQuestions(questions *[]quizData) {
	if questions == nil {
		return
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*questions), func(i, j int) {
		(*questions)[i], (*questions)[j] = (*questions)[j], (*questions)[i]
	})

}

func playQuiz(timer *time.Timer, questions []quizData) (result string, err error) {
	if timer == nil {
		return result, fmt.Errorf("playQuiz func : timer is nil")
	}
	if len(questions) == 0 {
		return result, fmt.Errorf("playQuiz func : no questions")
	}

	correctAnswers := 0

	for i, p := range questions {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)

		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			return fmt.Sprintf("Time is over: You scored %d out of %d\n", correctAnswers, len(questions)), err
		case answer := <-answerCh:
			if answer == p.answer {
				correctAnswers++
			}
		}

	}

	return fmt.Sprintf("You scored %d out of %d\n", correctAnswers, len(questions)), err
}

func main() {
	csvFilename, quizTimeLimitSecs, randomize := parseCmdFlags()

	lines, err := readCsv(csvFilename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	questions, err := parseCsv(lines, csvFilename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if randomize {
		mixQuestions(&questions)
	}

	timer := time.NewTimer(time.Duration(quizTimeLimitSecs) * time.Second)

	result, err := playQuiz(timer, questions)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}
