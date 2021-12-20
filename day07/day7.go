//refactored!

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

var inputFile = flag.String("inputFile", "../inputs/input7.txt", "Relative file path to use as input.")

func getMedian(positions []int) int {
	sort.Ints(positions)
	middle := len(positions) / 2
	medianValue := 0
	if len(positions)%2 == 1 {
		medianValue = positions[middle]
	} else {
		medianValue = (positions[middle-1] + positions[middle]) / 2
	}
	return medianValue
}

func getAverage(positions []int) int {
	sum := 0
	for i := 0; i < len(positions); i++ {
		sum += positions[i]
	}
	avg := int(math.Floor(float64(sum) / float64(len(positions))))
	return avg
}

func calculateFuel2(positions []int, point int) int {
	fuel := 0
	for _, v := range positions {
		diff := int(math.Abs(float64(v) - float64(point)))
		fuel += (diff * (diff + 1)) / 2
	}
	return fuel
}

func calculateFuel(positions []int, point int) int {
	fuel := 0
	for _, v := range positions {
		diff := int(math.Abs(float64(v) - float64(point)))
		fuel += diff
	}
	return fuel
}

func main() {
	start := time.Now()
	flag.Parse()
	bytes, _ := ioutil.ReadFile(*inputFile)
	contents := string(bytes)
	positions := []int{}
	for _, v := range strings.Split(contents, ",") {
		intv, _ := strconv.Atoi(v)
		positions = append(positions, intv)
	}
	partOne := calculateFuel(positions, getMedian(positions))
	partTwo := calculateFuel2(positions, getAverage(positions))
	fmt.Printf("Part one: %v \nPart two: %v \n", partOne, partTwo)
	duration := time.Since(start)
	fmt.Println(duration)
}
