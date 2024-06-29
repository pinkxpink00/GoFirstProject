package main

import "fmt"

func main() {
	name := "german"
	fmt.Println(name, "", len(name))

	var consoleRead string
	fmt.Println("Waititng you...")
	fmt.Scan(&consoleRead)
	fmt.Println(consoleRead)
}
