package dto

import "time"

type ClassBookingResponse struct {
	ID                 uint                          `json:"id,omitempty"`
	InvoiceNo          string                        `json:"invoice_no,omitempty"`
	PaymentMethod      string                        `json:"payment_method,omitempty"`
	MemberName         string                        `json:"member_name,omitempty"`
	MemberPhone        string                        `json:"member_phone,omitempty"`
	MemberEmail        string                        `json:"member_email,omitempty"`
	Note               string                        `json:"note,omitempty"`
	Status             uint                          `json:"status,omitempty"`
	Total              float64                       `json:"total,omitempty"`
	ClassBookingDetail []*ClassBookingDetailResponse `json:"class_booking_detail,omitempty"`
	CreatedAt          time.Time                     `json:"created_at,omitempty"`
	UpdatedAt          time.Time                     `json:"updated_at,omitempty"`
}

type ClassBookingListResponse []*ClassBookingResponse

type ClassBookingDetailResponse struct {
	ID               uint      `json:"id,omitempty"`
	ClassName        string    `json:"class_name,omitempty"`
	ClassDescription string    `json:"class_description,omitempty"`
	ClassMeetLink    string    `json:"class_meet_link,omitempty"`
	ClassCategory    uint      `json:"class_category,omitempty"`
	ClassImage       string    `json:"class_image,omitempty"`
	ClassPrice       float64   `json:"class_price,omitempty"`
	ClassDate        time.Time `json:"class_date,omitempty"`
	ClassStartTime   time.Time `json:"class_start_time,omitempty"`
	ClassEndTime     time.Time `json:"class_end_time,omitempty"`
	ClassTrainerName string    `json:"class_trainer_name,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}
