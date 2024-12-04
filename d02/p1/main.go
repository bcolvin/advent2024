package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("d02/p1/input.txt")
	if err != nil {
		panic(fmt.Errorf("unable to read input file: %v", err))
	}
	rows := strings.Split(string(f), "\n")
	if len(rows) == 0 {
		fmt.Println("0")
	}
	res := make([]int, len(rows))
	for i, row := range rows {
		items := strings.Split(strings.TrimSpace(row), " ")
		if len(items) == 1 {
			res[i] = 1
		} else if len(items) > 1 {
			b, err := isSafe(items)
			if err != nil {
				fmt.Errorf("error checking safety for row[%d] %s: %v", i, row, err)
				res[i] = 0
			} else {
				res[i] = b
			}
		}
	}
	sum := 0
	for i := 0; i < len(res); i++ {
		if res[i] > 0 {
			sum++
		}
	}
	fmt.Println(sum)
}

func isSafe(vals []string) (int, error) {
	prev, err := strconv.Atoi(vals[0])
	desc := 0
	asc := 0
	if err != nil {
		return 0, fmt.Errorf("error converting %s to int: %v", vals[0], err)
	}
	for i := 1; i < len(vals); i++ {
		num, err := strconv.Atoi(vals[i])
		if err != nil {
			return 0, fmt.Errorf("error converting %s to int: %v", vals[i], err)
		}
		if num < prev && prev-num > 0 && prev-num < 4 {
			desc++
		} else if num > prev && num-prev > 0 && num-prev < 4 {
			asc++
		} else {
			return 0, nil
		}
		prev = num
	}
	if desc != 0 && asc != 0 {
		return 0, nil
	} else if desc != 0 {
		return desc, nil
	} else if asc != 0 {
		return asc, nil
	}
	return 0, nil
}
