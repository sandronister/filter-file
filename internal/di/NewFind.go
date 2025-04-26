package di

import (
	"github.com/sandronister/filter-file/internal/usecase"
	fsbussiness "github.com/sandronister/filter-file/pkg/fs_bussiness"
)

func NewFind(folder string) *usecase.FindUseCase {
	return usecase.NewFindUseCase(fsbussiness.NewFSBussiness(), folder)
}
