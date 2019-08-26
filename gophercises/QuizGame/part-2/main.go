package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	q, a string
}

func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[0]),
		}
	}
	return ret
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var (
		timeLimit = 30
		correct   = 0
	)
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)

		go func() {
			xxx
		}()

		select {
		xxx
		}
	}

}
