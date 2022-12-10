package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type cpu struct {
	nbCycles  int
	x         int
	strPrint  []int
	strValues []int
	screen    [6][]rune
}

func newCpu(strP []int) *cpu {
	screen := [6][]rune{
		[]rune("........................................"),
		[]rune("........................................"),
		[]rune("........................................"),
		[]rune("........................................"),
		[]rune("........................................"),
		[]rune("........................................"),
	}

	return &cpu{
		nbCycles:  0,
		x:         1,
		strPrint:  strP,
		strValues: []int{},
		screen:    screen,
	}
}

func (c *cpu) calcSum() int {
	nb := 0
	for _, c := range c.strValues {
		nb += c
	}
	return nb
}

func (c *cpu) getStr() int {
	return c.x * c.nbCycles
}

func (c *cpu) drawPixel() {
	spritPos := c.x
	vert := c.nbCycles / 40
	hor := c.nbCycles % 40

	if hor >= spritPos-1 && hor <= spritPos+1 {
		c.screen[vert][hor] = '#'
	}
}

func (c *cpu) drawScreen() {
	for _, s := range c.screen {
		fmt.Println(string(s))
	}
}

func (c *cpu) IncrCycle() {
	c.drawPixel()
	c.nbCycles++
	if slices.Contains[int](c.strPrint, c.nbCycles) {
		c.strValues = append(c.strValues, c.getStr())
	}
}

func (c *cpu) noop() {
	c.IncrCycle()
}

func (c *cpu) Add(v int) {
	for i := 0; i < 2; i++ {
		c.IncrCycle()
	}
	c.x += v
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
	cpu := newCpu([]int{20, 60, 100, 140, 180, 220})

	for _, l := range lines {
		tab := strings.Split(l, " ")
		switch tab[0] {
		case "noop":
			cpu.noop()
		case "addx":
			a, _ := strconv.Atoi(tab[1])
			cpu.Add(a)
		}
	}
	fmt.Println(cpu.calcSum())

	cpu.drawScreen()
}
