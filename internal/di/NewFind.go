package di

import (
	"github.com/sandronister/filter-file/internal/dto"
	"github.com/sandronister/filter-file/internal/usecase"
	fsbussiness "github.com/sandronister/filter-file/pkg/fs_bussiness"
)

func NewFind(info *dto.SearchResult) *usecase.FindUseCase {
	return usecase.NewFindUseCase(fsbussiness.NewFSBussiness(), info)
}
