package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/site-monnitoring/logs"
	"github.com/site-monnitoring/menu"
)

const monnitor = 13

const delay = 1

func main() {
	menu.ShowIntrod()
	for {
		menu.CreateMenu()
		command := readCommand()
		chooseCommand(command)
	}
}

func readCommand() int {
	var command int

	fmt.Scan(&command)
	fmt.Println("\nYou entered:", command)

	return command
}

func chooseCommand(command int) {
	switch command {
	case 1:
		startMonitoring()
	case 2:
		logs.PrintLogs()
	case 0:
		fmt.Println("\nExiting...")
		os.Exit(0)
	default:
		fmt.Println("\nInvalid command.")
		os.Exit(-1)
	}
}

func startMonitoring() {
	fmt.Println("\nMonitoring...")

	sites := readFile()
	for i := 0; i < monnitor; i++ {
		for _, site := range sites {
			fmt.Println("\nChecking site:", site)
			testsite(site)
		}
		fmt.Println("\n-------------------------------------------------")
		time.Sleep(delay * time.Hour)
	}
}

func testsite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Site is down! Status Code:", resp.StatusCode)
		logs.RegisterLogs(site, false)
	} else {
		fmt.Println("Site is up! Status Code:", resp.StatusCode)
		logs.RegisterLogs(site, true)
	}
}

func readFile() []string {
	var sites []string

	osFile, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(osFile)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	fmt.Println("\nSites:", sites)

	osFile.Close()

	return sites
}
