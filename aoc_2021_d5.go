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
	lines, err := readinput(`C:\Users\alael\Documents\Go\AoC\2021\data\inp_d5.txt`)
	if err != nil {
		return
	}

	coords := make(map[string]map[string]int)
	for _, l := range lines {
		coords = addcoords(coords, l)
	}
	/// part 1
	c := 0
	for k, _ := range coords {
		for _, v := range coords[k] {
			if v > 1 {
				c += 1
			}
		}
	}
	fmt.Println("part 1:", c)
	/// part 2
	for _, l := range lines {
		coords = handlediags(coords, l)
	}
	c = 0
	for k, _ := range coords {
		for _, v := range coords[k] {
			if v > 1 {
				c += 1
			}
		}
	}
	fmt.Println("part 2:", c)
	fmt.Println(time.Since(start))
}

func addcoords(coords map[string]map[string]int, coordset string) map[string]map[string]int {
	sx, sy, ex, ey := coordform(coordset)
	var offset int
	if sx == ex {
		if sy > ey {
			offset = -1
		} else {
			offset = 1
		}
		for sy != ey+offset {
			coords = mapadd(coords, sx, sy)
			sy += offset
		}
	}
	if sy == ey {
		if sx > ex {
			offset = -1
		} else {
			offset = 1
		}
		for sx != ex+offset {
			coords = mapadd(coords, sx, sy)
			sx += offset
		}
	}
	return coords
}
func handlediags(coords map[string]map[string]int, coordset string) map[string]map[string]int {
	sx, sy, ex, ey := coordform(coordset)
	var xoffset int
	var yoffset int
	if sx != ex && sy != ey {
		if sx > ex {
			xoffset = -1
		} else {
			xoffset = 1
		}
		if sy > ey {
			yoffset = -1
		} else {
			yoffset = 1
		}
		for sx != ex+xoffset {
			coords = mapadd(coords, sx, sy)
			sx += xoffset
			sy += yoffset
		}
	}
	return coords
}

func mapadd(coords map[string]map[string]int, x int, y int) map[string]map[string]int {
	strx := strconv.Itoa(x)
	stry := strconv.Itoa(y)
	if _, ok := coords[strx]; ok {
		if _, ok := coords[strx][stry]; ok {
			coords[strx][stry] += 1
		} else {
			coords[strx][stry] = 1
		}
	} else {
		coords[strx] = make(map[string]int)
		coords[strx][stry] = 1
	}
	return coords
}

func coordform(coordset string) (int, int, int, int) {
	t1 := strings.Split(coordset, " -> ")
	s := t1[0]
	e := t1[1]
	ts := strings.Split(s, ",")
	sx := ts[0]
	sy := ts[1]
	te := strings.Split(e, ",")
	ex := te[0]
	ey := te[1]
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
