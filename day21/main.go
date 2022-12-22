package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	name     string
	op       string
	relation []string
	number   int64
	yelled   bool
}

func findEq(m *monkey, mksMap map[string]*monkey) string {
	e1, e2 := "", ""
	if len(m.relation) == 2 {
		e1 = findEq(mksMap[m.relation[0]], mksMap)
		e2 = findEq(mksMap[m.relation[1]], mksMap)
		return "(" + e1 + m.op + e2 + ")"
	}
	if m.name == "humn" {
		return "x"
	}
	str := fmt.Sprintf("%d", m.number)
	return str
}

func solveEq(m *monkey, mksMap map[string]*monkey) int64 {
	var e1, e2 int64 = 0, 0
	if len(m.relation) == 2 {
		e1 = solveEq(mksMap[m.relation[0]], mksMap)
		e2 = solveEq(mksMap[m.relation[1]], mksMap)
		switch m.op {
		case "+":
			return e1 + e2
		case "-":
			return e1 - e2
		case "*":
			return e1 * e2
		case "/":
			return e1 / e2
		}
	}
	m.yelled = true
	return m.number
}

func solveEqX(m *monkey, nbToMatch, currNb int64, mksMap map[string]*monkey) int64 {
	if m.name == "humn" {
		return currNb
	}

	if !strings.Contains(findEq(mksMap[m.relation[0]], mksMap), "x") {
		nb := solveEq(mksMap[m.relation[0]], mksMap)
		switch m.op {
		case "+":
			currNb = currNb - nb
		case "-":
			currNb = (currNb - nb) / -1
		case "*":
			currNb = currNb / nb
		case "/":
			currNb = currNb * nb
		}
		//fmt.Println(currNb)
		currNb = solveEqX(mksMap[m.relation[1]], nbToMatch, currNb, mksMap)
	} else if !strings.Contains(findEq(mksMap[m.relation[1]], mksMap), "x") {
		nb := solveEq(mksMap[m.relation[1]], mksMap)
		switch m.op {
		case "+":
			currNb = currNb - nb
		case "-":
			currNb = currNb + nb
		case "*":
			currNb = currNb / nb
		case "/":
			currNb = currNb * nb
		}
		currNb = solveEqX(mksMap[m.relation[0]], nbToMatch, currNb, mksMap)
	}

	return currNb
}

func makeMonkeyYellButIYellToo(mks []*monkey, mksMap map[string]*monkey) int64 {
	root := mksMap["root"]
	eq1 := findEq(mksMap[root.relation[0]], mksMap)
	eq2 := findEq(mksMap[root.relation[1]], mksMap)
	var nbToMatch int64
	var eqRes int64
	if !strings.Contains(eq1, "x") {
		nbToMatch = solveEq(mksMap[root.relation[0]], mksMap)
		fmt.Println("nb to match", nbToMatch)
		eqRes = solveEqX(mksMap[root.relation[1]], nbToMatch, nbToMatch, mksMap)
	} else if !strings.Contains(eq2, "x") {
		nbToMatch = solveEq(mksMap[root.relation[1]], mksMap)
		fmt.Println("nb to match", nbToMatch)
		eqRes = solveEqX(mksMap[root.relation[0]], nbToMatch, nbToMatch, mksMap)
	}

	return eqRes
}

func parseMks(lines []string, part2 bool) ([]*monkey, map[string]*monkey) {
	monkeys := make([]*monkey, len(lines))
	mksMap := make(map[string]*monkey)
	for i, l := range lines {
		m := &monkey{
			yelled: false,
		}
		tab := strings.Split(l, ":")
		m.name = tab[0]
		operation := strings.ReplaceAll(tab[1], " ", "")
		if part2 && (m.name == "root" || m.name == "humn") {
			if m.name == "humn" {
				m.number = 0
				m.yelled = true
			} else {
				ops := strings.Split(operation, "+")
				m.op = "="
				m.relation = append(m.relation, ops[0], ops[1])
			}
		} else if strings.Contains(operation, "+") {
			ops := strings.Split(operation, "+")
			m.op = "+"
			m.relation = append(m.relation, ops[0], ops[1])
		} else if strings.Contains(operation, "*") {
			ops := strings.Split(operation, "*")
			m.op = "*"
			m.relation = append(m.relation, ops[0], ops[1])
		} else if strings.Contains(operation, "-") {
			ops := strings.Split(operation, "-")
			m.op = "-"
			m.relation = append(m.relation, ops[0], ops[1])
		} else if strings.Contains(operation, "/") {
			ops := strings.Split(operation, "/")
			m.op = "/"
			m.relation = append(m.relation, ops[0], ops[1])
		} else {
			nb, err := strconv.ParseInt(operation, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			m.yelled = true
			m.number = nb
		}
		mksMap[m.name] = m
		monkeys[i] = m
	}

	return monkeys, mksMap
}

// Solving works like this :
// This is the equation with exemple value for part2
// ((4+(2*(x-3)))/4) = 150
// /4 -> x / 4 = 150 -> 150 * 4 = 600
// 4+ -> 4 + x = 600 -> 600 - 4 = 596
// 2* -> 2 * x = 596 -> 596 / 2 = 298
// -3 -> x - 3 = 298 -> 298 + 3 = 301

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

	_, mksMap := parseMks(lines, false)

	nb := solveEq(mksMap["root"], mksMap)
	fmt.Println("part1:", nb)

	monkeys, mksMap := parseMks(lines, true)
	nb = makeMonkeyYellButIYellToo(monkeys, mksMap)

	fmt.Println("part2:", nb)
}
