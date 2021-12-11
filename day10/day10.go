package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

var inputFile = flag.String("inputFile", "../inputs/input10.txt", "Relative file path to use as input.")

func scoringLines(s string) int {
	scoreMap := make(map[byte]int)
	scoreMap[')'] = 3
	scoreMap[']'] = 57
	scoreMap['}'] = 1197
	scoreMap['>'] = 25137

	score := 0
	if len(s) < 2 {
		return scoreMap[s[0]]
	}
	closers := "}])>"
	stack := []string{}
	for i := 0; i < len(s); i++ {

		current := string(s[i])
		if current == "{" {
			stack = append(stack, "}")
		}
		if current == "(" {
			stack = append(stack, ")")
		}
		if current == "[" {
			stack = append(stack, "]")
		}
		if current == "<" {
			stack = append(stack, ">")
		}
		if strings.Contains(closers, string(s[i])) {
			if len(stack) < 1 {
				score += scoreMap[s[i]]
				return score
			}

			if current != stack[len(stack)-1] {
				score += scoreMap[s[i]]
				return score
			}
			stack = remove(stack, len(stack)-1)
		}
	}
	return score
}
func scoringLinesPartTwo(s string) []string {

	closers := "}])>"
	stack := []string{}
	for i := 0; i < len(s); i++ {

		current := string(s[i])
		if current == "{" {
			stack = append(stack, "}")
		}
		if current == "(" {
			stack = append(stack, ")")
		}
		if current == "[" {
			stack = append(stack, "]")
		}
		if current == "<" {
			stack = append(stack, ">")
		}
		if strings.Contains(closers, string(s[i])) {
			if len(stack) < 1 {
				return []string{}
			}

			if current != stack[len(stack)-1] {
				return []string{}
			}
			stack = remove(stack, len(stack)-1)
		}
	}
	if len(stack) > 0 {
		return stack
	}
	return []string{}
}
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	flag.Parse()
	bytes, _ := ioutil.ReadFile(*inputFile)
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	totalscore := 0
	//part one
	for _, v := range lines {
		totalscore += scoringLines(v)
	}
	fmt.Println(totalscore)

	//part two
	secondscores := []int{}
	scoreMap2 := make(map[string]int)
	scoreMap2[")"] = 1
	scoreMap2["]"] = 2
	scoreMap2["}"] = 3
	scoreMap2[">"] = 4
	for _, v := range lines {
		stack := scoringLinesPartTwo(v)
		if len(stack) > 0 {
			for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
				stack[i], stack[j] = stack[j], stack[i]
			}
			linescore := 0
			for _, x := range stack {
				linescore *= 5
				linescore += scoreMap2[x]
			}
			secondscores = append(secondscores, linescore)
		}
	}
	sort.SliceStable(secondscores, func(i, j int) bool { return secondscores[i] < secondscores[j] })
	middle := math.Floor(float64(len(secondscores) / 2))
	fmt.Println(secondscores[int(middle)])
}
