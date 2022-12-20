package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
)

func computeScore(r *ring.Ring, at int) int64 {
	for *r.Value.(*int64) != 0 {
		r = r.Next()
	}
	for i := 0; i < at; i++ {
		r = r.Next()
	}
	return *r.Value.(*int64)
}

func mixing(nb []*int64, nb2 *ring.Ring) int64 {
	for _, n := range nb {
		for nb2.Value.(*int64) != n {
			nb2 = nb2.Next()
		}
		nb := *nb2.Value.(*int64) % (int64(nb2.Len()) - 1)
		if nb > 0 {
			for i := 0; i < int(nb); i++ {
				nb2.Value, nb2.Next().Value = nb2.Next().Value, nb2.Value
				nb2 = nb2.Next()
			}
		}
		if nb < 0 {
			for i := 0; i > int(nb); i-- {
				nb2.Value, nb2.Prev().Value = nb2.Prev().Value, nb2.Value
				nb2 = nb2.Prev()
			}
		}
	}

	return computeScore(nb2, 1000) + computeScore(nb2, 2000) + computeScore(nb2, 3000)
}

func printRing(r *ring.Ring) {
	for *r.Value.(*int64) != 0 {
		r = r.Next()
	}
	r.Do(func(a any) {
		fmt.Print(*a.(*int64), " ")
	})
	fmt.Println()
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

	nb := []*int64{}
	nb2 := ring.New(len(lines))
	for _, l := range lines {
		v, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		vptr := &v
		nb = append(nb, vptr)
		nb2.Value = vptr
		nb2 = nb2.Next()
	}

	score := mixing(nb, nb2)
	fmt.Println("part1", score)

	nb2 = ring.New(len(lines))
	for _, n := range nb {
		*n *= 811589153
		nb2.Value = n
		nb2 = nb2.Next()
	}

	for i := 0; i < 10; i++ {
		score = mixing(nb, nb2)
	}
	fmt.Println("part2", score)
}
