package dto

import "gym/cmd/domain/trainer/entity"

func CreateTrainerResponse(trainer *entity.Trainer) *TrainerResponse {
	return &TrainerResponse{
		ID:        trainer.ID,
		Name:      trainer.Name,
		Email:     trainer.Email,
		Avatar:    trainer.Email,
		CreatedAt: trainer.CreatedAt,
		UpdatedAt: trainer.UpdatedAt,
	}
}

func CreateTrainerListResponse(trainers *entity.TrainerList) *TrainerListResponse {
	trainerResp := TrainerListResponse{}
	for _, p := range *trainers {
		trainer := CreateTrainerResponse(p)
		trainerResp = append(trainerResp, trainer)
	}
	return &trainerResp
}
