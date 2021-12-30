package main

import (
	"fmt"
	"math"
)

type player struct {
	space int
	score int
}

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

func main() {
	//part one
	p1 := player{10, 0}
	p2 := player{2, 0}
	die := die{0, 0}
	for {
		p1.turn(&die)
		if p1.score >= 1000 {
			break
		}
		p2.turn(&die)
		if p2.score >= 1000 {
			break
		}
	}
	losingScore := int(math.Min(float64(p1.score), float64(p2.score)))
	fmt.Println("Part one: ", losingScore*die.total)

	//part two

}
