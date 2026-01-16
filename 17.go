package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func t17() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	numberStr := scanner.Text()

	steArr := strings.Split(numberStr, " ")
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		arr[i], _ = strconv.Atoi(steArr[i])
	}

	result := -1

	for i := N - 1; i >= 0; i-- {
		if arr[i]%2 == 0 {
			result = arr[i]
			break
		}
	}

	fmt.Println(result)
}
