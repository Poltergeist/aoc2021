package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/poltergeist/aoc2020/fileLoader"
)

func main() {
	input := strings.Split(string(fileLoader.LoadFile("./010-input")), "\n")
	count := -1
	lastNum := math.Inf(-1)
	for _, l := range input {
		lNum, _ := strconv.ParseFloat(l, 2)
		if lastNum < lNum {
			count += 1
		}
		lastNum = lNum
	}
	fmt.Println(count)
}
