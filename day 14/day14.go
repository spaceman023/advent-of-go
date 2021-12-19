package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type counter map[string]int

func step(chain counter, polymers counter, pairs map[string]string) (counter, counter) {
	newChain := make(counter)
	newPolymers := make(counter)
	for k := range chain {
		newChain[k] = 0
	}
	for k, v := range polymers {
		newPolymers[k] = v
	}
	for i, v := range chain {
		newChain[string(i[0])+pairs[i]] += v
		newChain[pairs[i]+string(i[1])] += v
		newPolymers[string(pairs[i])] += v
	}
	return newChain, newPolymers
}

func main() {
	bytes, _ := ioutil.ReadFile("../inputs/input14.txt")
	parsed := strings.Split(string(bytes), "\n\n")
	start := parsed[0]
	lines := strings.Split(parsed[1], "\n")
	chain := make(counter)
	polymers := make(counter)
	pairs := make(map[string]string)

	for _, v := range lines {
		entrySplit := strings.Split(v, " -> ")
		pairs[entrySplit[0]] = entrySplit[1]
		chain[entrySplit[0]] = 0
	}

	for i := 0; i < len(start)-1; i++ {
		pair := string(start[i]) + string(start[i+1])
		chain[string(pair)]++
	}

	for i := 0; i < len(start); i++ {
		polymers[string(start[i])]++
	}

	for i := 0; i < 40; i++ {
		chain, polymers = step(chain, polymers, pairs)
	}

	max := 0
	min := int(^uint(0) >> 1)
	for _, v := range polymers {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)
}
