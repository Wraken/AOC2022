package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func removeFirst(s []int64) []int64 {
	ret := make([]int64, 0)
	ret = append(ret, s[:0]...)
	return append(ret, s[0+1:]...)
}

func mulOp(x, y int64) int64 {
	return x * y
}

func addOp(x, y int64) int64 {
	return x + y
}

type monkey struct {
	id    int64
	items []int64

	op    func(int64, int64) int64
	opVal int64

	testVal    int64
	monkeyGoTo [2]int64

	inspected int64
}

func (m *monkey) inspectFirst() (int64, int64) {
	m.inspected++

	nb := m.items[0]
	m.items = removeFirst(m.items)

	var new_w int64
	if m.opVal == -1 {
		new_w = m.op(nb, nb)
	} else {
		new_w = m.op(m.opVal, nb)
	}
	new_w = int64(math.Round(float64(new_w / 3)))

	if new_w%m.testVal == 0 {
		return new_w, m.monkeyGoTo[0]
	}
	return new_w, m.monkeyGoTo[1]
}

func playTurn(mks []*monkey) {
	for _, m := range mks {
		for len(m.items) > 0 {
			i, to := m.inspectFirst()
			mks[to].items = append(mks[to].items, i)
		}
	}
}

func (m *monkey) inspectFirstPart2(mod int64) (int64, int64) {
	m.inspected++

	nb := m.items[0]
	m.items = removeFirst(m.items)

	var new_w int64
	if m.opVal == -1 {
		new_w = m.op(nb, nb) % mod
	} else {
		new_w = m.op(m.opVal, nb) % mod
	}

	if new_w%m.testVal == 0 {
		return new_w, m.monkeyGoTo[0]
	}
	return new_w, m.monkeyGoTo[1]
}

func playTurnPart2(mks []*monkey, mod int64) {
	for _, m := range mks {
		for len(m.items) > 0 {
			i, to := m.inspectFirstPart2(mod)
			mks[to].items = append(mks[to].items, i)
		}
	}
}

func parseInput(lines []string, log bool) []*monkey {
	id := 0
	mks := []*monkey{}
	for i := 0; i < len(lines); i += 7 {
		nm := &monkey{
			id: int64(id),
		}

		//time parse
		items := []int64{}
		tab := strings.Split(strings.TrimSpace(strings.ReplaceAll(strings.Split(lines[i+1], ":")[1], ",", "")), " ")
		for _, i := range tab {
			nbi, _ := strconv.ParseInt(i, 10, 64)
			items = append(items, nbi)
		}
		nm.items = items

		// op parse
		op := ""
		opVal := ""
		fmt.Sscanf(strings.TrimSpace(lines[i+2]), "Operation: new = old %s %s", &op, &opVal)
		if op == "+" {
			nm.op = addOp
		} else if op == "*" {
			nm.op = mulOp
		}
		if opVal == "old" {
			nm.opVal = -1
		} else {
			v, _ := strconv.ParseInt(opVal, 10, 64)
			nm.opVal = v
		}

		// test parse
		tVal := 0
		fmt.Sscanf(strings.TrimSpace(lines[i+3]), "Test: divisible by %d", &tVal)
		nm.testVal = int64(tVal)

		var toM int64
		fmt.Sscanf(strings.TrimSpace(lines[i+4]), "If true: throw to monkey %d", &toM)
		nm.monkeyGoTo[0] = toM
		fmt.Sscanf(strings.TrimSpace(lines[i+5]), "If false: throw to monkey %d", &toM)
		nm.monkeyGoTo[1] = toM

		mks = append(mks, nm)
		id++
	}
	return mks
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

	mks := parseInput(lines, false)
	mks2 := parseInput(lines, true)

	// Part1
	for i := 0; i < 20; i++ {
		playTurn(mks)
	}
	sort.SliceStable(mks, func(i, j int) bool {
		return mks[i].inspected > mks[j].inspected
	})

	fmt.Println(mks[0].inspected * mks[1].inspected)

	// Part2
	var mod int64
	mod = 1
	// This is possible because all test values are PRIME
	// If not we should have used LCM
	// Explanation here https://www.reddit.com/r/adventofcode/comments/zizi43/2022_day_11_part_2_learning_that_it_was_that/iztt8mx?utm_medium=android_app&utm_source=share&context=3
	for _, m := range mks2 {
		mod *= m.testVal
	}
	for i := 1; i <= 10000; i++ {
		playTurnPart2(mks2, mod)
	}
	sort.SliceStable(mks2, func(i, j int) bool {
		return mks2[i].inspected > mks2[j].inspected
	})

	fmt.Println(mks2[0].inspected * mks2[1].inspected)
}
