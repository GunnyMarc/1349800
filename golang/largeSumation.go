// Golang program to print the tables up to given number
// using "for" loop.

package main

import "fmt"

func main() {
	var num int = 0

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)

	for loop := 2; loop <= num; loop++ {
		for count := 1; count <= 10; count++ {
			fmt.Printf("%d ", loop*count)
		}
		fmt.Println()
	}
}

