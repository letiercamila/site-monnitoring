package menu

import "fmt"

func ShowIntrod() {
	name := "Camila"
	var version float32 = 1.0

	fmt.Println("Hello, ms.", name)
	fmt.Println("This program is running in version:", version)
}

func CreateMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Start monitoring")
	fmt.Println("2. Show Logs")
	fmt.Println("0. Exit")

	fmt.Println("\nEnter a number: ")
}
