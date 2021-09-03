package repository

import "github.com/vandyahmad24/go-unit-test/entity"

type CategoryRepository interface {
	FindById(Id int) *entity.Category
}
