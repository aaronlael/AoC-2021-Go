package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	lines, err := readinput(`C:\Users\alael\Documents\Go\AoC\2021\data\inp_d1.txt`)
	if err != nil {
		return
	}
	fmt.Println(ascendtrio(lines))
	fmt.Println(time.Since(start))
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

func ascendtrio(lines []string) int {
	ascendcount := 0
	for i, _ := range lines {
		if i < len(lines)-3 {
			c := thethree(lines[i], lines[i+1], lines[i+2])
			n := thethree(lines[i+1], lines[i+2], lines[i+3])
			if n > c {
				ascendcount += 1
			}
		}
	}
	return ascendcount
}

func thethree(a string, b string, c string) int {
	ia, _ := strconv.Atoi(a)
	ib, _ := strconv.Atoi(b)
	ic, _ := strconv.Atoi(c)
	return ia + ib + ic
}
