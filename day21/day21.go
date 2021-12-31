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

//part two types and methods

type game struct {
	p1       player
	p2       player
	turn     bool
	gameover bool
}

func (p *player) nextTurn(move int) {
	p.space = (p.space-1+move)%10 + 1
	p.score += p.space
}
func (g game) iterate(move int, counter1 *int, counter2 *int) game {
	if g.turn {
		g.p1.nextTurn(move)
	} else {
		g.p2.nextTurn(move)
	}
	g.turn = !g.turn
	if g.p1.score >= 21 {
		*counter1++
		g.gameover = true
		return g
	}
	if g.p2.score >= 21 {
		*counter2++
		g.gameover = true
		return g
	}
	return g
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
	dieRollProbs := make(map[int]int)
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			for z := 1; z <= 3; z++ {
				dieRollProbs[x+y+z]++
			}
		}
	}
	fmt.Println(dieRollProbs)
	queue := []game{}
	starterGame := game{
		player{10, 0},
		player{2, 0},
		true,
		false,
	}
	counter1 := 0
	counter2 := 0
	queue = append(queue, starterGame)
	for len(queue) > 0 {
		if queue[0].gameover {
			queue = queue[1:]
			continue
		}
		fmt.Println(len(queue), counter1, counter2)

		for k, v := range dieRollProbs {
			for i := 0; i < v; i++ {
				queue = append(queue, queue[0].iterate(k, &counter1, &counter2))
			}
		}
		queue = queue[1:]
	}
	fmt.Println(counter1, counter2)
}
