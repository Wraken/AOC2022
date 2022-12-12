package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
)

type coord struct {
	x int
	y int
}

type point struct {
	p coord
	d int
}

type BFSnaze struct {
	array      [][]int
	visited    [][]bool
	rows, cols int
	start, end coord
	q          *list.List
}

func MakeBFS(arr [][]int, start coord, exit coord) BFSnaze {
	rows := len(arr)
	cols := len(arr[0])
	bfs := BFSnaze{
		array:   arr,
		visited: make([][]bool, rows),
		rows:    rows,
		cols:    cols,
		start:   start,
		end:     exit,
		q:       list.New(),
	}
	for i := 0; i < rows; i++ {
		bfs.visited[i] = make([]bool, cols)
	}
	bfs.q.PushBack(point{p: start, d: 0})
	bfs.visited[start.y][start.x] = true
	return bfs
}

func (b *BFSnaze) print() {
	str := ""
	for y, l := range b.array {
		for x, i := range l {
			if b.visited[y][x] {
				str += fmt.Sprint("\033[31m", string(rune(i+'a'-1)), "\033[0m")
			} else {
				str += fmt.Sprint(string(rune(i + 'a' - 1)))
			}
		}
		str += fmt.Sprintln()
	}
	fmt.Println(str)
}

func (b *BFSnaze) pushIfGoodtchi(c coord, qp point, dist int) {
	p := b.array[c.y][c.x] - b.array[qp.p.y][qp.p.x]
	if p <= 1 {
		b.q.PushBack(point{p: c, d: dist})
		b.visited[c.y][c.x] = true
	}
}

func (b *BFSnaze) findEligibleNeighbors(qp point) {
	dist := qp.d + 1
	c := coord{x: qp.p.x, y: qp.p.y - 1}
	if c.y >= 0 && !b.visited[c.y][c.x] {
		b.pushIfGoodtchi(c, qp, dist)
	}
	c = coord{x: qp.p.x, y: qp.p.y + 1}
	if c.y < b.rows && !b.visited[c.y][c.x] {
		b.pushIfGoodtchi(c, qp, dist)
	}
	c = coord{x: qp.p.x - 1, y: qp.p.y}
	if c.x >= 0 && !b.visited[c.y][c.x] {
		b.pushIfGoodtchi(c, qp, dist)
	}
	c = coord{x: qp.p.x + 1, y: qp.p.y}
	if c.x < b.cols && !b.visited[c.y][c.x] {
		b.pushIfGoodtchi(c, qp, dist)
	}
	// b.print()
	// time.Sleep(time.Millisecond * 20)
}

func (b *BFSnaze) FindShortestPath() int {
	for b.q.Len() != 0 {
		qpo := b.q.Front()
		qp := qpo.Value.(point)
		b.q.Remove(qpo)
		if qp.p.x == b.end.x && qp.p.y == b.end.y {
			return qp.d
		}
		b.findEligibleNeighbors(qp)
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

	maze := [][]int{}

	exit := coord{}
	entry := coord{}

	allA := []coord{}

	for y, l := range lines {
		tmp := []int{}
		for x, c := range l {
			switch c {
			case 'S':
				entry = coord{x: x, y: y}
				tmp = append(tmp, 0) // 0
			case 'E':
				exit = coord{x: x, y: y}
				tmp = append(tmp, int('z'-'a'))
			default:
				if c == 'a' {
					allA = append(allA, coord{x: x, y: y})
				}
				tmp = append(tmp, int(c-'a'))
			}
		}
		maze = append(maze, tmp)
	}

	bfs := MakeBFS(maze, entry, exit)
	fmt.Println(bfs.FindShortestPath())

	res := []int{}
	for _, a := range allA {
		bfs := MakeBFS(maze, a, exit)
		r := bfs.FindShortestPath()
		res = append(res, r)
	}

	ret := res[0]
	for _, r := range res {
		if ret > r && r != -1 {
			ret = r
		}
	}
	fmt.Println(ret)
}
