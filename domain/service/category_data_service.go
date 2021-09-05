package service

import (
	"github.com/wangjiandev/category/domain/model"
	"github.com/wangjiandev/category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
	FindAllCategory() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

// NewCategoryDataService 创建
func NewCategoryDataService(categoryRepository repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{categoryRepository}
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

// AddCategory 插入
func (u *CategoryDataService) AddCategory(category *model.Category) (int64, error) {
	return u.CategoryRepository.CreateCategory(category)
}

// DeleteCategory 删除
func (u *CategoryDataService) DeleteCategory(categoryID int64) error {
	return u.CategoryRepository.DeleteCategoryByID(categoryID)
}

// UpdateCategory 更新
func (u *CategoryDataService) UpdateCategory(category *model.Category) error {
	return u.CategoryRepository.UpdateCategory(category)
}

// FindCategoryByID 根据ID查找
func (u *CategoryDataService) FindCategoryByID(categoryID int64) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByID(categoryID)
}

// FindAllCategory 查找全部
func (u *CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return u.CategoryRepository.FindAll()
}

// FindCategoryByName 根据名称查找
func (u *CategoryDataService) FindCategoryByName(name string) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByName(name)
}

// FindCategoryByLevel 根据级别查找
func (u *CategoryDataService) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByLevel(level)
}

// FindCategoryByParent 根据父ID查找
func (u *CategoryDataService) FindCategoryByParent(parent int64) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByParent(parent)
}
