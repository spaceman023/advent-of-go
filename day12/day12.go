package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

var inputFile = flag.String("inputFile", "../inputs/input12sample.txt", "Relative file path to use as input.")

type graph struct {
	nodes []node
	edges []edge
}

type node struct {
	name      string
	outgoings []edge
	incomings []edge
}

type edge struct {
	name string
	from node
	to   node
}

func containsNode(s []node, n node) bool {
	for _, v := range s {
		if v.name == n.name {
			return true
		}
	}
	return false
}
func containsEdge(s []edge, n edge) bool {
	for _, v := range s {
		if v.name == n.name {
			return true
		}
	}
	return false
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
func main() {
	flag.Parse()
	bytes, _ := ioutil.ReadFile(*inputFile)
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	var g = graph{[]node{}, []edge{}}
	fmt.Println(lines)
	for _, v := range lines {
		split := strings.Split(v, "-")
		nodeA := node{split[0], []edge{}, []edge{}}
		nodeB := node{split[1], []edge{}, []edge{}}
		edgeA := edge{split[0] + split[1], nodeA, nodeB}
		if !containsNode(g.nodes, nodeA) {
			g.nodes = append(g.nodes, nodeA)
		}
		if !containsNode(g.nodes, nodeB) {
			g.nodes = append(g.nodes, nodeB)
		}
		if !containsEdge(g.edges, edgeA) {
			g.edges = append(g.edges, edgeA)
		}
	}
}
