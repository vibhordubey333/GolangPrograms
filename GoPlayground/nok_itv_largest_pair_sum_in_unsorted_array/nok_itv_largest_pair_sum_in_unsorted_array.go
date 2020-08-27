package main

import (
	"log"
	"math"
)

func main() {

	intArr := []int{12, 34, 10, 6, 40}
	sum := FindLargestSum(intArr)
	log.Println("Largest sum is : ", sum)
}

func FindLargestSum(arr []int) {
	first := Large
}

func LargestValue(arr []int) int {
	large := math.MaxInt64
	for i := 0; i < len(arr); i++ {
		if arr[i] < small {
			small = arrr[i]
		}
	}
	return large
}

func SmallValue(arr []int) int {

	small := math.MinInt64
	for i := 0; i < len(arr); i++ {
		if arr[i] < small {
			small = arrr[i]
		}
	}
	return small
}
