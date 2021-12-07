package main

import "fmt"

func main() {
	result := palindrome("kodok")
	fmt.Println(result)
	result = palindrome("jinggaaku")
	fmt.Println(result)
	result = palindrome("123321")
	fmt.Println(result)
}

func palindrome(str string) string {

	for i := 0; i < len(str); i++ {
		alphabetCheck := str[len(str)-1-i]
		if string(str[i]) != string(alphabetCheck) {
			return str + " itu bukan palindrome"
		}
	}
	return str + " itu palindrome"
}
