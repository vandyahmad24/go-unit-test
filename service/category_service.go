package service

import (
	"errors"
	"github.com/vandyahmad24/go-unit-test/entity"
	"github.com/vandyahmad24/go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService)Get(id int)(*entity.Category, error){
	category := service.Repository.FindById(id)
	if category == nil{
		return category, errors.New("Category not found")
	}else{
		return category, nil
	}
}