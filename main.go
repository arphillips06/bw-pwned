package main

import (
	"bw-hibp-check/bitwarden"
	"fmt"
	"log"
	"syscall"

	"golang.org/x/term"
)

func main() {
	fmt.Println("Bitwarden -> HIBP checker starting ...")
	bitwarden.GetStatus()

	fmt.Printf("Enter vault password (hidden): ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
	}
	fmt.Println()

	password := string(bytePassword)
	unlockResp, err := bitwarden.UnlockVault(password)
	if err != nil {
		log.Fatalf("Unlock failed: %v", err)
	}
	log.Printf("Unlocked message: %s", unlockResp.Data.Title)

	fmt.Print("Enter Bitwarden item ID: ")
	var itemID string
	fmt.Scanln(&itemID)

	item, err := bitwarden.GetItem(itemID)
	if err != nil {
		log.Fatalf("Get item failed: %v", err)
	}

	fmt.Printf("\nItem: %s\nUsername: %s\nPassword: %s\n",
		item.Data.Name, item.Data.Login.Username, item.Data.Login.Password)

	itemsResp, err := bitwarden.ListAllItems()
	if err != nil {
		log.Fatalf("List all items failed: %v", err)
	}
	fmt.Printf("Found %d vault items\n", len(itemsResp.Data.Data))

}
