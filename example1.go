package main

import "fmt"

func main() {
	//Initalize i
	i := 2

	//use the fmt package to print to stdOut
	fmt.Print("write ", i, " as ")

	//Here’s a basic switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("... ummm ... I don't know")
	}
}
