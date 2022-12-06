package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func removeFromSlices(s []string, i int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:i]...)
	return append(ret, s[i+1:]...)
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

	tab := map[int][]string{}

	i := 0
	for lines[i][1] != '1' {
		i++
	}

	for y, r := range lines[i] {
		tmp := i - 1
		if r != ' ' {
			v, _ := strconv.Atoi(string(r))
			tab[v] = []string{}
			for tmp >= 0 && len(lines[tmp]) >= y {
				if lines[tmp][y] != ' ' {
					tab[v] = append(tab[v], string(lines[tmp][y]))
				}
				tmp--
			}
		}
	}

	for !strings.Contains(lines[i], "move") {
		i++
	}

	tab2 := map[int][]string{}
	for k, v := range tab {
		tab2[k] = v
	}

	for _, l := range lines[i:] {
		// Parsing shit
		nl := strings.Replace(l, "move", "", -1)
		nl = strings.Replace(nl, "from", "", -1)
		nl = strings.Replace(nl, "to", "", -1)
		nl = strings.Join(strings.Fields(nl), " ")
		z := strings.Split(nl, " ")
		nb, _ := strconv.Atoi(z[0])
		from, _ := strconv.Atoi(z[1])
		to, _ := strconv.Atoi(z[2])

		for nb > 0 {
			tab[to] = append(tab[to], tab[from][len(tab[from])-1])
			tab[from] = removeFromSlices(tab[from], len(tab[from])-1)
			nb--
		}
	}

	keys := make([]int, 0, len(tab))
	for k := range tab {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Print(tab[k][len(tab[k])-1])
	}
	fmt.Println()

	for _, l := range lines[i:] {
		// Parsing shit
		nl := strings.Replace(l, "move", "", -1)
		nl = strings.Replace(nl, "from", "", -1)
		nl = strings.Replace(nl, "to", "", -1)
		nl = strings.Join(strings.Fields(nl), " ")
		z := strings.Split(nl, " ")
		nb, _ := strconv.Atoi(z[0])
		from, _ := strconv.Atoi(z[1])
		to, _ := strconv.Atoi(z[2])

		tmp := []string{}
		y := 0
		for nb > 0 {
			tmp = append(tmp, tab2[from][len(tab2[from])-1-y])
			nb--
			y++
		}
		for i := len(tmp) - 1; i >= 0; i-- {
			tab2[to] = append(tab2[to], tmp[i])
			tab2[from] = removeFromSlices(tab2[from], len(tab2[from])-1)
		}
	}
	keys = make([]int, 0, len(tab2))
	for k := range tab2 {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Print(tab2[k][len(tab2[k])-1])
	}
	fmt.Println()
}
