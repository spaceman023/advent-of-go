package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var inputFile = flag.String("inputFile", "../inputs/input9.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, _ := ioutil.ReadFile(*inputFile)
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]

	//get all values into a two-dimensional array
	elevationMap := [][]int{}
	for _, v := range lines {
		arr := []int{}
		for _, x := range v {
			arr = append(arr, (int(x))-'0')
		}
		elevationMap = append(elevationMap, arr)
	}

	//define x,y coord types
	type Point struct {
		x int
		y int
	}

	type Points []Point

	//transform 2-d array into a map with points as keys and heights as values
	pointMap := make(map[Point]int)
	for i, v := range elevationMap {
		for j, x := range v {
			thispoint := Point{j, i}
			pointMap[thispoint] = x
		}
	}

	//part one
	var lowestPoints = Points{}
	lowestSum := 0
	//getting neighbors and assigning to lowest
	for key, value := range pointMap {
		var neighbors = Points{Point{key.x, key.y + 1}, Point{key.x, key.y - 1}, Point{key.x - 1, key.y}, Point{key.x + 1, key.y}}
		islowest := true
		for _, x := range neighbors {
			if val, ok := pointMap[x]; ok {
				if value >= val {
					islowest = false
				}
			}

		}
		if islowest {
			lowestPoints = append(lowestPoints, key)
		}
	}
	for _, v := range lowestPoints {
		lowestSum += pointMap[v] + 1
	}
	fmt.Println(lowestSum)
	//part two
	var basins = []map[Point]int{}
	for _, v := range lowestPoints {
		basin := make(map[Point]int)
		queue := Points{v}
		for len(queue) > 0 {
			curr := queue[0]
			var neighbors = Points{Point{curr.x, curr.y + 1}, Point{curr.x, curr.y - 1}, Point{curr.x - 1, curr.y}, Point{curr.x + 1, curr.y}}
			for _, x := range neighbors {
				if val, ok := pointMap[x]; ok { //not outside the map
					if pointMap[x] == 9 {
						continue
					}
					if _, ok := basin[x]; ok { //not already in basin
					} else {
						queue = append(queue, x)
						basin[x] = val //add to basin
					}
				}
			}
			queue = queue[1:]
		}
		basins = append(basins, basin)
	}
	sort.SliceStable(basins, func(i, j int) bool { return len(basins[i]) > len(basins[j]) })
	basintotal := 1
	for _, v := range basins[0:3] {
		basintotal *= len(v)
	}
	fmt.Println(basintotal)
}
