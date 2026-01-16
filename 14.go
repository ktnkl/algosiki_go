package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func t14() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	numbersStr := scanner.Text()

	strArr := strings.Split(numbersStr, " ")
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		arr[i], _ = strconv.Atoi(strArr[i])
	}

	i := 0
	j := 1

	for j < len(arr) {
		if arr[i] > arr[j] {
			j++
		} else {
			i = j
			j++
		}
	}

	fmt.Println(arr[i])
}
