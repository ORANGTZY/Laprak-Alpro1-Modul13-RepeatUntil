package main

import "fmt"

func main() {
	var number int
	var continueLoop bool
	for continueLoop = true; continueLoop; {
		fmt.Scan(&number)
		continueLoop = number <= 0 

	// for {
	//    	if continueLoop == false {
	// 		break
	// }
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", number)
}