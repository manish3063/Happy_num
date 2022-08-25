package main

import "fmt"

func main() {
	var num int = 19

	result := addDigits(num)
	fmt.Println(result)
}
func add(num int) bool {
	res := addDigits(num)
	if res == 1 {
		return true
	} else {
		return false
	}
}

func addDigits(num int) int {

	var sum int = 0
	var res int = 0
	for num > 0 {
		res = num % 10

		sum = sum + res*res
		num = num / 10

	}
	//fmt.Println(sum)
	if sum > 9 {
		sum = addDigits(sum)

	}
	return sum

}
