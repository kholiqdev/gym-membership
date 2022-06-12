package dto

import "gym/cmd/domain/class_category/entity"

func CreateClassCategoryResponse(classCategory *entity.ClassCategory) *ClassCategoryResponse {
	return &ClassCategoryResponse{
		ID:   classCategory.ID,
		Name: classCategory.Name,
		Slug: classCategory.Slug,
	}
}

func CreateClassCategoryListResponse(classCategorys *entity.ClassCategoryList) *ClassCategoryListResponse {
	classCategoryResp := ClassCategoryListResponse{}
	for _, p := range *classCategorys {
		classCategory := CreateClassCategoryResponse(p)
		classCategoryResp = append(classCategoryResp, classCategory)
	}
	return &classCategoryResp
}
