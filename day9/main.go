package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type knot struct {
	x int
	y int
}

func moveRight(k *knot) {
	k.x += 1
}
func moveLeft(k *knot) {
	k.x -= 1
}
func moveUP(k *knot) {
	k.y += 1
}
func moveDown(k *knot) {
	k.y -= 1
}

func adjustPosition(h, t *knot) bool {
	moved := true
	for moved {
		if h.x-t.x > 1 {
			if h.y-t.y >= 1 {
				moveUP(t)
			} else if h.y-t.y <= -1 {
				moveDown(t)
			}
			moveRight(t)
			continue
		}
		if h.x-t.x < -1 {
			if h.y-t.y >= 1 {
				moveUP(t)
			} else if h.y-t.y <= -1 {
				moveDown(t)
			}
			moveLeft(t)
			continue
		}
		if h.y-t.y > 1 {
			if h.x-t.x >= 1 {
				moveRight(t)
			} else if h.x-t.x <= -1 {
				moveLeft(t)
			}
			moveUP(t)
			continue
		}
		if h.y-t.y < -1 {
			if h.x-t.x >= 1 {
				moveRight(t)
			} else if h.x-t.x <= -1 {
				moveLeft(t)
			}
			moveDown(t)
			continue
		}
		moved = false
	}
	return false
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

	t := &knot{x: 0, y: 0}
	h := &knot{x: 0, y: 0}

	alreadyVisited := map[string]bool{}
	alreadyVisited["0-0"] = true

	for _, l := range lines {
		tab := strings.Split(l, " ")
		dir := tab[0]
		nb, _ := strconv.Atoi(tab[1])

		for i := 0; i < nb; i++ {
			switch dir {
			case "R":
				moveRight(h)
			case "L":
				moveLeft(h)
			case "D":
				moveDown(h)
			case "U":
				moveUP(h)
			}
			//Adjust t position
			adjustPosition(h, t)
			alreadyVisited[fmt.Sprintf("%d-%d", t.x, t.y)] = true
		}
	}
	fmt.Println("part1:", len(alreadyVisited))

	ks := make([]*knot, 10)
	for i := 0; i < 10; i++ {
		ks[i] = &knot{x: 0, y: 0}
	}
	alreadyVisited2 := map[string]bool{}
	alreadyVisited2["0-0"] = true

	for _, l := range lines {
		tab := strings.Split(l, " ")
		dir := tab[0]
		nb, _ := strconv.Atoi(tab[1])

		for i := 0; i < nb; i++ {
			switch dir {
			case "R":
				moveRight(ks[0])
			case "L":
				moveLeft(ks[0])
			case "D":
				moveDown(ks[0])
			case "U":
				moveUP(ks[0])
			}

			for id := 1; id <= 9; id++ {
				adjustPosition(ks[id-1], ks[id])
			}
			alreadyVisited2[fmt.Sprintf("%d-%d", ks[9].x, ks[9].y)] = true
		}
	}
	fmt.Println("part2", len(alreadyVisited2))
}
