package usecase

import (
	"fmt"
	"strings"

	"github.com/sandronister/filter-file/internal/dto"
	"github.com/sandronister/filter-file/internal/entity"
	"github.com/sandronister/filter-file/pkg/fs_bussiness/types"
)

type FindUseCase struct {
	fs    types.IFSBussiness
	info  *dto.SearchResult
	count int
}

func NewFindUseCase(fs types.IFSBussiness, info *dto.SearchResult) *FindUseCase {
	return &FindUseCase{
		fs:    fs,
		info:  info,
		count: 0,
	}
}

func (f *FindUseCase) ReadFile(content chan<- string, path string) {
	fileContent, err := f.fs.OpenPDF(path)
	if err != nil {
		return
	}
	content <- fileContent
}

func (f *FindUseCase) MoveFile(entity <-chan entity.FileEntity) {

	for fileEnt := range entity {
		if strings.Contains(fileEnt.Content, fileEnt.Keyword) {
			fullPath := f.info.NewDirectory + "/" + fileEnt.Name
			if err := f.fs.CopyPDF(fileEnt.Path, fullPath); err != nil {
				fmt.Printf("Error moving file %s: %v\n", fileEnt.Path, err)
				return
			}
			f.count++
		}
	}
}

func (f *FindUseCase) GetFilesWithKeyword() error {

	maxwork := 5
	content := make(chan entity.FileEntity)

	err := f.fs.CreateDirectory(f.info.NewDirectory)

	if err != nil {
		return err
	}

	fmt.Println("Searching for files for directory:", f.info.Directory)

	files, err := f.fs.ListPDFInDirectory(f.info.Directory)
	if err != nil {
		return err
	}

	for range maxwork {
		go f.MoveFile(content)
	}

	f.GetContent(content, files)

	if f.count == 0 {
		fmt.Println("No files found with the specified keyword.")
		return nil
	}

	fmt.Printf("Total files moved: %d\n", f.count)

	return nil
}

func (f *FindUseCase) GetContent(content chan<- entity.FileEntity, files []string) {
	for _, file := range files {
		fullPath := f.info.Directory + "/" + file

		fileContent, err := f.fs.OpenPDF(fullPath)
		if err != nil {
			return
		}

		content <- entity.FileEntity{
			Content: fileContent,
			Path:    fullPath,
			Keyword: f.info.Keyword,
			Name:    file,
		}
	}

}
