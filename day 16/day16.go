//work in progress
package main

import (
	"fmt"
	"math"
	"strconv"
)

func hexToBin(hex string) string {
	finalStr := ""

	for _, v := range hex {
		ui, _ := strconv.ParseUint(string(v), 16, 64)
		finalStr += fmt.Sprintf("%04b", ui)
	}
	return finalStr
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func convertBinaryToDecimal(n string) int {
	number, _ := strconv.Atoi(n)
	decimal := 0
	counter := 0.0
	remainder := 0

	for number != 0 {
		remainder = number % 10
		decimal += remainder * int(math.Pow(2.0, counter))
		number = number / 10
		counter++
	}
	return decimal
}

func parseStrToPacket(input string) packet {
	version := convertBinaryToDecimal(input[:3])
	id := convertBinaryToDecimal(input[3:6])
	fmt.Println(version, id)
	if id == 4 {
		inner := input[6:]
		return packet{version, id, 2, inner}
	}
	lengthType := input[6] - '0'
	subPacketLength := convertBinaryToDecimal(input[7:22])
	fmt.Println(subPacketLength)
	inner := input[22 : 22+subPacketLength]
	return packet{version, id, int(lengthType), inner}
}

func parseLiteral(p packet) (int, int) {
	trailing := len(p.inner) % 5
	literal := ""
	keepReading := true
	chunkStart := 0
	chunkStop := 5
	for keepReading {
		chunk := p.inner[chunkStart:chunkStop]
		fmt.Println(chunk)
		if chunk[0] == '0' {
			keepReading = false
		}
		literal += chunk[1:]
		chunkStart += 5
		chunkStop += 5
	}
	return convertBinaryToDecimal(literal), trailing
}

type packet struct {
	version    int
	id         int
	lengthType int
	inner      string
}

func main() {
	bin := hexToBin("38006F45291200")
	fmt.Println(bin)
	fmt.Println("110100010100101001000100100")
	binner := parseStrToPacket(bin).inner
	fmt.Println(parseLiteral(parseStrToPacket(binner)))
}
