package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func t132() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	numbersStr := scanner.Text()

	scanner.Scan()
	element, _ := strconv.Atoi(scanner.Text())

	strArr := strings.Split(numbersStr, " ")
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		arr[i], _ = strconv.Atoi(strArr[i])
	}

	filtredArr := make([]string, 0)

	for i := 0; i < len(arr); i++ {
		if arr[i] != element {
			filtredArr = append(filtredArr, strconv.Itoa(arr[i]))
		}
	}

	fmt.Println(strings.Join(filtredArr, " "))
}
