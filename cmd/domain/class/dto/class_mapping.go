package dto

import "gym/cmd/domain/class/entity"

func CreateClassResponse(entity *entity.Class) *ClassResponse {
	return &ClassResponse{
		ID:            entity.ID,
		ClassCategory: entity.ClassCategoryID,
		CategoryName:  entity.CategoryName,
		Name:          entity.Name,
		TrainerName:   entity.TrainerName,
		Description:   entity.Description,
		Price:         entity.Price,
		MeetLink:      entity.MeetLink.String,
		Image:         entity.Image,
		Date:          entity.Date,
		StartTime:     entity.StartTime,
		EndTime:       entity.EndTime,
		IsOffline:     entity.IsOffline,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
	}
}

func CreateClassListResponse(enities *entity.ClassList) *ClassResponseList {
	offlineCLassResp := ClassResponseList{}
	for _, p := range *enities {
		class := CreateClassResponse(p)
		offlineCLassResp = append(offlineCLassResp, class)
	}
	return &offlineCLassResp
}
