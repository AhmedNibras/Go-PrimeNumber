package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	Wait()

	end := time.Now()
	duration := end.Sub(start)
	fmt.Println(duration)
	fmt.Println()
}

func Wait() {
	var wg sync.WaitGroup
	routines := 5
	dataSize := 1000
	chunk := dataSize / routines
	start := 0
	end := 0
	for i := 0; i < routines; i++ {
		start = end
		end = start + chunk
		wg.Add(1)
		fmt.Printf("%v -> %v\n", start, end)
		go PrimeNumbers(start, end, &wg)
	}
	wg.Wait()
}

func PrimeNumbers(num1, num2 int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	count := 0

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
	// fmt.Println("Sum: ", a)
	fmt.Println("Count: ", count)
}
