package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/vandyahmad24/go-unit-test/entity"
)

type  CategoryRepositoryMock struct {
	Mock mock.Mock
}


func (repository *CategoryRepositoryMock) FindById(Id int) *entity.Category {
	arguments := repository.Mock.Called(Id)
	if arguments.Get(0) == nil{
		return nil
	}else{
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
