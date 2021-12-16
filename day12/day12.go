//took me far too long, but it's finally done!

package main

import (
	"fmt"
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

func main() {

	inputString := `he-JK
wy-KY
pc-XC
vt-wy
LJ-vt
wy-end
wy-JK
end-LJ
start-he
JK-end
pc-wy
LJ-pc
at-pc
xf-XC
XC-he
pc-JK
vt-XC
at-he
pc-he
start-at
start-XC
at-LJ
vt-JK`
	caves := make(map[string][]string)
	lines := strings.Split(inputString, "\n")
	for _, v := range lines {
		entrySplit := strings.Split(v, "-")
		caves[entrySplit[0]] = append(caves[entrySplit[0]], entrySplit[1])
		caves[entrySplit[1]] = append(caves[entrySplit[1]], entrySplit[0])
	}
	start := "start"
	end := "end"
	paths := make(map[string]int)
	queue := [][]string{}
	queue = append(queue, []string{start})
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		if path[len(path)-1] == end {
			collapse := ""
			for _, v := range path {
				collapse = collapse + v + "."
			}
			paths[collapse]++
			continue
		}
		dsmall := hasDoubleSmall(path)
		for _, v := range caves[path[len(path)-1]] {
			if v == "start" {
				continue
			}
			if !islower(v) {
				n := make([]string, len(path))
				copy(n, path)
				n = append(n, v)
				queue = append(queue, n)
			}

			if islower(v) && !contains(path, v) {
				x := make([]string, len(path))
				copy(x, path)
				x = append(x, v)
				queue = append(queue, x)
			}

			if islower(v) && contains(path, v) && !dsmall {
				y := make([]string, len(path))
				copy(y, path)
				y = append(y, v)
				queue = append(queue, y)
			}
		}

	}
	fmt.Println(len(paths))
}
