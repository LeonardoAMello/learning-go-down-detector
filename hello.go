package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
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
	sites := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.wikipedia.org",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.netflix.com",
		"https://www.reddit.com",
		"https://www.yahoo.com",
		"https://www.ebay.com",
		"https://www.bing.com",
		"https://www.msn.com",
		"https://www.pinterest.com",
		"https://www.tumblr.com",
		"https://www.paypal.com",
		"https://www.espn.com",
		"https://www.bbc.com",
		"https://www.cnn.com",
	}

	for _, site := range sites {
		fmt.Println(checkSite(site))
	}
}

func checkSite(site string) string {
	response, err := http.Get(site)

	if err != nil {
		return site + " NOK"
	}

	return site + " " + strconv.Itoa(response.StatusCode)
}

func showHistory() {

}

func main() {
	fmt.Println("Down Detector")
	printMenu()
}
