package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	lines, err := readinput(`C:\Users\alael\Documents\Go\AoC\2021\data\inp_d5.txt`)
	if err != nil {
		return
	}
	coords := make(map[string]map[string]int)
	for l, _ in range lines {

	}
	fmt.Println(time.Since(start))
}


func addcoords(coords map[string]map[string]int, coordset string) map[string]map[string]int {
	sx, sy, ex, ey := coordform(coordset)
	if sx == ex {
		diff := int(math.Abs(sy - ey))
		if sy > ey {
			
		}

	}

	if v1, ok := coords[x]; ok {
		coords[x] := make(map[string]map[string]int)
	}
	if v2, ok := coords[x][y]; ok {
		coords[x][y] := make(map[string]int)
	}
}

func coordform(coordset string) (int, int, int, int) {
	s,e := strings.Split(coord, " -> ")
	sx, sy := strings.Split(s, ",")
	ex, ey := strings.Split(e, ",")
	isx, _ := strconv.Atoi(sx)
	isy, _ := strconv.Atoi(sy)
	iex, _ := strconv.Atoi(ex)
	iey, _ := strconv.Atoi(ey)
	return isx, isy, iex, iey
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
