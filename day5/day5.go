//yay I made it back to Go
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile5 = flag.String("inputFile", "../inputs/input5.txt", "Path")

type coordinate struct {
	x int
	y int
}

// func newCoordinate(x int, y int) *coordinate {
// 	c := coordinate{x: x, y: y}
// 	return &c
// }
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {

	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile5)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	pairs := [][]string{}

	for _, s := range split {
		pairs = append(pairs, strings.Split(s, " -> "))
	}

	//fmt.Printf("%v", pairs[len(pairs)-1])

	m := make(map[coordinate]int)
	n := make(map[coordinate]int)
	for _, s := range pairs {
		//parse the coordinates
		point1 := strings.Split(s[0], ",")
		point2 := strings.Split(s[1], ",")
		ipoint1x, _ := strconv.Atoi(point1[0])
		ipoint1y, _ := strconv.Atoi(point1[1])
		ipoint2x, _ := strconv.Atoi(point2[0])
		ipoint2y, _ := strconv.Atoi(point2[1])

		//create a coordinate type for each point
		coord1 := coordinate{x: ipoint1x, y: ipoint1y}
		coord2 := coordinate{x: ipoint2x, y: ipoint2y}

		//filter for only horizontal and vertical
		//loop through the coordinates between the two points creating new points and adding them to the map
		if coord1.x == coord2.x {
			for i := min(coord1.y, coord2.y); i <= max(coord1.y, coord2.y); i++ {
				midCoord := coordinate{x: coord1.x, y: i}
				m[midCoord] += 1
			}
		}

		if coord1.y == coord2.y {
			for i := min(coord1.x, coord2.x); i <= max(coord1.x, coord2.x); i++ {
				midCoord := coordinate{x: i, y: coord1.y}
				m[midCoord] += 1
			}
		}
	}
	overlaps := 0
	for _, element := range m {
		if element > 1 {
			overlaps += 1
		}
	}
	//fmt.Printf("%v", overlaps)

	//part two
	for _, s := range pairs {
		//parse the coordinates
		point1 := strings.Split(s[0], ",")
		point2 := strings.Split(s[1], ",")
		ipoint1x, _ := strconv.Atoi(point1[0])
		ipoint1y, _ := strconv.Atoi(point1[1])
		ipoint2x, _ := strconv.Atoi(point2[0])
		ipoint2y, _ := strconv.Atoi(point2[1])

		//create a coordinate type for each point
		coord1 := coordinate{x: ipoint1x, y: ipoint1y}
		coord2 := coordinate{x: ipoint2x, y: ipoint2y}

		//filter for only horizontal and vertical
		//loop through the coordinates between the two points creating new points and adding them to the map
		if coord1.x == coord2.x {
			for i := min(coord1.y, coord2.y); i <= max(coord1.y, coord2.y); i++ {
				midCoord := coordinate{x: coord1.x, y: i}
				n[midCoord] += 1
			}
		}

		if coord1.y == coord2.y {
			for i := min(coord1.x, coord2.x); i <= max(coord1.x, coord2.x); i++ {
				midCoord := coordinate{x: i, y: coord1.y}
				n[midCoord] += 1
			}
		}
		//start bottom left
		if coord1.y < coord2.y && coord1.x < coord2.x {
			j := coord1.y
			for i := coord1.x; i <= coord2.x; i++ {
				midCoord := coordinate{x: i, y: j}
				n[midCoord] += 1
				j++
			}
		}
		//start top left
		if coord1.y > coord2.y && coord1.x < coord2.x {
			j := coord1.y
			for i := coord1.x; i <= coord2.x; i++ {
				midCoord := coordinate{x: i, y: j}
				n[midCoord] += 1
				j--
			}
		}
		//start bottom right
		if coord1.y < coord2.y && coord1.x > coord2.x {
			j := coord1.y
			for i := coord1.x; i >= coord2.x; i-- {
				midCoord := coordinate{x: i, y: j}
				n[midCoord] += 1
				j++
			}
		}
		//start top right
		if coord1.y > coord2.y && coord1.x > coord2.x {
			j := coord1.y
			for i := coord1.x; i >= coord2.x; i-- {
				midCoord := coordinate{x: i, y: j}
				n[midCoord] += 1
				j--
			}
		}
	}
	overlaps2 := 0
	for _, element := range n {
		if element > 1 {
			overlaps2 += 1
		}
	}
	fmt.Printf("%v", overlaps2)
}
