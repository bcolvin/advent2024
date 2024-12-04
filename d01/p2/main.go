package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("d01/p2/input.txt")
	if err != nil {
		panic(fmt.Errorf("unable to read input file: %v", err))
	}
	rows := strings.Split(string(f), "\n")
	if len(rows) == 0 {
		fmt.Println("0")
	}
	left, right, err := convertArrays(rows)
	if err != nil {
		panic(fmt.Errorf("unable to convert arrays: %v", err))
	}
	sum := 0
	for i := 0; i < len(left); i++ {
		sum += left[i] * right[left[i]]
	}
	fmt.Println(sum)
}

func convertArrays(rows []string) ([]int, map[int]int, error) {
	left := make([]int, len(rows))
	right := make(map[int]int, len(rows))
	re := regexp.MustCompile("\\s+")
	for i, r := range rows {
		if len(strings.TrimSpace(r)) == 0 {
			break
		}
		val := strings.Split(re.ReplaceAllString(r, " "), " ")
		if len(val) != 2 {
			return nil, nil, fmt.Errorf("unable to parse row %d: '%s'", i, r)
		}
		num, err := strconv.Atoi(val[0])
		if err != nil {
			return nil, nil, fmt.Errorf("Invalid value at %d: %s\n", r, val[0])
		}
		left[i] = num
		num, err = strconv.Atoi(val[1])
		if err != nil {
			return nil, nil, fmt.Errorf("Invalid value at %d: %s\n", r, val[1])
		}
		right[num]++
	}
	sort.Ints(left)
	return left, right, nil
}
