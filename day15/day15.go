//super hacky solution to this problem
//many thanks to https://github.com/fzipp/astar for the base algorithm and data structures that I have butchered herein
package main

import (
	"container/heap"
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

type item struct {
	value    interface{}
	priority float64
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*item))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type node interface{}

//Graph is a graph
type Graph interface {
	Neighbours(n node) []node
}

//CostFunc is a costfunc
type CostFunc func(a node, b node, m map[image.Point]int) float64

//Path is a path
type Path []node

func newPath(start node) Path {
	return []node{start}
}

func (p Path) last() node {
	return p[len(p)-1]
}

func (p Path) cont(n node) Path {
	newPath := make([]node, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, n)
	return newPath
}

//Cost is a method
func (p Path) Cost(m map[image.Point]int) (c float64) {
	for i := 1; i < len(p); i++ {
		c += float64(m[p[i].(image.Point)])
	}
	return c
}

//FindPath is a method
func FindPath(g Graph, start, dest node, d, h CostFunc, m map[image.Point]int) Path {

	closed := make(map[node]bool)

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &item{value: newPath(start)})

	for pq.Len() > 0 {
		p := heap.Pop(pq).(*item).value.(Path)
		n := p.last()
		if closed[n] {
			continue
		}
		if n == dest {
			return p
		}
		closed[n] = true

		for _, nb := range g.Neighbours(n) {
			newPath := p.cont(nb)
			heap.Push(pq, &item{
				value:    newPath,
				priority: -(newPath.Cost(m) + h(nb, dest, m)),
			})
		}
	}
	return nil
}

type point struct {
	x, y int
}
type graph map[node][]node

func newGraph() graph {
	return make(map[node][]node)
}

func (g graph) link(a, b node) graph {
	g[a] = append(g[a], b)
	g[b] = append(g[b], a)
	return g
}
func (g graph) Neighbours(n node) []node {
	return g[n]
}
func check(p image.Point, m map[image.Point]int) bool {
	if _, ok := m[p]; ok {
		return true
	}
	return false
}
func addNeighbors(p image.Point, m map[image.Point]int) []image.Point {
	candidates := []image.Point{}
	up := image.Point{p.X, p.Y + 1}
	down := image.Point{p.X, p.Y - 1}
	left := image.Point{p.X - 1, p.Y}
	right := image.Point{p.X + 1, p.Y}
	candidates = append(candidates, up, down, left, right)
	neighbors := []image.Point{}
	for _, v := range candidates {
		if check(v, m) {
			neighbors = append(neighbors, v)
		}
	}
	return neighbors
}

func nodeDist(a node, b node, m map[image.Point]int) float64 {
	p := a.(image.Point)
	q := b.(image.Point)
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}
func tile(m map[image.Point]int, oY int, oX int) map[image.Point]int {
	tiledX := make(map[image.Point]int)
	tiled := make(map[image.Point]int)
	for i := 0; i < 5; i++ {
		for k, v := range m {
			newpoint := image.Point{k.X + oX*i, k.Y}
			newVal := 1
			if v+i <= 9 {
				newVal = v + i
			} else {
				newVal = v + i - 9
			}
			tiledX[newpoint] = newVal
		}
	}
	for i := 0; i < 5; i++ {
		for k, v := range tiledX {
			newpoint := image.Point{k.X, k.Y + oY*i}
			newVal := 1
			if v+i <= 9 {
				newVal = v + i
			} else {
				newVal = v + i - 9
			}
			tiled[newpoint] = newVal
		}
	}
	return tiled
}
func main() {
	allPoints := make(map[image.Point]int)
	bytes, _ := ioutil.ReadFile("../inputs/input15.txt")
	parsed := strings.Split(string(bytes), "\n")
	maxX := len(parsed[0])
	maxY := len(parsed)
	for y, row := range parsed {
		for x, cell := range row {
			thisnode := image.Point{x, y}
			allPoints[thisnode] = int(cell - '0')
		}
	}
	allPoints = tile(allPoints, maxX, maxY)
	g := newGraph()
	for point := range allPoints {
		neighbors := addNeighbors(point, allPoints)
		for _, neighbor := range neighbors {
			g.link(point, neighbor)
		}
	}

	a := image.Point{0, 0}
	d := image.Point{499, 499}
	start := time.Now()
	p := FindPath(g, a, d, nodeDist, nodeDist, allPoints)
	if p == nil {
		fmt.Println("No path found.")
		return
	}
	totalRisk := 0
	for _, n := range p {
		totalRisk += allPoints[n.(image.Point)]
	}
	fmt.Println(totalRisk - allPoints[image.Point{0, 0}])
	fmt.Println(time.Since(start))
}
