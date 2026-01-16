package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tbp3() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	a := scanner.Text()
	aStrArr := strings.Split(a, " ")
	N, _ := strconv.Atoi(aStrArr[0])
	x, _ := strconv.Atoi(aStrArr[1])
	y, _ := strconv.Atoi(aStrArr[2])
	z, _ := strconv.Atoi(aStrArr[3])

	scanner.Scan()
	numbersStr := scanner.Text()

	strArr := strings.Split(numbersStr, " ")
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		arr[i], _ = strconv.Atoi(strArr[i])
	}

	xDiff := x
	yDiff := y
	zDiff := z

	for i := 0; i < len(arr); i++ {
		fmt.Println("diff", arr[i]%x, "xDiff", xDiff)
		if x-(arr[i]%x) < xDiff {
			xDiff = x - (arr[i] % x)
		}
		fmt.Println("diff", arr[i]%y, "yDiff", yDiff)
		if y-(arr[i]%y) < yDiff {
			yDiff = y - (arr[i] % y)
		}
		fmt.Println("diff", arr[i]%z, "zDiff", zDiff)
		if z-(arr[i]%z) < zDiff {
			zDiff = z - (arr[i] % z)
		}
	}

	// Нужно переработать подсчет разниц, в решении нейронки исп. массив, т.е. в массив добавляется разница из каждого шага. Возможно стоит хранить там не все разницы,
	fmt.Println(xDiff + yDiff + zDiff)
}
