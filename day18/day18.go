package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func snailAddition(s1 string, s2 string) string {
	return fmt.Sprintf("[%v,%v]", s1, s2)
}
func replaceAtIndex1(str string, replacement string, index int, numtoreplace int) string {
	return str[:index] + string(replacement) + str[index+numtoreplace:]
}
func sumSnail(s string) string {
	z := s
	re := regexp.MustCompile(`\[\d+,\d+\]`)
	for re.FindAllString(z, -1) != nil {
		t := re.FindAllString(z, -1)
		for _, v := range t {
			nre := regexp.MustCompile(`\d+`)
			l := nre.FindAllString(v, -1)
			a, _ := strconv.Atoi(l[0])
			b, _ := strconv.Atoi(l[1])
			total := a*3 + b*2
			g := strconv.Itoa(total)
			z = strings.ReplaceAll(z, v, g)
		}
	}
	return z
}
func checkExplode(fullstr string) (bool, []int, string) {
	re := regexp.MustCompile(`\[\d+,\d+\]`)
	depth := 0
	for i, v := range fullstr {
		if i == len(fullstr)-6 {
			break
		}
		check := re.Find([]byte(fullstr[i : i+6]))
		if v == '[' {
			depth++
		}
		if v == ']' {
			depth--
		}
		if depth >= 5 && check != nil {
			return true, []int{i, i + len(string(check))}, string(check)
		}
	}
	return false, nil, ""
}

func checkSplit(fullstr string) (bool, string) {
	re := regexp.MustCompile(`\d\d`)
	split := re.FindString(fullstr)
	if len(split) > 0 {
		return true, split
	}
	return false, ""
}

func explode(indexes []int, substr string, fullstr string) string {
	re := regexp.MustCompile(`\d+`)
	z := fullstr
	imatches := re.FindAllIndex([]byte(substr), -1)
	lint, _ := strconv.Atoi(substr[imatches[0][0]:imatches[0][1]])
	rint, _ := strconv.Atoi(substr[imatches[1][0]:imatches[1][1]])
	leftIndex := indexes[0]
	left := z[:leftIndex]
	lmatches := re.FindAllIndex([]byte(left), -1)

	if lmatches != nil {
		closest := lmatches[len(lmatches)-1]
		m := left[closest[0]:closest[1]]
		n, _ := strconv.Atoi(m)
		repl := strconv.Itoa(lint + n)
		left = replaceAtIndex1(left, repl, closest[0], len(m))
	}
	right := z[leftIndex+len(substr):]
	rmatches := re.FindAllIndex([]byte(right), -1)
	if rmatches != nil {
		closest := rmatches[0]
		m := right[closest[0]:closest[1]]
		n, _ := strconv.Atoi(m)
		repl := strconv.Itoa(rint + n)
		right = replaceAtIndex1(right, repl, closest[0], len(m))
	}
	z = left + "0" + right
	return z
}
func split(substr string, fullstr string) (string, string, int) {
	index := strings.Index(fullstr, substr)
	substrint, _ := strconv.Atoi(substr)
	a := strconv.Itoa(int(math.Floor(float64(substrint) / float64(2))))
	b := strconv.Itoa(int(math.Ceil(float64(substrint) / float64(2))))
	pair := fmt.Sprintf("[%v,%v]", a, b)
	return strings.Replace(fullstr, substr, pair, 1), pair, index
}
func checkDepth(j int, fullstr string) bool {
	depth := 0
	for i := 0; i <= j; i++ {
		if fullstr[i] == '[' {
			depth++
		}
		if fullstr[i] == ']' {
			depth--
		}
	}
	return depth >= 5
}
func reduce(str string) string {
	z := str
	for {
		noexplode := true
		nosplit := true
		a, b, c := checkExplode(z)
		if a {
			noexplode = false
			z = explode(b, c, z)
		} else {
			c, d := checkSplit(z)
			if c {
				nosplit = false
				f, s, i := split(d, z)
				z = f
				if checkDepth(i, z) {
					z = explode([]int{i, i + len(s)}, s, z)
				}

			}
		}
		if noexplode && nosplit {
			break
		}
	}
	return z
}
func main() {
	bytes, _ := ioutil.ReadFile("../inputs/input18.txt")
	lines := strings.Split(string(bytes), "\n")
	curr := lines[0]
	for i := 1; i < len(lines); i++ {
		sum := snailAddition(curr, lines[i])
		sum = reduce(sum)
		curr = sum
	}
	fmt.Println(sumSnail(curr))
	max := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			total, _ := strconv.Atoi(sumSnail(reduce(snailAddition(lines[i], lines[j]))))
			if total > max {
				max = total
			}
		}
	}
	fmt.Println(max)
}
