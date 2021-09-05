package handler

import (
	context "context"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/wangjiandev/category/common"
	"github.com/wangjiandev/category/domain/model"
	"github.com/wangjiandev/category/domain/service"
	category "github.com/wangjiandev/category/proto/category"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

// CreateCategory 新增
func (c *Category) CreateCategory(ctx context.Context, categoryRequest *category.CategoryRequest, createCategoryResponse *category.CreateCategoryResponse) error {
	ca := &model.Category{}
	err := common.SwapTo(categoryRequest, ca)
	if err != nil {
		return err
	}
	categoryId, err := c.CategoryDataService.AddCategory(ca)
	if err != nil {
		return err
	}
	createCategoryResponse.CategoryId = categoryId
	createCategoryResponse.Message = "创建成功"
	return nil
}

// UpdateCategory 更新
func (c *Category) UpdateCategory(ctx context.Context, categoryRequest *category.CategoryRequest, updateCategoryResponse *category.UpdateCategoryResponse) error {
	ca := &model.Category{}
	err := common.SwapTo(categoryRequest, ca)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(ca)
	if err != nil {
		return err
	}
	updateCategoryResponse.Message = "更新成功"
	return nil
}

func (c *Category) DeleteCategory(ctx context.Context, deleteCategoryRequest *category.DeleteCategoryRequest, deleteCategoryResponse *category.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(deleteCategoryRequest.CategoryId)
	if err != nil {
		return err
	}
	deleteCategoryResponse.Message = "删除成功"
	return nil
}

func (c *Category) FindCategoryByName(ctx context.Context, findByNameRequest *category.FindByNameRequest, categoryResponse *category.CategoryResponse) error {
	ca, err := c.CategoryDataService.FindCategoryByName(findByNameRequest.CategoryName)
	if err != nil {
		return err
	}
	categoryResponse.CategoryId = ca.ID
	categoryResponse.CategoryName = ca.CategoryName
	categoryResponse.CategoryLevel = ca.CategoryLevel
	categoryResponse.CategoryImage = ca.CategoryImage
	categoryResponse.CategoryParent = ca.CategoryParent
	categoryResponse.CategoryDescription = ca.CategoryDescription
	return nil
	//return common.SwapTo(ca, categoryResponse)
}

func (c *Category) FindCategoryById(ctx context.Context, findByIdRequest *category.FindByIdRequest, categoryResponse *category.CategoryResponse) error {
	ca, err := c.CategoryDataService.FindCategoryByID(findByIdRequest.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(ca, categoryResponse)
}

func (c *Category) FindCategoryByLevel(ctx context.Context, findByLevelRequest *category.FindByLevelRequest, categoryListResponse *category.CategoryListResponse) error {
	categories, err := c.CategoryDataService.FindCategoryByLevel(findByLevelRequest.CategoryLevel)
	if err != nil {
		return err
	}
	return categoriesToResponse(categories, categoryListResponse)
}

func (c *Category) FindCategoryByParent(ctx context.Context, findByParentRequest *category.FindByParentRequest, categoryListResponse *category.CategoryListResponse) error {
	categories, err := c.CategoryDataService.FindCategoryByParent(findByParentRequest.CategoryParent)
	if err != nil {
		return err
	}
	return categoriesToResponse(categories, categoryListResponse)
}

func (c *Category) FindAllCategory(ctx context.Context, findAllRequest *category.FindAllRequest, categoryListResponse *category.CategoryListResponse) error {
	categories, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	return categoriesToResponse(categories, categoryListResponse)
}

func categoriesToResponse(categories []model.Category, categoryListResponse *category.CategoryListResponse) error {
	for _, cg := range categories {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		categoryListResponse.CategoryList = append(categoryListResponse.CategoryList, cr)
	}
	return nil
}
