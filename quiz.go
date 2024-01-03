package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question/answer")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "Randomize the question order")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("There was an error opening the csv file: %s\n", *filename))
	}
	defer file.Close()

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse csv file: %s\n", *filename))
	}

	problems := parseLines(lines)

	if *shuffle {
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		// Create  an answer channel
		answerCh := make(chan string)
		go func() {
			// Set up closure
			var answer string
			fmt.Scanf("%s\n", &answer)
			// Sent the user's answer to the answer channel
			answerCh <- answer
		}()
		// if we get a message from the timer  then we can stop
		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		// when answerCh receives an answer send it to this var and check it
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("Final Score: %d/%d\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0], a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
