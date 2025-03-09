package main

import (
	"fmt"
)

// Scan scands a new folder for Git repositories
func scan(folder string) {
	fmt.Print("Found folders:\n\n")
	repositories := recursiveScanFolder(folder)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}