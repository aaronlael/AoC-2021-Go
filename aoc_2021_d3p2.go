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
	oxy := lines
	co2 := lines
	for i, _ := range lines[0] {
		if len(oxy) > 1 {
			oxy = binpos(oxy, i, "o2")
		}
		if len(co2) > 1 {
			co2 = binpos(co2, i, "co2")
		}
		if len(co2) == 1 && len(oxy) == 1 {
			ioxy, _ := strconv.ParseInt(oxy[0], 2, 64)
			ico2, _ := strconv.ParseInt(co2[0], 2, 64)
			fmt.Println(ioxy * ico2)
		}
	}

	fmt.Println(time.Since(start))
}

func binpos(lines []string, i int, mode string) []string {
	var out []string
	var vals []string
	for _, line := range lines {
		vals = append(vals, string(line[i]))
	}
	jvals := strings.Join(vals, "")
	if mode == "o2" {
		if strings.Count(jvals, "0") > strings.Count(jvals, "1") {
			out = filt(lines, "0", i)
		} else {
			out = filt(lines, "1", i)
		}
	} else {
		if strings.Count(jvals, "0") > strings.Count(jvals, "1") {
			out = filt(lines, "1", i)
		} else {
			out = filt(lines, "0", i)
		}
	}
	return out
}

func filt(lines []string, f string, i int) []string {
	var tout []string
	for _, l := range lines {
		if string(l[i]) == f {
			tout = append(tout, l)
		}
	}
	return tout
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
