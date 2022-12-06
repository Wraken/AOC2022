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

	var ru rune
	ru = 'a'
	prio := map[string]int{}
	i := 1
	for ru <= 'z' {
		prio[string(ru)] = i
		ru += 1
		i++
	}
	ru = 'A'
	for ru <= 'Z' {
		prio[string(ru)] = i
		ru += 1
		i++
	}
	fmt.Println(prio)

	grp := [][]string{}
	nt := []string{}
	for y, l := range lines {
		nt = append(nt, l)
		if (y+1)%3 == 0 {
			grp = append(grp, nt)
			nt = []string{}
		}
	}

	score := 0
	for _, l := range lines {
		l1 := l[:len(l)/2]
		l2 := l[len(l)/2:]

		for _, r := range l1 {
			if strings.Contains(l2, string(r)) {
				score += prio[string(r)]
				break
			}
		}
	}
	fmt.Println(score)

	score = 0
	for _, gr := range grp {
		for _, r := range gr[0] {
			if strings.Contains(gr[1], string(r)) && strings.Contains(gr[2], string(r)) {
				score += prio[string(r)]
				break
			}
		}
	}
	fmt.Println(score)
}
