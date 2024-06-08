package main

import (
	"fmt"
	"os"
)

func printMenu() {
	fmt.Println()
	fmt.Println("Please select an option:")

	fmt.Println("1 - Start new analisys")
	fmt.Println("2 - Show analisys history")
	fmt.Println("3 - Exit")

	var option int
	fmt.Print("Option: ")
	fmt.Scan(&option)
	fmt.Println()

	switch option {
	case 1:
		startAnalisys()
	case 2:
		showHistory()
	case 3:
		fmt.Println("Exiting program...")
		os.Exit(0)
	default:
		fmt.Println("Invalid command")
	}

	printMenu()
}

func startAnalisys() {

}

func showHistory() {

}

func main() {
	fmt.Println("Down Detector")
	printMenu()
}
