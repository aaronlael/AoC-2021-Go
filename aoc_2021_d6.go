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
	lines, err := readinput(`C:\Users\alael\Documents\Go\AoC\2021\data\inp_d6.txt`)
	if err != nil {
		return
	}
	/// part 1
	/// days := 80
	/// part 2
	days := 256
	d := make(map[string]int)
	for i := 0; i < 9; i++ {
		d[strconv.Itoa(i)] = 0
	}
	for _, x := range strings.Split(lines[0], ",") {
		d = dictpopulate(x, 1, d)
	}
	fmt.Println(d)
	t := make(map[string]int)
	for i := 0; i < days; i++ {
		for j := 0; j < 9; j++ {
			t[strconv.Itoa(j)] = 0
		}
		for k, _ := range d {
			if k == "0" {
				t["6"] += d[k]
				t["8"] += d[k]
			} else {
				ik, _ := strconv.Atoi(k)
				ik -= 1
				tk := strconv.Itoa(ik)
				t[tk] += d[k]
			}
		}
		for k, v := range t {
			d[k] = v
		}
	}
	c := 0
	for _, v := range d {
		c += v
	}
	fmt.Println(c)
	fmt.Println(time.Since(start))
}

func dictpopulate(key string, val int, d map[string]int) map[string]int {
	td := make(map[string]int)
	for k, v := range d {
		td[k] = v
	}
	td[key] += val
	return td
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
