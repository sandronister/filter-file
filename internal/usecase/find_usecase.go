package usecase

import (
	"fmt"
	"strings"

	"github.com/sandronister/filter-file/internal/entity"
	"github.com/sandronister/filter-file/pkg/fs_bussiness/types"
)

type FindUseCase struct {
	fs     types.IFSBussiness
	folder string
	count  int
}

func NewFindUseCase(fs types.IFSBussiness, folder string) *FindUseCase {
	return &FindUseCase{
		fs:     fs,
		folder: folder,
		count:  0,
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
			if err := f.fs.CopyPDF(fileEnt.Path, f.folder+"/"+fileEnt.Name); err != nil {
				fmt.Printf("Error moving file %s: %v\n", fileEnt.Path, err)
				return
			}
			f.count++
		}
	}
}

func (f *FindUseCase) GetFilesWithKeyword(path, keyword string) error {

	maxwork := 5
	content := make(chan entity.FileEntity)

	err := f.fs.CreateDirectory(f.folder)

	if err != nil {
		return err
	}

	files, err := f.fs.ListPDFInDirectory(path)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		fmt.Println("No files found in the specified directory.")
		return nil
	}

	for range maxwork {
		go f.MoveFile(content)
	}

	f.GetContent(content, path, keyword, files)

	if f.count == 0 {
		fmt.Println("No files found with the specified keyword.")
		return nil
	}

	fmt.Printf("Total files moved: %d\n", f.count)

	return nil
}

func (f *FindUseCase) GetContent(content chan<- entity.FileEntity, path, keyword string, files []string) {
	for _, file := range files {
		fileContent, err := f.fs.OpenPDF(path + "/" + file)
		if err != nil {
			return
		}

		content <- entity.FileEntity{
			Content: fileContent,
			Path:    path + "/" + file,
			Keyword: keyword,
			Name:    file,
		}
	}

}
