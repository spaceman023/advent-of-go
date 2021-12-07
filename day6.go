//no memory overflows here 😎

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/input6.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	fishmap := make(map[int]int)
	for i := 0; i <= 8; i++ {
		fishmap[i] = 0
	}
	for _, v := range strings.Split(contents, ",") {
		intv, _ := strconv.Atoi(v)
		fishmap[intv] += 1
	}
	for i := 1; i <= 256; i++ {
		tempMap := make(map[int]int)
		for k, v := range fishmap {
			tempMap[k] = v
		}
		for i := 0; i <= 6; i++ {
			fishmap[i] = tempMap[i+1]
		}
		fishmap[6] = tempMap[7] + tempMap[0]
		fishmap[7] = tempMap[8]
		fishmap[8] = tempMap[0]
	}
	totalFish := 0
	for _, v := range fishmap {
		totalFish += v
	}
	fmt.Println(totalFish)
}
