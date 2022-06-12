package dto

import "gym/cmd/domain/class_booking/entity"

func CreateClassBookingResponse(booking *entity.ClassBooking) ClassBookingResponse {

	var classDetailResp []*ClassBookingDetailResponse
	for _, c := range booking.ClassBookingDetail {
		classDetailResp = append(classDetailResp, &ClassBookingDetailResponse{
			ID:               c.ID,
			ClassName:        c.ClassName,
			ClassDescription: c.ClassDescription,
			ClassImage:       c.ClassImage,
			ClassPrice:       c.ClassPrice,
			ClassCategory:    c.ClassCategory,
			ClassDate:        c.ClassDate,
			ClassStartTime:   c.ClassStartTime,
			ClassEndTime:     c.ClassEndTime,
			ClassTrainerName: c.ClassTrainerName,
			CreatedAt:        c.CreatedAt,
			UpdatedAt:        c.UpdatedAt,
		})
	}
	bookingResp := ClassBookingResponse{
		ID:                 booking.ID,
		InvoiceNo:          booking.InvoiceNo,
		PaymentMethod:      booking.PaymentMethod,
		MemberName:         booking.MemberName,
		MemberPhone:        booking.MemberPhone,
		MemberEmail:        booking.MemberEmail,
		Note:               booking.Note.String,
		Status:             booking.Status,
		Total:              booking.Total,
		ClassBookingDetail: classDetailResp,
		CreatedAt:          booking.CreatedAt,
		UpdatedAt:          booking.UpdatedAt,
	}

	return bookingResp
}

func CreateOfflineListResponse(bookings *entity.ClassBookingList) ClassBookingListResponse {
	bookingsResp := ClassBookingListResponse{}
	for _, p := range *bookings {
		booking := CreateClassBookingResponse(p)
		bookingsResp = append(bookingsResp, &booking)
	}
	return bookingsResp
}
