package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	name      string
	dir       []*dir
	files     []*file
	totalSize int
	prevDir   *dir
}

type file struct {
	dir  string
	size int
}

var allDir []*dir

func calcDirSize(dir *dir) int {
	for _, d := range dir.dir {
		dir.totalSize += calcDirSize(d)
	}

	for _, f := range dir.files {
		dir.totalSize += f.size
	}

	fmt.Println(dir.name, dir.totalSize)
	allDir = append(allDir, dir)
	return dir.totalSize
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

	var currDir, firstDir *dir
	currDir = nil
	for _, l := range lines {
		cmds := strings.Split(l, " ")

		switch cmds[0] {
		case "$":
			switch cmds[1] {
			case "cd":
				if currDir == nil && cmds[2] == "/" {
					currDir = &dir{
						name: "/",
					}
					firstDir = currDir
				} else if cmds[2] == ".." {
					currDir = currDir.prevDir
				} else {
					for _, d := range currDir.dir {
						if d.name == cmds[2] {
							currDir = d
							break
						}
					}
				}
			}
		default:
			if nb, err := strconv.Atoi(cmds[0]); err == nil {
				currDir.files = append(currDir.files, &file{
					size: nb,
					dir:  currDir.name,
				})
			} else if cmds[0] == "dir" {
				currDir.dir = append(currDir.dir, &dir{
					prevDir: currDir,
					name:    cmds[1],
				})
			}
		}
	}
	calcDirSize(firstDir)
	score := 0
	for _, d := range allDir {
		if d.totalSize <= 100000 {
			score += d.totalSize
		}
	}
	fmt.Println(score)

	//part 2

	unusedSpace := 70000000 - firstDir.totalSize
	mostLittleDir := firstDir.totalSize
	for _, d := range allDir {
		if unusedSpace+d.totalSize >= 30000000 && mostLittleDir > d.totalSize {
			mostLittleDir = d.totalSize
		}
	}
	fmt.Println(mostLittleDir)
}
