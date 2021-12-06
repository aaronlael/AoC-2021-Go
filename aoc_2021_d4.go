package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readinput(`C:\Users\alael\Documents\Go\AoC\2021\data\inp_d4_1.txt`)
	if err != nil {
		return
	}
	var winvals []int
	cards := formcards(lines)
	calls, err := readinput(`C:\Users\alael\Documents\Go\AoC\2021\data\inp_d4_2.txt`)
	if err != nil {
		return
	}
	calls = strings.Split(calls[0], ",")
	var winners [][][]string
	var twinvals []int
	for _, call := range calls {
		cards = update(cards, call)
		winners = hoorayforthewinners(cards)
		cards, twinvals, winners = winhandler(cards, winners, call)
		winvals = append(winvals, twinvals...)
		if cards == nil {
			fmt.Println(winvals[0])
			fmt.Println(winvals[len(winvals)-1])
			break
		}

	}
}

func winhandler(cards [][][]string, winners [][][]string, call string) ([][][]string, []int, [][][]string) {
	var winvals []int
	icall, _ := strconv.Atoi(call)
	for _, w := range winners {
		for i, c := range cards {
			if slicecompare(c, w) {
				winvals = append(winvals, cardval(w)*icall)
				if len(cards) > 1 {
					if i < len(cards)-1 {
						cards = append(cards[:i], cards[i+1:]...)
					} else {
						cards = cards[:i]
					}
				} else {
					cards = nil
				}
			}
		}
	}
	winners = nil
	return cards, winvals, winners
}

func slicecompare(card1 [][]string, card2 [][]string) bool {
	for i, _ := range card1 {
		for j, _ := range card1[i] {
			if card1[i][j] != card2[i][j] {
				return false
			}
		}
	}
	return true
}

func cardval(card [][]string) int {
	v := 0
	for i, _ := range card {
		for j, _ := range card[i] {
			if card[i][j] != "x" {
				t, _ := strconv.Atoi(card[i][j])
				v += t
			}
		}
	}
	return v
}

func hoorayforthewinners(cards [][][]string) [][][]string {
	var winners [][][]string
	for i, _ := range cards {
		for j, _ := range cards[i] {
			if count(cards[i][j], "x") == 5 {
				winners = append(winners, cards[i])
				break
			}
		}
		if count(cards[i][0], "x") > 0 {
			for k, _ := range cards[i][0] {
				if cards[i][0][k] == "x" {
					c := 1
					for l := 1; l < 5; l++ {
						if cards[i][l][k] == "x" {
							c += 1
						}
					}
					if c == 5 {
						winners = append(winners, cards[i])
					}

				}
			}
		}
	}
	return winners
}

func update(cards [][][]string, call string) [][][]string {
	for i, _ := range cards {
		for j, _ := range cards[i] {
			for k, _ := range cards[i][j] {
				if cards[i][j][k] == call {
					cards[i][j][k] = "x"
				}
			}
		}
	}
	return cards
}

func formcards(lines []string) [][][]string {
	var cards [][][]string
	var card [][]string
	for _, line := range lines {
		if line == "" {
			continue
		}
		line = strings.Replace(line, "  ", " ", 0)
		c := strings.Split(line, " ")
		var c2 []string
		for _, x := range c {
			if x != "" {
				c2 = append(c2, x)
			}
		}
		card = append(card, c2)
		if len(card) == 5 {
			cards = append(cards, card)
			card = nil
		}

	}
	return cards
}

func count(line []string, val string) int {
	i := 0
	for _, x := range line {
		if x == val {
			i += 1
		}
	}
	return i
}

func readinput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}
