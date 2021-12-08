package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	lines, err := readinput(`C:\Users\alael\Documents\Go\AoC\2021\data\inp_d8.txt`)
	if err != nil {
		return
	}
	var out []string
	for _, l := range lines {
		l0, l1 := lineparser(l)
		out = append(out, decrypter(l0, l1))
	}
	fmt.Println(len(out))
	fmt.Println(time.Since(start))
}

func decrypter(l0 []string, l1 []string) string {
	key := make(map[string][]string)
	for _, v := range l0 {
		if len(v) == 2 {
			key["1"] = sort.StringSlice(strings.Split(v, ""))
		}
		if len(v) == 3 {
			key["7"] = sort.StringSlice(strings.Split(v, ""))
		}
		if len(v) == 4 {
			key["4"] = sort.StringSlice(strings.Split(v, ""))
		}
		if len(v) == 8 {
			key["8"] = sort.StringSlice(strings.Split(v, ""))
		}
	}
	for _, v := range l0 {
		if len(v) == 5 {
			s := strings.Split(v, "")
			if intersection(key["1"], s) == 2 {
				key["3"] = sort.StringSlice(s)
			} else if intersection(key["4"], s) == 3 {
				key["5"] = sort.StringSlice(s)
			} else {
				key["2"] = sort.StringSlice(s)
			}
		} else if len(v) == 6 {
			s := strings.Split(v, "")
			if intersection(key["1"], s) == 1 {
				key["6"] = sort.StringSlice(s)
			} else if intersection(key["4"], s) == 4 {
				key["9"] = sort.StringSlice(s)
			} else {
				key["0"] = sort.StringSlice(s)
			}
		}
	}
	return decoder(key, l1)
}

func decoder(key map[string][]string, v1 []string) string {
	var out string
	for _, x := range v1 {
		tx := sort.StringSlice(strings.Split(x, ""))
		for k, v := range key {
			if slicecompare(tx, v) {
				out += k
			}
		}
	}
	return out
}

func lineparser(line string) ([]string, []string) {
	t := strings.Split(line, " | ")
	t0 := strings.Split(t[0], " ")
	t1 := strings.Split(t[1], " ")
	return t0, t1
}

func intersection(v1 []string, v2 []string) int {
	c := 0
	for _, i := range v1 {
		for _, j := range v2 {
			if i == j {
				c += 1
			}
		}
	}
	return c
}

func slicecompare(card1 []string, card2 []string) bool {
	for i, _ := range card1 {
		for j, _ := range card1[i] {
			if card1[i][j] != card2[i][j] {
				return false
			}
		}
	}
	return true
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
