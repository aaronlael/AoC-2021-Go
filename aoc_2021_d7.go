package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	lines, err := readinput(`C:\Users\alael\Documents\Go\AoC\2021\data\inp_d7.txt`)
	if err != nil {
		return
	}
	crabmap := assigncrabs(lines[0])
	/// regular flag for p1, anything else for p2
	cr := (crabrave("series", crabmap))
	sort.Ints(cr)
	fmt.Println(cr[0])
	fmt.Println(time.Since(start))
}

func assigncrabs(crabs string) map[string]int {
	cm := make(map[string]int)
	c := strings.Split(crabs, ",")
	for _, x := range c {
		if _, ok := cm[x]; !(ok) {
			cm[x] = 0
		}
		cm[x] += 1
	}
	return cm
}

func crabrave(mode string, crabs map[string]int) []int {
	var out []int
	for i := 0; i < 10000; i++ {
		var c int
		for ke, va := range crabs {
			ike, _ := strconv.Atoi(ke)
			if mode == "regular" {
				c += int(math.Abs(float64(ike-i))) * va
			} else {
				big := math.Abs(float64(ike - i))
				c += int(((1.0+big)/2.0)*big) * va
			}
		}
		out = append(out, c)
	}
	return out
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
