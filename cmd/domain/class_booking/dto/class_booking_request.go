package dto

type ClassBookingCreateRequest struct {
	ClassIds      []int  `json:"class_ids" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
	Note          string `json:"note" form:"note"`
}
