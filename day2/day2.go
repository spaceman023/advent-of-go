package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile2 = flag.String("inputFile", "../inputs/input2.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile2)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	//split = split[:len(split)-1]

	coords := make(map[string]int)
	coords["x"] = 0
	coords["y"] = 0

	//part one
	for _, i := range split {
		parsed := strings.Split(i, " ")
		direction := parsed[0]
		distance, _ := strconv.Atoi(parsed[1])
		switch direction {
		case "forward":
			coords["x"] += distance

		case "down":
			coords["y"] += distance
		case "up":
			coords["y"] -= distance
		}
	}
	fmt.Printf("%v \n", coords["x"]*coords["y"])

	//part two
	coords["x"] = 0
	coords["y"] = 0
	aim := 0
	for _, i := range split {
		parsed := strings.Split(i, " ")
		direction := parsed[0]
		distance, _ := strconv.Atoi(parsed[1])
		switch direction {
		case "forward":
			coords["x"] += distance
			coords["y"] += (distance * aim)
		case "down":
			aim += distance
		case "up":
			aim -= distance
		}
	}
	fmt.Printf("%v \n", coords["x"]*coords["y"])
}
