package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	//part1
	score := 0
	for _, l := range lines {
		pairs := strings.Split(l, ",")
		var e1, e2 = [2]int{}, [2]int{}

		p := strings.Split(pairs[0], "-")
		e1[0], _ = strconv.Atoi(p[0])
		e1[1], _ = strconv.Atoi(p[1])

		p = strings.Split(pairs[1], "-")
		e2[0], _ = strconv.Atoi(p[0])
		e2[1], _ = strconv.Atoi(p[1])

		if (e1[0] >= e2[0] && e1[1] <= e2[1]) || (e2[0] >= e1[0] && e2[1] <= e1[1]) {
			score++
		}
	}
	fmt.Println(score)

	//part2
	score = 0
	for _, l := range lines {
		pairs := strings.Split(l, ",")
		var e1, e2 = [2]int{}, [2]int{}

		p := strings.Split(pairs[0], "-")
		e1[0], _ = strconv.Atoi(p[0])
		e1[1], _ = strconv.Atoi(p[1])

		p = strings.Split(pairs[1], "-")
		e2[0], _ = strconv.Atoi(p[0])
		e2[1], _ = strconv.Atoi(p[1])

		for i := e1[0]; i <= e1[1]; i++ {
			if i >= e2[0] && i <= e2[1] {
				score++
				break
			}
		}
	}
	fmt.Println(score)
}
