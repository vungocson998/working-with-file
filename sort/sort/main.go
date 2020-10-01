package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	mySort("../../files/input.txt")

	elapsed := time.Since(start)
	fmt.Printf("Time: %s\n", elapsed)
}

func readFile(filePath string, start int, end int) (numbers []int) {
	data, _ := ioutil.ReadFile(filePath)

	numbersStr := strings.Split(string(data), "\t")

	for _, v := range numbersStr {
		vInt, _ := strconv.Atoi(v)
		numbers = append(numbers, vInt)
	}

	return numbers
}

func splitNumArr(numbers []int) (left, right []int) {
	k := len(numbers)
	left = numbers[0 : k/2]
	right = numbers[k/2+1:]

	return left, right
}

func merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	count := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			slice[count] = left[i]
			count, i = count+1, i+1
		} else {
			slice[count] = right[j]
			count, j = count+1, j+1
		}
	}
	for i < len(left) {
		slice[count] = left[i]
		count, i = count+1, i+1
	}
	for j < len(right) {
		slice[count] = right[j]
		count, j = count+1, j+1
	}

	return slice
}

func partSort(nums []int) (slice []int) {
	sort.Ints(nums)

	slice = nums

	return
}

func mySort(filePath string) (slice []int) {

	numbers := readFile("../../files/input.txt", 0, 5)

	left, right := splitNumArr(numbers)

	left = partSort(left)
	right = partSort(right)

	slice = merge(left, right)

	fmt.Printf("Sorted array: %v\n", numbers)
	return
}
