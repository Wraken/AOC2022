package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	scan := bufio.NewScanner(r)

	scan.Split(bufio.ScanLines)

	e := []int{}
	tmp := 0
	for scan.Scan() {
		l := scan.Text()
		fmt.Println(l)
		if l == "" {
			e = append(e, tmp)
			tmp = 0
		} else {
			v, _ := strconv.Atoi(l)
			tmp += v
		}
	}
	sort.SliceStable(e, func(i, j int) bool {
		return e[i] > e[j]
	})

	fmt.Println(e[0] + e[1] + e[2])
}
