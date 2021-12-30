package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coord struct {
	r, c int
}

func (c coord) getNeighbors(l litmap) []coord {
	coords := []coord{}
	for rd := -1; rd <= 1; rd++ {
		for cd := -1; cd <= 1; cd++ {
			coords = append(coords, coord{c.r + rd, c.c + cd})
		}
	}
	return coords
}

type alg map[int]bool

type litmap map[coord]bool

func (l litmap) iterate(a alg, i int) litmap {
	newMap := make(litmap)
	var d bool

	if a[0] && !a[511] {
		d = (i%2 == 0)
	} else {
		d = false
	}

	min := -2 * i
	max := 100 + 2*i
	for r := min; r < max; r++ {
		for c := min; c < max; c++ {
			bin := ""
			z := coord{r, c}
			for _, v := range z.getNeighbors(l) {
				if val, ok := l[v]; ok && val {
					bin += "1"
				} else if !ok && d {
					bin += "1"
				} else {
					bin += "0"
				}
			}
			newMap[coord{r, c}] = a[strToBin(bin)]
		}
	}
	return newMap
}

func strToBin(s string) int {
	i, _ := strconv.ParseInt(s, 2, 64)
	return int(i)
}

func parseInput(s string) (alg, litmap) {
	bytes, _ := ioutil.ReadFile(s)
	chunks := strings.Split(string(bytes), "\n\n")
	lines := strings.Split(chunks[1], "\n")
	a := make(alg)
	b := make(litmap)
	for i, v := range chunks[0] {
		if v == '#' {
			a[i] = true
		}
	}
	for r, l := range lines {
		for c, v := range l {
			if v == '.' {
				b[coord{r, c}] = false
			} else {
				b[coord{r, c}] = true
			}
		}
	}
	return a, b
}
func main() {
	algorithm, lightmap := parseInput("../inputs/input20.txt")
	for i := 1; i <= 50; i++ {
		lightmap = lightmap.iterate(algorithm, i)
	}
	count := 0
	for _, v := range lightmap {
		if v {
			count++
		}
	}
	fmt.Println(count)
}
