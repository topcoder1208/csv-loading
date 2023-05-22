package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// Open the CSV file
	csvFile, err := os.Open("customers.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Create a scanner to read the CSV file line by line
	scanner := bufio.NewScanner(csvFile)

	// Initialize a map to hold the counts for each domain
	domainCounts := make(map[string]int)

	// Loop over each line in the CSV file
	for scanner.Scan() {
		// Split the line into columns
		columns := strings.Split(scanner.Text(), ",")

		// Parse the email address from the second column
		email := strings.TrimSpace(columns[2])

		// Extract the domain from the email address
		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			continue // skip invalid email addresses
		}
		domain := strings.ToLower(parts[1])

		// Increment the count for the domain
		domainCounts[domain]++
	}

	// Sort the domains lexicographically
	domains := make([]string, 0, len(domainCounts))
	for domain := range domainCounts {
		domains = append(domains, domain)
	}
	sort.Strings(domains)

	// Output the results
	for _, domain := range domains {
		count := domainCounts[domain]
		fmt.Printf("%s: %d\n", domain, count)
	}
}
