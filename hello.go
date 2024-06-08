package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	sites := loadSites()

	for _, site := range sites {
		fmt.Println(checkSite(site))
	}
}

func loadSites() []string {
	sites := []string{}

	file, _ := os.Open("sites.txt")

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		sites = append(sites, strings.TrimSpace(line))

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
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
