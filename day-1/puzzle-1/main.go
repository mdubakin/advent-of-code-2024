package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

const filename = "input.txt"

func main() {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var (
		numL, numR     []int
		countL, countR = make(map[int]int), make(map[int]int)
		reader         = bufio.NewReader(f)
	)
	for {
		// read line
		line, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}

		// get left and right number
		l, r := getLeftRight(line)

		// add data
		addData(l, &numL, countL)
		addData(r, &numR, countR)
	}

	// sort
	slices.Sort(numL)
	slices.Sort(numR)

	// get closures
	getNextValL := getNextVal(numL, countL)
	getNextValR := getNextVal(numR, countR)

	var res int
	for {
		nextL, nextR := getNextValL(), getNextValR()
		if nextL == -1 || nextR == -1 {
			break
		}

		max, min := getMaxMin(nextL, nextR)
		res += max - min
	}

	fmt.Println(res)
}

func getNextVal(s []int, m map[int]int) func() int {
	var index int
	val := s[index]
	return func() int {
		if m[val] <= 0 {
			index++
			if index == len(s) {
				return -1
			}
			val = s[index]
		}
		m[val]--
		return val
	}
}

func getMaxMin(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

func getLeftRight(line []byte) (int, int) {
	splittedLine := strings.Fields(string(line))

	l, err := strconv.Atoi(splittedLine[0])
	if err != nil {
		panic(err)
	}
	r, err := strconv.Atoi(splittedLine[1])
	if err != nil {
		panic(err)
	}
	return l, r
}

func addData(num int, s *[]int, m map[int]int) {
	if !slices.Contains(*s, num) {
		*s = append(*s, num)
	}
	m[num]++
}
