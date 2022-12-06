package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Find characters repetitions in a string, return false if a char is not unique in s
func isUnique(s string) bool {
	for _, r := range s {
		nb := 0
		for _, r2 := range s {
			if r == r2 {
				nb += 1
			}
			if nb > 1 {
				return false
			}
		}
	}
	return true
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

	// part1
	score := 0
	nbchar := 4
	for _, l := range lines {
		b := false
		i := 0
		for b == false {
			substr := l[i : i+nbchar]
			if isUnique(substr) {
				score += i + nbchar
				b = true
			}
			i++
		}
	}
	fmt.Println(score)

	// part2
	score = 0
	nbchar = 14
	for _, l := range lines {
		b := false
		i := 0
		for b == false {
			substr := l[i : i+nbchar]
			if isUnique(substr) {
				score += i + nbchar
				b = true
			}
			i++
		}
	}
	fmt.Println(score)
}
