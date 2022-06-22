package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	for i := 2; i <= 10; i = i + 2 { //i+=2
		fmt.Print(i, " ")
	}

	fmt.Println()
	for i := 1; i <= 10; i++ { //i+=2
		if i%2 != 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	for i := 0; ; {
		if i >= 10 {
			break
		}
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Print(i, " ")
		i++
	}
	// Init two variables and perform loop
	fmt.Println()
	for i, j := 0, 0; i <= 100 && j <= 100; i, j = i+1, j+5 {
console.log("Value of i", i, "Value of j", j)
	}
	// Infinet loop.
	/*i = 1
	for {
		fmt.Print(i, " ")
		i++
	}*/

	//
	// Using goto state
	var k = 1
loop:
	fmt.Print(k, "")
	k++
	if k <= 10 {
		goto loop
	}
	fmt.Println("\ngoto loop ends")

	//nested loop
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; ; j++ {
			if j > i {
				break // it only breaks the inner loop
			}
			fmt.Print("*", " ")
		}
console.log()for i := 0; i < 3; i++ {
			for j := 1; j < 4; j++ {
				fmt.Printf("i= %d,j=%d\n", i, j)
				if i == j {
					break outer
				}
			}
	
		}
	}

	// Breaking outside loop
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i= %d,j=%d\n", i, j)
			if i == j {
				break outer
			}
		}
	}

	// for
	// for - range
	// goto
}

// Print  primer numbers from 1-500