package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// scme := map[string]int{
	// 	"X": 1,
	// 	"Y": 2,
	// 	"Z": 3,
	// }

	scop := map[string]int{
		"A": 1, // X ROCK
		"B": 2, // Y PAPER
		"C": 3, // Z SCISOR
	}

	score := 0
	for _, l := range lines {
		c := strings.Split(l, " ")

		if c[1] == "X" {
			v := scop[c[0]]
			i := v + 2
			if i > 3 {
				i -= 3
			}
			score += i
		} else if c[1] == "Y" {
			score += scop[c[0]]
			score += 3
		} else if c[1] == "Z" {
			v := scop[c[0]]
			i := v - 2
			if i <= 0 {
				i += 3
			}
			score += i
			score += 6
		}
	}
	fmt.Println("score", score)

}
