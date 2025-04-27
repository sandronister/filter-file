package main

import (
	"fmt"
	"os"

	"github.com/sandronister/filter-file/internal/di"
	"github.com/sandronister/filter-file/internal/dto"
	inputclear "github.com/sandronister/filter-file/pkg/input_clear"
)

func main() {

	searchDTO := &dto.SearchResult{}

	if len(os.Args) < 2 {
		searchDTO.Directory = "."
	}

	if len(os.Args) > 1 {
		searchDTO.Directory = os.Args[1]
	}

	fmt.Printf("Listing files in directory: %s\n", searchDTO.Directory)

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

	searchDTO.NewDirectory = newDirectory
	searchDTO.Keyword = search

	usecase := di.NewFind(searchDTO)

	fmt.Println("Searching for files with keyword:", search)

	err = usecase.GetFilesWithKeyword()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
