package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	wg.Add(5)
	go printPrimeNumbers(2, 10, &wg)
	go printPrimeNumbers(20, 30, &wg)
	go printPrimeNumbers(30, 40, &wg)
	go printPrimeNumbers(50, 60, &wg)
	go printPrimeNumbers(60, 70, &wg)

	wg.Wait()
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println(duration)
	fmt.Println()
}

func printPrimeNumbers(num1, num2 int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	count := 0
	if num1 < 2 || num2 < 2 {
		fmt.Println("Numbers must be greater than 2.")
		return
	}
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= num1/2; i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num1)
			sum += num1
			count++
		}
		num1++
	}
	fmt.Println()
	fmt.Println("Sum: ", sum)
	fmt.Println("Count: ", count)
}
