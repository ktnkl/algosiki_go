package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Задача о слиянии двух отсортированных списков с литкода, решение на слайсах
func lc21() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numbersStr1 := scanner.Text()
	strArr1 := strings.Split(numbersStr1, " ")
	var slice1 []int

	scanner.Scan()
	numbersStr2 := scanner.Text()
	strArr2 := strings.Split(numbersStr2, " ")
	var slice2 []int

	for i := 0; i < len(strArr1); i++ {
		number, _ := strconv.Atoi(strArr1[i])
		slice1 = append(slice1, number)
	}

	for i := 0; i < len(strArr2); i++ {
		number, _ := strconv.Atoi(strArr2[i])
		slice2 = append(slice2, number)
	}

	i, j := 0, 0
	var result []int

	for i < len(slice1) && j < len(slice2) {
		if slice1[i] < slice2[j] {
			result = append(result, slice1[i])
			i++
		} else {
			result = append(result, slice2[j])
			j++
		}
	}

	result = append(result, slice1[i:]...)
	result = append(result, slice2[j:]...)

	fmt.Println(result)
}
