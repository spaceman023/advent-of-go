package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "../inputs/input13.txt", "Relative file path to use as input.")

type point struct {
	x, y int
}
type points []point

func executeFold(p points, f string) points {
	fLine, _ := strconv.Atoi(string(f[1:]))
	newP := points{}
	if f[0] == 'y' {
		for _, v := range p {
			if v.y > fLine {
				newPoint := point{v.x, 2*fLine - v.y}
				if !contains(newP, newPoint) {
					newP = append(newP, newPoint)
				}
			} else {
				if !contains(newP, v) {
					newP = append(newP, point{v.x, v.y})
				}
			}
		}
	}
	if f[0] == 'x' {
		for _, v := range p {
			if v.x > fLine {
				newPoint := point{2*fLine - v.x, v.y}
				if !contains(newP, newPoint) {
					newP = append(newP, newPoint)
				}
			} else {
				if !contains(newP, v) {
					newP = append(newP, point{v.x, v.y})
				}
			}
		}
	}
	return newP
}
func contains(ps points, p point) bool {
	for _, v := range ps {
		if v == p {
			return true
		}
	}
	return false
}
func printPaper(p points) {
	xd := 0
	yd := 0
	for _, v := range p {
		if v.x > xd {
			xd = v.x
		}
		if v.y > yd {
			yd = v.y
		}
	}
	paper := make([][]string, yd+1)
	for i := range paper {
		paper[i] = make([]string, xd+1)
	}
	for y := 0; y <= yd; y++ {
		for x := 0; x <= xd; x++ {
			if contains(p, point{x, y}) {
				paper[y][x] = "#"
			} else {
				paper[y][x] = "."
			}
		}
	}
	for _, v := range paper {
		fmt.Println(v)
	}
}
func main() {
	flag.Parse()
	bytes, _ := ioutil.ReadFile(*inputFile)
	contents := string(bytes)
	lines := strings.Split(contents, "\n\n")
	parsefolds := strings.Split(lines[1], "\n")
	lines = strings.Split(lines[0], "\n")
	folds := []string{}
	for _, v := range parsefolds {
		x := strings.Split(v, "=")
		fold := string(x[0][len(x[0])-1]) + x[1]
		folds = append(folds, fold)
	}

	paper := points{}
	for _, v := range lines {
		pair := strings.Split(v, ",")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		paper = append(paper, point{x, y})
	}

	//part one
	paper1 := executeFold(paper, folds[0])
	fmt.Println(len(paper1))
	//part two
	finalPaper := paper
	for _, v := range folds {
		finalPaper = executeFold(finalPaper, v)
	}
	printPaper(finalPaper)
}
