package bitwarden

import (
	"fmt"
	"log"
)

func HandleCheckStatus() {
	status, err := GetStatus()
	if err != nil {
		log.Fatalf("Check status failed: %v", err)
	}
	fmt.Printf("Status: %s\n", status.Data.Template.Status)
}
