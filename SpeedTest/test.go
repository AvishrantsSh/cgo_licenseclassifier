package main

import "C"
import "math"

//export isPrime
func isPrime(num int64) bool {
	if num < 2 {
		return false
	}
	var i int64 = 2
	for ; i < int64(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

//export totalPrime
func totalPrime(num int64) int {
	count := 0
	var i int64 = 2
	for ; i <= num; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func main() {}
