package repository

import "learn-golang/belajar-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}