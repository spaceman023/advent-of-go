//brute force!

package main

import "fmt"

type target struct {
	x1, x2, y1, y2 int
}
type pos struct {
	x, y int
}

func fire(xv int, yv int, t target) (bool, int) {
	p := pos{0, 0}
	maxY := 0
	for i := 0; i < 1500; i++ {
		if p.y > maxY {
			maxY = p.y
		}
		p.x += xv
		p.y += yv
		if p.x >= t.x1 && p.x <= t.x2 && p.y >= t.y1 && p.y <= t.y2 {
			p.y += yv
			return true, maxY
		}

		if xv != 0 {
			if xv > 0 {
				xv--
			} else {
				xv++
			}
		}
		yv--
	}
	return false, 0
}
func main() {
	//fuck parsing we're hardcoding
	t := target{206, 250, -105, -57}
	topHeight := 0
	successVelocities := make(map[pos]bool)
	for x := 500; x >= 0; x-- {
		for y := -500; y <= 500; y++ {
			success, height := fire(x, y, t)
			if success {
				successVelocities[pos{x, y}] = true
				if height > topHeight {
					topHeight = height
				}
			}
		}
	}
	fmt.Println(topHeight, len(successVelocities))
}
