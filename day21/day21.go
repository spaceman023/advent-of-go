//not mine, reverse engineering one from https://github.com/Torakushi
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	boardGame                  = &board{}
	cache                      = map[string][]int{}
	universesCreatedByLaunches = map[int]int{}
)

func main() {
	fmt.Println("DAY21")
	if err := process(); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func process() error {
	// First Part
	initBoardGame(&detDice{})
	fmt.Printf("Part one, the loosing score * rolls number: %d\n\n", boardGame.detPlay())

	// Second Part
	initBoardGame(nil)
	getUniversesCreateByLaunches()
	r := boardGame.quanticPlay()
	max, index := r[0], 1
	if r[1] > max {
		max = r[1]
		index = 2
	}

	fmt.Printf("Player1 wins on %d universe and Player2 wins on %d ... CONGRATZ Player%d, you win more !\n",
		r[0], r[1], index)
	return nil
}

func readDatas() error {
	boardGame = &board{}
	file, err := os.Open("../inputs/input21.txt")
	if err != nil {
		return fmt.Errorf("error while opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`Player (\d+) starting position: (\d+)`)
	p := []*player{}
	for scanner.Scan() {
		s := re.FindAllStringSubmatch(scanner.Text(), -1)[0]
		index, _ := strconv.Atoi(s[1])
		position, _ := strconv.Atoi(s[2])
		p = append(p, &player{
			index:    index,
			position: position,
		})
	}
	boardGame.players = p
	return nil
}

type board struct {
	dice    *detDice
	players []*player
	turn    int
}

func (b *board) detPlay() int {
	winnerIndex := 0
L:
	for {
		for _, player := range b.players {
			score := b.dice.rolls(3)
			player.updateScoreAndPosition(score)
			if player.score >= 1000 {
				fmt.Printf("Player %d Win with score: %d!\n", player.index, player.score)
				winnerIndex = player.index - 1
				break L
			}
		}
	}
	return b.players[(winnerIndex+1)%2].score * b.dice.rollCount
}

func (b *board) quanticPlay() []int {
	wins := make([]int, 2)
	for k, v := range universesCreatedByLaunches {
		oB := b.Copy()
		oB.players[b.turn-1].updateScoreAndPosition(k)
		if oB.players[b.turn-1].score >= 21 {
			wins[b.turn-1] += v
		} else {
			oB.turn = b.turn + 1
			if oB.turn > 2 {
				oB.turn = 1
			}
			//
			r, ok := cache[oB.getState()]
			if !ok {
				r = oB.quanticPlay()
				cache[oB.getState()] = r
			}

			wins[0] += v * r[0]
			wins[1] += v * r[1]
		}
	}
	return wins
}

type player struct {
	index    int
	position int
	score    int
}

func (p *player) updateScoreAndPosition(r int) {
	newP := p.position + r
	if newP%10 == 0 {
		p.position = 10
	} else {
		p.position = newP % 10
	}
	p.score += p.position
}

type detDice struct {
	current   int
	rollCount int
}

func (d *detDice) rolls(l int) int {
	sum := 0
	for i := 1; i <= l; i++ {
		d.current++
		d.rollCount++
		if d.current == 101 {
			d.current = 1
		}
		sum += d.current
	}
	return sum
}

func initBoardGame(d *detDice) error {
	if err := readDatas(); err != nil {
		return err
	}
	boardGame.dice = d
	boardGame.turn = 1
	return nil
}

func (b *board) getState() string {
	return fmt.Sprintf("T%d-po%d-pt%d-so%d-st%d",
		b.turn,
		b.players[0].position,
		b.players[1].position,
		b.players[0].score,
		b.players[1].score,
	)
}

func (b *board) Copy() *board {
	return &board{
		turn: b.turn,
		dice: b.dice,
		players: []*player{
			{
				position: b.players[0].position,
				score:    b.players[0].score,
				index:    b.players[0].index,
			},
			{
				position: b.players[1].position,
				score:    b.players[1].score,
				index:    b.players[1].index,
			},
		},
	}
}

func getUniversesCreateByLaunches() {
	universesCreatedByLaunches = map[int]int{}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				universesCreatedByLaunches[i+k+j]++
			}
		}
	}
}
