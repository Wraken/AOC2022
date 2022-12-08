package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func checkUP(treeMap [][]int, x, y int) (bool, int) {
	v := treeMap[x][y]

	d := 0
	for i := x - 1; i >= 0; i-- {
		d++
		if treeMap[i][y] >= v {
			return false, d
		}
	}
	fmt.Println("tree", x, y, "visible from up")
	return true, d
}

func checkDOWN(treeMap [][]int, x, y int) (bool, int) {
	v := treeMap[x][y]

	d := 0
	for i := x + 1; i < len(treeMap); i++ {
		d++
		if treeMap[i][y] >= v {
			return false, d
		}
	}
	fmt.Println("tree", x, y, "visible from down")

	return true, d
}

func checkRIGHT(treeMap [][]int, x, y int) (bool, int) {
	v := treeMap[x][y]

	d := 0
	for i := y + 1; i < len(treeMap[y]); i++ {
		d++
		if treeMap[x][i] >= v {
			return false, d
		}
	}
	fmt.Println("tree", x, y, "visible from right")

	return true, d
}

func checkLEFT(treeMap [][]int, x, y int) (bool, int) {
	v := treeMap[x][y]

	d := 0
	for i := y - 1; i >= 0; i-- {
		d++
		if treeMap[x][i] >= v {
			return false, d
		}
	}
	fmt.Println("tree", x, y, "visible from left")

	return true, d
}

func checkTreePosition(treeMap [][]int, x, y int) bool {
	b := false
	if e, _ := checkUP(treeMap, x, y); e {
		b = true
	} else if e, _ := checkDOWN(treeMap, x, y); e {
		b = true
	} else if e, _ := checkLEFT(treeMap, x, y); e {
		b = true
	} else if e, _ := checkRIGHT(treeMap, x, y); e {
		b = true
	}
	return b
}

func calcViewScore(treeMap [][]int, x, y int) int {
	score := 0
	_, d := checkUP(treeMap, x, y)
	score = d
	_, d = checkDOWN(treeMap, x, y)
	score *= d
	_, d = checkRIGHT(treeMap, x, y)
	score *= d
	_, d = checkLEFT(treeMap, x, y)
	score *= d

	return score
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(r)

	scan.Split(bufio.ScanLines)

	lines := []string{}

	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	treeMap := [][]int{}
	for _, l := range lines {
		trow := []int{}
		for _, r := range l {
			i, _ := strconv.Atoi(string(r))
			trow = append(trow, i)
		}
		treeMap = append(treeMap, trow)
	}

	//part1
	score := 0
	for x, treeLine := range treeMap {
		for y := range treeLine {

			// Side check
			if x == 0 || x == len(treeMap)-1 {
				score++
			} else if y == 0 || y == len(treeLine)-1 {
				score++
			} else if checkTreePosition(treeMap, x, y) {
				score++
			}
		}
	}
	fmt.Println(score)

	//part2
	bestScore := 0
	for x, treeLine := range treeMap {
		for y := range treeLine {
			s := calcViewScore(treeMap, x, y)
			if s > bestScore {
				bestScore = s
			}
		}
	}
	fmt.Println(bestScore)
}
