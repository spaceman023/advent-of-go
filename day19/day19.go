package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

type scanners map[string]scanner

type scanner struct {
	x, y, z   int
	report    []beacon
	positions [][]beacon
}
type beacon struct {
	coords []int
}

func getPossiblePositions(b beacon) []beacon {
	c := b.coords
	o := []int{c[0], c[1], c[2]}
	list := combin.Permutations(3, 3)
	fmt.Println(list)
	positions := []beacon{}
	for _, v := range list {
		positions = append(positions,
			beacon{[]int{o[v[0]], o[v[1]], o[v[2]]}},
		)
	}
	return positions
}
func diffScanners(s1 scanner, s2 scanner) map[int]int {
	diffs := map[int]int{}
	for _, v := range s1.report {
		for _, w := range s2.report {
			for _, x := range w.coords {
				diffs[v.coords[0]-x]++
				diffs[v.coords[1]-x]++
				diffs[v.coords[2]-x]++
			}
		}
	}
	return diffs
}
func main() {

	bytes, _ := ioutil.ReadFile("../inputs/input19.txt")
	chunks := strings.Split(string(bytes), "\n\n")
	scanners := scanners{}
	for _, n := range chunks {
		s := scanner{}
		lines := strings.Split(n, "\n")
		for _, v := range lines[1:] {
			ints := []int{}
			for _, k := range strings.Split(v, ",") {
				i, _ := strconv.Atoi(k)
				ints = append(ints, i)
			}
			s.report = append(s.report, beacon{ints})
		}
		for _, v := range s.report {
			s.positions = append(s.positions, getPossiblePositions(v))
		}

		re := regexp.MustCompile(`\d+`)
		name := re.Find([]byte(lines[0]))
		scanners[string(name)] = s
	}
	//fmt.Println(scanners)
	//fmt.Println(lines)
	d := diffScanners(scanners["0"], scanners["1"])
	//top := map[int]int{}
	max := 0
	for _, v := range d {
		if v > max {
			max = v
		}
	}
	for i, v := range d {
		if v > max-1 {
			fmt.Println(i, v)
		}
	}

	//fmt.Println(top)
}
