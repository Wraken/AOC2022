package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type coord struct {
	x int
	y int
}

type cave struct {
	maps [700][700]rune
	minX int
	minY int
	maxY int
	maxX int

	// for print
	pMaxX int
	pMinX int
	pMaxY int
}

func (c *cave) printMap() {
	str := ""
	for y := 0; y < c.pMaxY+4; y++ {
		for x := c.pMinX - 4; x < c.pMaxX+4; x++ {
			str += fmt.Sprint(string(c.maps[y][x]))
		}
		str += fmt.Sprintln()
	}
	fmt.Println(str)
}

func (c *cave) makeItRain(floor bool) int {
	if floor {
		for x := 0; x < 700; x++ {
			c.maps[c.maxY+2][x] = '#'
		}
	}

	nb := 0
	abyssing := false
	for !abyssing {
		sand := coord{
			x: 500,
			y: 0,
		}
		//can't spawn more
		if c.maps[0][500] == 'o' {
			break
		}
		for {
			b := true
			if c.maps[sand.y+1][sand.x] == '.' {
				c.maps[sand.y][sand.x] = '.'
				sand.y++
				c.maps[sand.y][sand.x] = 'o'
				b = false
			} else if c.maps[sand.y+1][sand.x] != '.' {
				if c.maps[sand.y+1][sand.x-1] == '.' {
					c.maps[sand.y][sand.x] = '.'
					sand.x--
					sand.y++
					c.maps[sand.y][sand.x] = 'o'
					b = false
				} else if c.maps[sand.y+1][sand.x+1] == '.' {
					c.maps[sand.y][sand.x] = '.'
					sand.x++
					sand.y++
					c.maps[sand.y][sand.x] = 'o'
					b = false
				}

			}
			//for print
			if sand.x > c.pMaxX {
				c.pMaxX = sand.x
			}
			if sand.x < c.pMinX {
				c.pMinX = sand.x
			}
			if sand.y > c.pMaxY {
				c.pMaxY = sand.y
			}
			c.printMap()
			time.Sleep(time.Millisecond * 10)

			if b && sand.x == 500 && sand.y == 0 {
				c.maps[sand.y][sand.x] = 'o'
				nb++
				break
			}
			if sand.y > c.maxY+10 {
				abyssing = true
				break
			}
			if b {
				nb++
				break
			}
		}
	}
	c.printMap()
	return nb
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

	maxX := 0
	maxY := 0
	minX := 700
	minY := 700
	coords := [][]coord{}
	for _, l := range lines {
		tab := strings.Split(strings.ReplaceAll(l, " -> ", " "), " ")
		c := []coord{}
		for _, t := range tab {
			xy := strings.Split(t, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			if x > maxX {
				maxX = x
			}
			if x < minX {
				minX = x
			}
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}
			c = append(c, coord{
				x: x,
				y: y,
			})
		}
		coords = append(coords, c)
	}

	maps := [700][700]rune{}
	for y, l := range maps {
		for x, _ := range l {
			maps[y][x] = '.'
		}
	}
	maps[0][500] = '+'
	for _, cl := range coords {
		for i := 0; i < len(cl)-1; i++ {
			ret := cl[i]
			if ret.x != cl[i+1].x {
				if ret.x < cl[i+1].x {
					for ret.x <= cl[i+1].x {
						maps[ret.y][ret.x] = '#'
						ret.x++
					}
				} else if ret.x > cl[i+1].x {
					for ret.x >= cl[i+1].x {
						maps[ret.y][ret.x] = '#'
						ret.x--
					}
				}
			} else if ret.y != cl[i+1].y {
				if ret.y < cl[i+1].y {
					for ret.y <= cl[i+1].y {
						maps[ret.y][ret.x] = '#'
						ret.y++
					}
				} else if ret.y > cl[i+1].y {
					for ret.y >= cl[i+1].y {
						maps[ret.y][ret.x] = '#'
						ret.y--
					}
				}
			}
		}
	}

	//Part 1
	c := cave{
		maps:  maps,
		minX:  minX,
		minY:  minY,
		maxY:  maxY,
		maxX:  maxX,
		pMaxX: 500 + 4,
		pMinX: 500 - 4,
		pMaxY: 2,
	}
	fmt.Println(c.makeItRain(false))

	//paart2
	c = cave{
		maps:  maps,
		minX:  minX,
		minY:  minY,
		maxY:  maxY,
		maxX:  maxX,
		pMaxX: 500 + 4,
		pMinX: 500 - 4,
		pMaxY: 2,
	}
	fmt.Println(c.makeItRain(true))
}
