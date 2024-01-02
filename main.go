package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var scanner *bufio.Scanner

func main () {
	filename := flag.String("file", "problems.csv", "Problem File")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		println("There was an error opening the csv file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	questions, err := reader.ReadAll()
	if err != nil {
		fmt.Println("There was an error reading the csv file:", err)
	}
	
	scanner = bufio.NewScanner(os.Stdin)
	score := 0

	for i, q := range(questions) {
		fmt.Printf("%d: %s = ", i + 1, q[0])
		scanner.Scan()
		if q[1] == strings.TrimSpace(scanner.Text()){
			score++
		}
	}
	fmt.Printf("Final score: %d/%d\n", score, len(questions))
}