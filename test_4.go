package main

import (
	"fmt"
	"strconv"
)

func main() {

	result := checkPrimeOrComposite(2)
	fmt.Println(result)
	result = checkPrimeOrComposite(3)
	fmt.Println(result)
	result = checkPrimeOrComposite(5)
	fmt.Println(result)
	result = checkPrimeOrComposite(7)
	fmt.Println(result)
	result = checkPrimeOrComposite(10)
	fmt.Println(result)
	result = checkPrimeOrComposite(9)
	fmt.Println(result)
	result = checkPrimeOrComposite(4)
	fmt.Println(result)
	result = checkPrimeOrComposite(1)
	fmt.Println(result)
}

func checkPrimeOrComposite(num int) string {
	for i := 2; i <= num; i++ {
		if num%i == 0 && num != i {
			return strconv.Itoa(num) + " bilangan komposit"
		}
	}

	if num <= 1 {
		return strconv.Itoa(num) + " bukan bilangan prima dan komposit"
	}
	return strconv.Itoa(num) + " bilangan prima"
}
