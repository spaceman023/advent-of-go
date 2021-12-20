//took me far too long, but it's finally done!

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func islower(str string) bool {
	if str[0] >= 'a' && str[0] <= 'z' {
		return true
	}
	return false
}

func hasDoubleSmall(currPath []string) bool {
	seen := make(map[string]int)
	for _, v := range currPath {
		seen[v]++
		if seen[v] > 1 && islower(v) {
			return true
		}
	}
	return false
}

func copyAndAppend(path []string, cave string) []string {
	y := make([]string, len(path))
	copy(y, path)
	y = append(y, cave)
	return y
}

func main() {

	bytes, _ := ioutil.ReadFile("../inputs/input12.txt")
	lines := strings.Split(string(bytes), "\n")
	caves := make(map[string][]string)
	for _, v := range lines {
		entrySplit := strings.Split(v, "-")
		caves[entrySplit[0]] = append(caves[entrySplit[0]], entrySplit[1])
		caves[entrySplit[1]] = append(caves[entrySplit[1]], entrySplit[0])
	}
	paths := [][]string{}
	queue := [][]string{}
	queue = append(queue, []string{"start"})
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		if path[len(path)-1] == "end" {
			paths = append(paths, path)
			continue
		}
		for _, v := range caves[path[len(path)-1]] {
			if v == "start" {
				continue
			}
			if !islower(v) || islower(v) && !contains(path, v) || islower(v) && contains(path, v) && !hasDoubleSmall(path) {
				queue = append(queue, copyAndAppend(path, v))
			}
		}
	}
	fmt.Println(len(paths))
}
