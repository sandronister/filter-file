package main

import (
	"fmt"
	"os"

	"github.com/sandronister/filter-file/internal/di"
	inputclear "github.com/sandronister/filter-file/pkg/input_clear"
)

func main() {

	var directory string

	if len(os.Args) < 2 {
		directory = "./"
	}

	if len(os.Args) > 2 {
		directory = os.Args[1]
	}

	fmt.Printf("Listing files in directory: %s\n", directory)

	newDirectory, err := inputclear.GetInputText("Enter the directory name: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	search, err := inputclear.GetInputText("Enter the keyword to search: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	usecase := di.NewFind(newDirectory)

	fmt.Println("Searching for files with keyword:", search)

	err = usecase.GetFilesWithKeyword(directory, search)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
