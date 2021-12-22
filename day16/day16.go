//adapted to golang pretty much entirely from anthony's python solution on anthony writes code https://www.youtube.com/watch?v=hq7j0CVatoA
//learned a ton about pointers, recursion, and bit manipulation!

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type packet struct {
	version int
	typeID  int
	n       int
	packets []packet
}

func read(numberToRead int, inputString string, currentIndex *int) (r int) {
	chunk := inputString[*currentIndex : *currentIndex+numberToRead]
	ret, _ := strconv.ParseInt(chunk, 2, 64)
	*currentIndex += numberToRead
	return int(ret)
}

func parsePacket(i *int, inputString string) packet {
	version := read(3, inputString, i)
	typeID := read(3, inputString, i)

	if typeID == 4 {
		chunk := read(5, inputString, i)
		n := chunk & 0b1111
		for chunk&0b10000 > 0 {
			chunk = read(5, inputString, i)
			n <<= 4
			n += chunk & 0b1111
		}
		return packet{version, typeID, n, nil}
	}
	mode := read(1, inputString, i)
	if mode == 0 {
		bitsLength := read(15, inputString, i)
		j := *i
		*i = *i + bitsLength
		packets := []packet{}
		for j < *i {
			packet := parsePacket(&j, inputString)
			packets = append(packets, packet)
		}
		return packet{version, typeID, 0, packets}
	}
	subpackets := read(11, inputString, i)
	packets := []packet{}
	for y := 0; y < subpackets; y++ {
		packet := parsePacket(i, inputString)
		packets = append(packets, packet)
	}
	return packet{version, typeID, 0, packets}
}

func val(p packet) int {
	switch p.typeID {
	case 0:
		sum := 0
		for _, v := range p.packets {
			sum += val(v)
		}
		return sum
	case 1:
		res := 1
		for _, v := range p.packets {
			res *= val(v)
		}
		return res
	case 2:
		min := 1<<63 - 1
		for _, v := range p.packets {
			if val(v) < min {
				min = val(v)
			}
		}
		return min
	case 3:
		max := 0
		for _, v := range p.packets {
			if val(v) > max {
				max = val(v)
			}
		}
		return max
	case 4:
		return p.n
	case 5:
		if val(p.packets[0]) > val(p.packets[1]) {
			return 1
		}
		return 0
	case 6:
		if val(p.packets[0]) < val(p.packets[1]) {
			return 1
		}
		return 0
	case 7:
		if val(p.packets[0]) == val(p.packets[1]) {
			return 1
		}
		return 0
	}
	return 0
}

func hexToBin(hex string) (string, error) {
	ui, _ := strconv.ParseUint(hex, 16, 64)
	format := fmt.Sprintf("%%0%db", len(hex)*4)
	return fmt.Sprintf(format, ui), nil
}

func main() {
	bytes, _ := ioutil.ReadFile("../inputs/input16.txt")
	parsedBin := ""
	for _, hex := range bytes {
		bin, _ := hexToBin(string(hex))
		parsedBin += bin
	}
	i := 0
	packet1 := parsePacket(&i, parsedBin)
	todo := []packet{packet1}
	total := 0
	for len(todo) > 0 {
		item := todo[len(todo)-1]
		total += item.version
		todo = todo[0 : len(todo)-1]
		todo = append(todo, item.packets...)
	}
	//part one
	fmt.Println(total)
	//part two
	fmt.Println(val(packet1))
}
