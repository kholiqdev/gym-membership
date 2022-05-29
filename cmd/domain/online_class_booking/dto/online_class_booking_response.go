package dto

import "time"

type OnlineClassBookingResponse struct {
	ID                       uint                                `json:"id,omitempty"`
	InvoiceNo                string                              `json:"invoice_no,omitempty"`
	PaymentMethod            string                              `json:"payment_method,omitempty"`
	MemberName               string                              `json:"member_name,omitempty"`
	MemberPhone              string                              `json:"member_phone,omitempty"`
	MemberEmail              string                              `json:"member_email,omitempty"`
	Note                     string                              `json:"note,omitempty"`
	Status                   uint                                `json:"status,omitempty"`
	Total                    float64                             `json:"total,omitempty"`
	OnlineClassBookingDetail []*OnlineClassBookingDetailResponse `json:"offline_class_booking_detail,omitempty"`
	CreatedAt                time.Time                           `json:"created_at,omitempty"`
	UpdatedAt                time.Time                           `json:"updated_at,omitempty"`
}

type OnlineClassBookingListResponse []*OnlineClassBookingResponse

type OnlineClassBookingDetailResponse struct {
	ID               uint      `json:"id,omitempty"`
	ClassName        string    `json:"class_name,omitempty"`
	ClassDescription string    `json:"class_description,omitempty"`
	ClassQuota       uint      `json:"class_quota,omitempty"`
	ClassImage       string    `json:"class_image,omitempty"`
	ClassPrice       float64   `json:"class_price,omitempty"`
	ClassDate        time.Time `json:"class_date,omitempty"`
	ClassTime        time.Time `json:"class_time,omitempty"`
	ClassDuration    uint      `json:"class_duration,omitempty"`
	ClassTrainerName string    `json:"class_trainer_name,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}
