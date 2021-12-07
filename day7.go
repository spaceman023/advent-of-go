//tried to go fast but am too stupid - will refactor this tomorrow

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/input7.txt", "Relative file path to use as input.")
var factVal = 1

func main() {

	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	positions := []int{}
	for _, v := range strings.Split(contents, ",") {
		intv, _ := strconv.Atoi(v)
		positions = append(positions, intv)
	}
	sort.Ints(positions)

	middle := len(positions) / 2
	medianValue := 0
	if len(positions)%2 == 1 {
		medianValue = positions[middle]
	} else {
		medianValue = (positions[middle-1] + positions[middle]) / 2
	}
	fuel := 0
	sum := 0
	for i := 0; i < len(positions); i++ {
		sum += positions[i]
	}
	avg := math.Floor(float64(sum) / float64(len(positions)))
	for _, v := range positions {
		diff := int(math.Abs(float64(v) - float64(avg)))
		total := 0
		for i := 0; i <= diff; i++ {
			total += i
		}
		fuel += total
	}
	fmt.Println(avg)
	fmt.Println(medianValue)
	fmt.Println(fuel)
}
