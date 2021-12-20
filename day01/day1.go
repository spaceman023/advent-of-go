package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile1 = flag.String("inputFile", "../inputs/input1.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile1)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]
	//fmt.Printf("%v", split)
	count := 0
	last := -1
	for _, s := range split {
		i, _ := strconv.Atoi(s)
		if last != -1 && i > last {
			count++
		}
		last = i
	}

	bcount := 0
	window := make([]int, 0)
	for i := 0; i < 3; i++ {
		val, _ := strconv.Atoi(split[i])
		window = append(window, val)
	}
	for i := 3; i < len(split); i++ {
		val, _ := strconv.Atoi(split[i])
		tot1 := sum(window)
		//fmt.Printf("first %v", window)
		window = append(window[1:3], val)
		//fmt.Printf("second %v", window)
		tot2 := sum(window)
		if tot2 > tot1 {
			bcount++
		}
	}
	fmt.Printf("%v", bcount)

}
func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
