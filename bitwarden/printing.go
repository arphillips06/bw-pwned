package bitwarden

import (
	"bw-hibp-check/models"
	"fmt"
	"strings"
)

func printResults(results []models.Result) {
	const (
		Red    = "\033[1;31m"
		Green  = "\033[1;32m"
		Yellow = "\033[1;33m"
		Cyan   = "\033[36m"
		Bold   = "\033[1m"
		Reset  = "\033[0m"
	)
	fmt.Println("Show passwords in output? (y/n):")
	var choice string
	fmt.Scanln(&choice)
	showPw := strings.ToLower(choice) == "y"
	breachedCount := 0
	safeCount := 0
	for _, r := range results {
		if r.Pwned {
			breachedCount++
		} else {
			safeCount++
		}
		if !r.Pwned {
			continue
		}
		pw := r.Password
		if !showPw {
			pw = "********"
		}
		fmt.Printf("%sBREACHED%s\n", Red, Reset)
		fmt.Printf("%sAccount:%s   %s\n", Bold, Reset, r.URI)
		fmt.Printf("%sUsername:%s  %s\n", Bold, Reset, r.Username)
		fmt.Printf("%sPassword:%s  %s\n", Bold, Reset, pw)
		fmt.Printf("%sBreaches:%s  %s%d%s\n", Bold, Reset, Yellow, r.PwnedCount, Reset)
		fmt.Println()
	}
	total := breachedCount + safeCount
	fmt.Println("================================")
	fmt.Printf("%sScan Summary%s\n", Bold, Reset)
	fmt.Println("================================")
	fmt.Printf("%sSAFE:%s      %s%d%s\n", Bold, Reset, Green, safeCount, Reset)
	fmt.Printf("%sBREACHED:%s  %s%d%s\n", Bold, Reset, Red, breachedCount, Reset)
	fmt.Printf("%sTOTAL:%s     %s%d%s\n", Bold, Reset, Cyan, total, Reset)
	fmt.Println("================================")
}
