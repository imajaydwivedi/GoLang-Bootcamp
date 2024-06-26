package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%v is fizzbuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%v is fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%v is buzz\n", i)
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
