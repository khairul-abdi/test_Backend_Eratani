package main

import (
	"fmt"
	"strconv"
)

func main() {
	polycarp(1000)
}

func polycarp(num int) {

	for i := 1; i <= num; i++ {
		strNum := strconv.Itoa(i)
		numEndWith, _ := strconv.Atoi(strNum[len(strNum)-1:])

		if (numEndWith != 3 && i%3 != 0) || i%3 != 0 {
			fmt.Println(i)
		}
	}
}
