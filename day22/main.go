package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type ord interface {
	int64 | string
}

type order[order ord] struct {
	order order
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

	maxLen := 0
	for _, l := range lines {
		if len(l) > maxLen && !strings.Contains(l, "R") {
			maxLen = len(l)
		}
	}

	ords := []order{}
	ret := 0
	for i, r := range lines[len(lines)-1] {
		if r >= '0' && r <= '9' {
			if ret == 0 {
				ret = i
			}
		} else {
			nb := strconv.A
		}
	}
}
