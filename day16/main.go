package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

type valve struct {
	name string
	flow int
	path []string
	open bool

	goTos []*goTo
}

// Represent how to go to a certain valv with cost and path to follow to go
type goTo struct {
	to   *valve
	cost int
	path []string
}

func findValve(name string, valves []*valve) *valve {
	for _, v := range valves {
		if v.name == name {
			return v
		}
	}
	return nil
}

func findPath(from, to *valve, valves []*valve) []string {
	distances := make(map[*valve]int)
	prev := make(map[*valve]*valve)
	visited := make(map[*valve]bool)

	prev[to] = to

	distances[from] = 0
	for _, v := range valves {
		if v.name != from.name {
			distances[v] = math.MaxInt32
		}
	}

	for len(visited) < len(valves) {
		minDist := math.MaxInt32
		var minV *valve
		for v, dist := range distances {
			if _, ok := visited[v]; !ok && dist < minDist {
				minDist = dist
				minV = v
			}
		}

		visited[minV] = true

		sides := minV.path
		for _, s := range sides {
			v := findValve(s, valves)
			if !visited[v] {
				dist := distances[minV] + 1
				if dist < distances[v] {
					distances[v] = dist
					prev[v] = minV
				}
			}
		}
	}

	if _, ok := distances[to]; !ok {
		fmt.Println("ERROR: no path found from", from, "to", to)
		return nil
	}

	path := []string{}

	for n := to; n != from; n = prev[n] {
		path = append(path, n.name)
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func calculateAllGoTo(valves []*valve) []*valve {
	for _, v := range valves {
		if v.name == "AA" || v.flow > 0 {
			for _, vv := range valves {
				if vv.flow > 0 {
					if v.name == vv.name {
						continue
					}
					gt := &goTo{
						to:   vv,
						cost: 0,
						path: findPath(v, vv, valves),
					}
					gt.cost = len(gt.path)
					v.goTos = append(v.goTos, gt)
				}
			}
		}
	}
	return valves
}

// Old func that didn't work idk why
func findBestMoveBROKEN(min, depth, currScore int, currMove string, currentPath []*goTo, valves []*valve) ([]*goTo, int, int) {
	allOpened := true
	for _, g := range valves {
		if g.flow > 0 && !g.open {
			allOpened = false
			break
		}
	}

	if depth == 0 || min >= 30 || allOpened {
		return currentPath, currScore, min
	}

	best := 0
	bestPath := []*goTo{}
	bM := 0
	for _, g := range findValve(currMove, valves).goTos {
		if !g.to.open && g.to.flow > 0 {
			if min+g.cost+1 <= 30 {
				min = min + g.cost + 1
				currScore += g.to.flow * (30 - min)
				g.to.open = true
				currentPath = append(currentPath, g)

				move, value, m := findBestMoveBROKEN(min, depth-1, currScore, g.to.name, currentPath, valves)

				if value > best {
					best = value
					bestPath = move
					bM = m
				}
				if depth == -1 {
					fmt.Println(best, bestPath)
				}

				currentPath = slices.Delete(currentPath, len(currentPath)-1, len(currentPath))
				g.to.open = false

				currScore -= g.to.flow * (30 - min)
				min = min - g.cost - 1
			}
		}
	}

	return bestPath, best, bM
}

// """Minimax""" like to find best moves and return best score found
func findBestMove(currTime, currReleased, releasedByTurn int, currMove *valve, valves []*valve, maxTime int) int {
	s := currReleased + (releasedByTurn * (maxTime - currTime))
	max := s

	for _, g := range currMove.goTos {
		if !g.to.open {
			cost := g.cost + 1
			if currTime+cost <= maxTime {
				// Do the move
				g.to.open = true

				// Calcs new value
				time := currTime + cost
				pressure := currReleased + cost*releasedByTurn
				flow := releasedByTurn + g.to.flow
				score := findBestMove(time, pressure, flow, g.to, valves, maxTime)

				if score > max {
					max = score
				}

				//Undo the move
				g.to.open = false
			}
		}
	}

	return max
}

type entity struct {
	currMove       *valve
	currTime       int
	currReleased   int
	releasedByTurn int
}

// """Minimax""" like to find best moves and return best score found
func findBestMoveWithDUMBO(man, dumbo *entity, valves []*valve, maxTime int) int {
	sMan := man.currReleased + (man.releasedByTurn * (maxTime - man.currTime))
	sDumbo := dumbo.currReleased + (dumbo.releasedByTurn * (maxTime - dumbo.currTime))
	max := sMan + sDumbo

	nm := &entity{
		currMove:       man.currMove,
		currTime:       man.currTime,
		currReleased:   man.currReleased,
		releasedByTurn: man.releasedByTurn,
	}
	nd := &entity{
		currMove:       dumbo.currMove,
		currTime:       dumbo.currTime,
		currReleased:   dumbo.currReleased,
		releasedByTurn: dumbo.releasedByTurn,
	}
	for _, gm := range man.currMove.goTos {
		if man.currMove.name == "AA" && dumbo.currMove.name == "AA" {
			fmt.Println("PROUT")
		}
		if !gm.to.open {
			cost := gm.cost + 1
			if man.currTime+cost <= maxTime {
				gm.to.open = true
				nm.currTime = man.currTime + cost
				nm.currReleased = man.currReleased + cost*man.releasedByTurn
				nm.releasedByTurn = man.releasedByTurn + gm.to.flow
				nm.currMove = gm.to
			}
		}
		for _, gd := range dumbo.currMove.goTos {
			if !gd.to.open {
				cost := gd.cost + 1
				if dumbo.currTime+cost <= maxTime {
					gd.to.open = true
					nd.currTime = dumbo.currTime + cost
					nd.currReleased = dumbo.currReleased + cost*dumbo.releasedByTurn
					nd.releasedByTurn = dumbo.releasedByTurn + gd.to.flow
					nd.currMove = gd.to
				}
			}
			if man.currMove.name != nm.currMove.name || dumbo.currMove.name != nd.currMove.name {
				score := findBestMoveWithDUMBO(nm, nd, valves, maxTime)
				if score > max {
					max = score
				}
				nm.currMove.open = false
				nd.currMove.open = false
			}
		}
	}

	return max
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

	valves := []*valve{}
	for _, l := range lines {
		p := strings.Split(l, ";")

		v := &valve{
			open:  false,
			path:  []string{},
			goTos: []*goTo{},
		}

		_, err := fmt.Sscanf(p[0], "Valve %s has flow rate=%d", &v.name, &v.flow)
		if err != nil {
			log.Fatal(err)
		}

		tab := strings.Split(p[1], " ")
		tab = tab[5:]
		for _, ta := range tab {
			ta = strings.ReplaceAll(ta, ",", "")
			v.path = append(v.path, ta)
		}

		valves = append(valves, v)
	}

	valves = calculateAllGoTo(valves)

	currMove := findValve("AA", valves)

	score := findBestMove(0, 0, 0, currMove, valves, 30)
	fmt.Println("Part1", score)

	score = findBestMoveWithDUMBO(&entity{
		currMove:       currMove,
		currTime:       0,
		currReleased:   0,
		releasedByTurn: 0,
	}, &entity{
		currMove:       currMove,
		currTime:       0,
		currReleased:   0,
		releasedByTurn: 0,
	}, valves, 26)
	fmt.Println("Part2", score)
}
