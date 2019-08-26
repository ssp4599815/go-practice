package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	q, a string
}

func parseLines(lines [][]string) []problem {
	// 初始化一个 列表
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			// 去除字符串两边的空行
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	fmt.Println(lines)

	// 创建解析函数
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanln(&answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}
