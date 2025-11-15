package helper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptItemID() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter item ID: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)

		if id != "" {
			return id
		}

		fmt.Println("Item ID cannot be empty.")
	}
}
