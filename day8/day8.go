package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

var inputFile = flag.String("inputFile", "../inputs/input8.txt", "Relative file path to use as input.")

func computeIntersects(unknown string, known string) int {
	count := 0
	for _, v := range known {
		if strings.Contains(unknown, string(v)) {
			count++
		}
	}
	return count
}
func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
func main() {
	start := time.Now()
	flag.Parse()
	bytes, _ := ioutil.ReadFile(*inputFile)
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	inputs := [][]string{}
	outputs := [][]string{}

	for i := 0; i < len(lines); i++ {
		splitLine := strings.Split(lines[i], " | ")
		inputs = append(inputs, strings.Split(splitLine[0], " "))
		outputs = append(outputs, strings.Split(splitLine[1], " "))
	}
	//Part One
	count := 0
	for i := range outputs {
		for _, x := range outputs[i] {
			if len(x) == 2 || len(x) == 4 || len(x) == 3 || len(x) == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
	//Part Two
	decodedOutputs := []string{}
	for i := range inputs {

		lettermap := make(map[int]string)
		lenfives := []string{}
		lensixes := []string{}
		for _, x := range inputs[i] {
			switch len(x) {
			case 2:
				lettermap[1] = x
			case 3:
				lettermap[7] = x
			case 4:
				lettermap[4] = x
			case 5:
				lenfives = append(lenfives, x)
			case 6:
				lensixes = append(lensixes, x)
			case 7:
				lettermap[8] = x
			}
			for _, x := range lenfives {
				if computeIntersects(x, lettermap[1]) == 2 {
					lettermap[3] = x
					continue
				}
				if computeIntersects(x, lettermap[4]) == 2 {
					lettermap[2] = x
					continue
				}
				if computeIntersects(x, lettermap[4]) == 3 && computeIntersects(x, lettermap[1]) == 1 {
					lettermap[5] = x
					continue
				}
			}
			for _, x := range lensixes {
				if computeIntersects(x, lettermap[1]) == 1 {
					lettermap[6] = x
					continue
				}
				if computeIntersects(x, lettermap[4]) == 4 {
					lettermap[9] = x
					continue
				}
				if computeIntersects(x, lettermap[7]) == 3 && computeIntersects(x, lettermap[4]) == 3 {
					lettermap[0] = x
					continue
				}
			}
		}
		decoded := ""
		for _, v := range outputs[i] {
			v = sortStringByCharacter(v)

			for key, value := range lettermap {
				value = sortStringByCharacter(value)
				if v == value {
					decoded += strconv.Itoa(key)
				}
			}
		}
		decodedOutputs = append(decodedOutputs, decoded)

	}

	finalsum := 0
	for _, v := range decodedOutputs {
		asnum, _ := strconv.Atoi(v)
		finalsum += asnum
	}
	fmt.Println(finalsum)
	fmt.Println(time.Since(start))
}
