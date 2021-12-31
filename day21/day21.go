package main

import (
	"fmt"
)

type player struct {
	space int
	score int
}

//part one types and methods
type die struct {
	current int
	total   int
}

func (p *player) turn(d *die) {
	totalMove := 0
	for i := 0; i < 3; i++ {
		totalMove += d.next()
	}
	p.space = (p.space-1+totalMove)%10 + 1
	p.score += p.space
}

func (d *die) next() int {
	d.current++
	d.total++
	if d.current > 100 {
		d.current = d.current % 100
	}
	return d.current
}
func weirdMod(i int) int {
	for i > 10 {
		i -= 10
	}
	return i
}

var cache = map[string][]int{}

//part two types and methods
func compute(p1 int, p1score int, p2 int, p2score int) []int {
	wins := make([]int, 2)
	dieRollProbs := make(map[int]int)
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			for z := 1; z <= 3; z++ {
				dieRollProbs[x+y+z]++
			}
		}
	}
	for _, v := range dieRollProbs {
		newp1 := weirdMod(p1 + v)
		newp1score := p1score + newp1
		if newp1score >= 21 {
			wins[0] += v
		} else {
			tmpwins := []int{0, 0}
			tmpwins = compute(
				p2,
				p2score,
				newp1,
				newp1score,
			)
			state1 := fmt.Sprintf("%d%d%d%d",
				p2,
				p2score,
				newp1,
				newp1score,
			)
			cache[state1] = tmpwins
			wins[0] += tmpwins[1] * v
			wins[1] += tmpwins[0] * v
		}

	}
	return wins
}

func main() {
	//part one
	p1 := player{10, 0}
	p2 := player{2, 0}
	die := die{0, 0}
	losingScore := 0
	for {
		p1.turn(&die)
		if p1.score >= 1000 {
			losingScore = p2.score
			break
		}
		p2.turn(&die)
		if p2.score >= 1000 {
			losingScore = p1.score
			break
		}
	}
	fmt.Println("Part one: ", losingScore*die.total)
	//part two
	answer := compute(4, 0, 8, 0)
	fmt.Println(answer[0], answer[1])
	fmt.Println(444356092776315, 341960390180808)
}
