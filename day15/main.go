package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"time"
)

type sensor struct {
	x int
	y int

	beaconX int
	beaconY int

	d int // manhattan distance truc
}

// or taxicab distance
func manhattanDistance(x, y, x1, y1 int) int {
	return int(math.Abs(float64(x1)-float64(x)) + math.Abs(float64(y1)-float64(y)))
}

func checkBeaconY(minX, maxX, y int, sensors []sensor) int {
	nb := 0
	for x := minX; x <= maxX; x++ {
		// check if sensors have reach this point, if yes, no beacon should be here
		for _, s := range sensors {
			d := manhattanDistance(s.x, s.y, x, y)
			if d <= s.d {
				// check if the position is not a beacon
				b := false
				for _, s := range sensors {
					if s.beaconX == x && s.beaconY == y {
						b = true
						break
					}
				}
				if !b {
					nb++
				}
				break
			}
		}
	}
	return nb
}

func findDistressBeacon(min, max int, sensors []sensor) int {
	// order sensors
	sort.SliceStable(sensors, func(i, j int) bool {
		return sensors[i].x < sensors[j].x
	})

	minY := min
	maxY := max

	// calculate all sensors occupation on all y
	// if occupation is less than maxX, it seems that we have the distress beacon
	for y := minY; y <= maxY; y++ {
		// calculate y distance from sensor to actual y
		currx := min
		for _, s := range sensors {
			d := manhattanDistance(s.x, s.y, s.x, y)
			//calculate wave length of the sensor at distance
			wl := (s.d-d)*2 + 1
			if wl > 0 {
				wX1 := s.x - (wl / 2)
				wX2 := s.x + (wl / 2)
				if wX1 < 0 {
					wX1 = 0
				}
				// Fill x with length of the wave; avoid overlapping
				if currx >= wX1 && currx <= wX2 {
					currx = wX2
				}
				if currx >= max {
					break
				}
			}
		}
		if currx < max {
			// +1 because the empty spot is just after
			return (currx+1)*4000000 + y
		}
	}

	return -1
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

	sensors := []sensor{}
	for _, l := range lines {
		s := sensor{}
		_, err := fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x, &s.y, &s.beaconX, &s.beaconY)
		if err != nil {
			log.Fatal(err)
		}
		s.d = int(manhattanDistance(s.x, s.y, s.beaconX, s.beaconY))
		sensors = append(sensors, s)
	}
	minX, maxX := sensors[0].x-sensors[0].d, sensors[0].x+sensors[0].d
	for _, s := range sensors {
		min, max := s.x-s.d, s.x+s.d
		if min < minX {
			minX = min
		}
		if max > maxX {
			maxX = max
		}
	}

	fmt.Println("Part 1:", checkBeaconY(minX, maxX, 2000000, sensors))
	t := time.Now()
	nb := findDistressBeacon(0, 4000000, sensors)
	elapsed := time.Since(t)
	fmt.Println(elapsed)
	fmt.Println(nb)
}
