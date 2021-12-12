package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "../inputs/input11.txt", "Relative file path to use as input.")

type point struct {
	x int
	y int
}

type points []point

type pointMap map[point]int

func getNeighbors(p point, m pointMap) points {
	neighbors := points{
		point{p.x, p.y + 1},
		point{p.x, p.y - 1},
		point{p.x - 1, p.y},
		point{p.x + 1, p.y},
		point{p.x + 1, p.y + 1},
		point{p.x + 1, p.y - 1},
		point{p.x - 1, p.y + 1},
		point{p.x - 1, p.y - 1}}
	filteredNeighbors := points{}
	for _, v := range neighbors {
		if _, ok := m[v]; ok { //not outside the map
			filteredNeighbors = append(filteredNeighbors, v)
		}
	}

	return filteredNeighbors
}

func update(m pointMap, s int) (pointMap, int, bool) {
	allFlashed := false
	flashed := map[point]bool{}
	currMap := m
	currScore := s
	for key := range currMap {
		currMap[key]++
	}
	stillFlashing := true
	for stillFlashing {
		stillFlashing = false
		for key, value := range currMap {
			if value > 9 && !flashed[key] {
				currScore++
				flashed[key] = true
				stillFlashing = true
				neighbors := getNeighbors(key, currMap)
				for _, v := range neighbors {
					currMap[v]++
				}
			}
		}
	}
	for key, value := range flashed {
		if value {
			currMap[key] = 0
		}
	}
	if len(flashed) == len(currMap) {
		allFlashed = true
	}
	return currMap, currScore, allFlashed
}

func main() {

	flag.Parse()
	bytes, _ := ioutil.ReadFile(*inputFile)
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	//get all values into a two-dimensional array
	elevationMap := [][]int{}
	for _, v := range lines {
		arr := []int{}
		for _, x := range v {
			arr = append(arr, (int(x))-'0')
		}
		elevationMap = append(elevationMap, arr)
	}

	//transform 2-d array into a map with points as keys and power as values
	pm := pointMap{}
	for i, v := range elevationMap {
		for j, x := range v {
			thispoint := point{j, i}
			pm[thispoint] = x
		}
	}
	jm := pointMap{} //make a copy for part two
	for k, v := range pm {
		jm[k] = v
	}
	//part one
	flashes := 0
	for i := 0; i < 100; i++ {
		pm, flashes, _ = update(pm, flashes)
	}
	fmt.Println(flashes)
	//part two
	flashes = 0
	allFlashed := false
	steps := 0
	for !allFlashed {
		steps++
		jm, flashes, allFlashed = update(jm, flashes)
	}
	fmt.Println(steps)
}
