package dto

import "gym/cmd/domain/online_class/entity"

func CreateOnlineClassResponse(entity *entity.OnlineClass) *OnlineClassResponse {
	return &OnlineClassResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Price:       entity.Price,
		Quota:       entity.Quota,
		Image:       entity.Image,
		Duration:    entity.Duration,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func CreateOnlineClassListResponse(enities *entity.OnlineClassList) *OnlineClassResponseList {
	onlineCLassResp := OnlineClassResponseList{}
	for _, p := range *enities {
		onlineClass := CreateOnlineClassResponse(p)
		onlineCLassResp = append(onlineCLassResp, onlineClass)
	}
	return &onlineCLassResp
}
