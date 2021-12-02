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
	for x, l := range input {
		if x + 2 >= len(input) {
			continue
		}
		l1, _ := strconv.ParseFloat(l, 2) 
		l2, _ := strconv.ParseFloat(input[x + 1], 2) 
		l3, _ := strconv.ParseFloat(input[x + 2], 2)
	
		if lastNum < l1 + l2 + l3 {
			count += 1
		}
		lastNum = l1 + l2 + l3
	}
	fmt.Println(count)
}
