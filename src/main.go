package main

import (
	"flag"
)

// Scan given a path, crawls it and its subfolders
// Searching for Git repositories

func scan(path string) {
	print("scan")
}

// Stats generate a nice graph of your Git contributions

func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}