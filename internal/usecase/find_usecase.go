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

func (f *FindUseCase) MoveFile(fileEnt entity.FileEntity) {
	if strings.Contains(fileEnt.Content, fileEnt.Keyword) {
		if err := f.fs.CopyPDF(fileEnt.Path, f.folder+"/"+fileEnt.Name); err != nil {
			fmt.Printf("Error moving file %s: %v\n", fileEnt.Path, err)
			return
		}
		f.count++
	}
}

func (f *FindUseCase) GetFilesWithKeyword(path, keyword string) error {

	err := f.fs.CreateDirectory(f.folder)

	if err != nil {
		return err
	}

	files, err := f.fs.ListPDFInDirectory(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		values, err := f.fs.OpenPDF(path + "/" + file)
		if err != nil {
			return err
		}
		fileEnt := entity.FileEntity{
			Content: values,
			Path:    path + "/" + file,
			Keyword: keyword,
			Name:    file,
		}

		f.MoveFile(fileEnt)

	}

	if f.count == 0 {
		fmt.Println("No files found with the specified keyword.")
		return nil
	}

	fmt.Printf("Total files moved: %d\n", f.count)

	return nil
}
