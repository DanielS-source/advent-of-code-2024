package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day1part1(input string) {
	lines := strings.Split(input, "\n")
	var list1 = make([]int64, len(lines))
	var list2 = make([]int64, len(lines))

	for i, line := range lines {
		pair := strings.Split(line, "   ")
		if len(pair) == 2 {

			p1, err := strconv.ParseInt(pair[0], 10, 64)
			if err != nil {
				fmt.Printf("Error processing line %q: %v\n", pair[0], err)
				panic(err)
			}

			p2, err := strconv.ParseInt(pair[1], 10, 64)
			if err != nil {
				fmt.Printf("Error processing line %q: %v\n", pair[1], err)
				panic(err)
			}

			list1[i] = p1
			list2[i] = p2
		}
	}

	list1 = quicksort(list1)
	list2 = quicksort(list2)

	//fmt.Println("Sorted list1:", list1)
	//fmt.Println("Sorted list2:", list2)

	total := 0.
	for i := 0; i < len(list1); i++ {
		dist := math.Abs(float64(list1[i] - list2[i]))
		//fmt.Printf("Distance for index %d of %f\n", i, dist)
		total += dist
	}

	fmt.Printf("Result %d\n", int64(total))

}

func quicksort(list []int64) []int64 {
	if len(list) <= 1 {
		return list
	}

	pivot := list[0]
	var left, right []int64

	for _, val := range list[1:] {
		if val < pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	return append(append(quicksort(left), pivot), quicksort(right)...)

}

func day1part2(input string) {
	lines := strings.Split(input, "\n")
	var list1 = make([]int64, len(lines))
	map2 := make(map[int64]int64)

	for i, line := range lines {
		pair := strings.Split(line, "   ")
		if len(pair) == 2 {

			p1, err := strconv.ParseInt(pair[0], 10, 64)
			if err != nil {
				fmt.Printf("Error processing line %q: %v\n", pair[0], err)
				panic(err)
			}

			p2, err := strconv.ParseInt(pair[1], 10, 64)
			if err != nil {
				fmt.Printf("Error processing line %q: %v\n", pair[1], err)
				panic(err)
			}

			list1[i] = p1
			map2[p2] += 1
		}
	}

	//fmt.Println("Sorted list1:", list1)
	//fmt.Println("Sorted list2:", map2)

	var similarityScore int64 = 0

	for i := 0; i < len(list1); i++ {
		key := list1[i]
		similarityScore += key * map2[key]
	}

	fmt.Printf("Similarity score of %d\n", similarityScore)

}
