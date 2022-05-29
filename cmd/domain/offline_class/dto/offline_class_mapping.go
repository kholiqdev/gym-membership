package dto

import "gym/cmd/domain/offline_class/entity"

func CreateOfflineClassResponse(entity *entity.OfflineClass) *OfflineClassResponse {
	return &OfflineClassResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Price:       entity.Price,
		Quota:       entity.Quota,
		Image:       entity.Image,
		Date:        entity.Date.String(),
		Time:        entity.Time.String(),
		Duration:    entity.Duration,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func CreateOfflineClassListResponse(enities *entity.OfflineClassList) *OfflineClassResponseList {
	offlineCLassResp := OfflineClassResponseList{}
	for _, p := range *enities {
		offlineClass := CreateOfflineClassResponse(p)
		offlineCLassResp = append(offlineCLassResp, offlineClass)
	}
	return &offlineCLassResp
}
