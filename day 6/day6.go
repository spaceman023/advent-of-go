//no memory overflows here 😎

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile6 = flag.String("inputFile", "../inputs/input6.txt", "Relative file path to use as input.")

func main() {

	//parsing the file
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile6)
	if err != nil {
		return
	}
	contents := string(bytes)

	//creating a map for each internal timer and the number of fish currently at that internal timer
	fishmap := make(map[int]int)
	for i := 0; i <= 8; i++ {
		fishmap[i] = 0
	}
	for _, v := range strings.Split(contents, ",") {
		intv, _ := strconv.Atoi(v)
		fishmap[intv] += 1
	}

	for i := 1; i <= 256; i++ { //loop through each day

		tempMap := make(map[int]int)
		for k, v := range fishmap { //make a temporary map to store the previous values
			tempMap[k] = v
		}
		for i := 0; i <= 6; i++ {
			fishmap[i] = tempMap[i+1] //fish at internal timers from 0 through 5 just equal the sum of fish at the internal timer 1 greater
		}
		fishmap[6] = tempMap[7] + tempMap[0] //fish at internal timer 7 will move down to 6 but the existing fish at 0 reset to 6 as well
		fishmap[7] = tempMap[8]              //fish at internal timer 8 become internal timer 7
		fishmap[8] = tempMap[0]              //fish at internal timer 0 create new fish and those begin at internal timer 8
	}
	totalFish := 0
	for _, v := range fishmap {
		totalFish += v
	}
	fmt.Println(totalFish)
}
