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
	lines, err := readinput(`C:\Users\alael\Documents\Go\AoC\2021\data\inp_d3.txt`)
	if err != nil {
		return
	}
	var gamma string
	for i, _ := range strings.Split(lines[0], "") {
		gamma += most(lines, i)
	}
	epsilon, _ := flipstring(gamma)
	en, _ := strconv.ParseInt(epsilon, 2, 64)
	gn, _ := strconv.ParseInt(gamma, 2, 64)
	fmt.Println(en * gn)
	fmt.Println(time.Since(start))
}

func most(lines []string, i int) string {
	count1, count0 := 0, 0
	for _, line := range lines {
		if string(line[i]) == "0" {
			count0 += 1
		} else {
			count1 += 1
		}
	}
	if count1 > count0 {
		return "1"
	} else {
		return "0"
	}
}

func flipstring(val string) (string, error) {
	var out string
	for _, x := range strings.Split(val, "") {
		if x == "0" {
			out += "1"
		} else {
			out += "0"
		}
	}
	return out, nil
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
