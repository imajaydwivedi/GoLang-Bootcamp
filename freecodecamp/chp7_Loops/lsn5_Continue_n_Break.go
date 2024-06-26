package main

import (
	"fmt"
)

//printPrimes(max):
//  for n in range(2, max+1):
//    if n is 2:
//      n is prime, print it
//    if n is even:
//      n is not prime, skip to next n
//    for i in range (3, sqrt(n), 2):
//      if i can be multiplied into n:
//        n is not prime, skip to next n
//    n is prime, print it

// https://www.boot.dev/lessons/5b5123f4-8632-4e91-9045-a7e604aba7d2
func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		isPrime := true
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		for i := 3; i*i < n+1; i += 2 {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
