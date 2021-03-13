// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	userString := "RADAR"
	PalindromeChecker(userString)
	userString = "MALAYALAM"
	PalindromeChecker(userString)
	userString = "XMALAYALAM"
	PalindromeChecker(userString)

}

func PalindromeChecker(userString string) {
	sliceArray := []rune(userString) // Cannot convert into []string array

	flag := 0
	sliceLength := len(userString) / 2
	for i := 0; i < sliceLength; i++ {
		if sliceArray[i] == sliceArray[(len(userString)-1)-i] {
			flag++
		}
	}

	if flag == sliceLength {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}
