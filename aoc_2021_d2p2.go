package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	lines, err := readinput(`C:\Users\alael\Documents\Go\AoC\2021\data\inp_d2.txt`)
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	y, x := 0, 0
	aim := 0
	for _, line := range lines {
		switch {
		case strings.Split(line, " ")[0] == "forward":
			xt, _ := strconv.Atoi(strings.Split(line, " ")[1])
			x += xt
			y += xt * aim
		case strings.Split(line, " ")[0] == "up":
			yt, _ := strconv.Atoi(strings.Split(line, " ")[1])
			aim -= yt
		case strings.Split(line, " ")[0] == "down":
			yt, _ := strconv.Atoi(strings.Split(line, " ")[1])
			aim += yt
		default:
			fmt.Println("You dun goofed A-A-ron")
		}
	}
	fmt.Println(x * y)
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
