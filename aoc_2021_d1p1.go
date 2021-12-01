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
	fmt.Println(ascend(lines))
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

func ascend(lines []string) int {
	ascendcount := 0
	for i, _ := range lines {
		if i != 0 {
			end, _ := strconv.Atoi(lines[i])
			start, _ := strconv.Atoi(lines[i-1])
			if end > start {
				ascendcount += 1
			}
		}
	}
	return ascendcount
}
