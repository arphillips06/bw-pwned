package main

import (
	"bw-hibp-check/bitwarden"
	"fmt"
)

func promptMenu() int {
	fmt.Println("What do you want to do?")
	fmt.Println("  1. Check status")
	fmt.Println("  2. Get single item")
	fmt.Println("  3. List all items")
	fmt.Print("Choose an option [1-3]: ")
	var choice int
	fmt.Scanln(&choice)
	return choice
}

func main() {
	fmt.Println("Bitwarden → HIBP checker starting ...")
	for {
		choice := promptMenu()
		switch choice {
		case 1:
			bitwarden.HandleCheckStatus()
		case 2:
			bitwarden.HandleGetSingleItem()
		case 3:
			bitwarden.HandleListAllItems()
		default:
			fmt.Println("Invalid choice. Exiting.")
			return
		}
	}
}
