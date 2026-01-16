package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func t15() {
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
	min_first := arr[i]
	min_second := arr[i+1]
	diff := arr[i+1] - arr[i]

	for i+1 < len(arr) {

		if arr[i+1]-arr[i] < diff {
			diff = arr[i+1] - arr[i]
			min_first = arr[i]
			min_second = arr[i+1]
		}

		i++
	}

	fmt.Println(min_first, min_second)
}
