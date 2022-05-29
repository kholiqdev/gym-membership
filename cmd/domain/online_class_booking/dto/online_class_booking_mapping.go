package dto

import "gym/cmd/domain/online_class_booking/entity"

func CreateOnlineClassBookingResponse(booking *entity.OnlineClassBooking) OnlineClassBookingResponse {

	var onlineClassDetailResp []*OnlineClassBookingDetailResponse
	for _, c := range booking.OnlineClassBookingDetail {
		onlineClassDetailResp = append(onlineClassDetailResp, &OnlineClassBookingDetailResponse{
			ID:               c.ID,
			ClassName:        c.ClassName,
			ClassDescription: c.ClassDescription,
			ClassImage:       c.ClassImage,
			ClassPrice:       c.ClassPrice,
			ClassQuota:       c.ClassQuota,
			CreatedAt:        c.CreatedAt,
			UpdatedAt:        c.UpdatedAt,
		})
	}
	bookingResp := OnlineClassBookingResponse{
		ID:                       booking.ID,
		InvoiceNo:                booking.InvoiceNo,
		PaymentMethod:            booking.PaymentMethod,
		MemberName:               booking.MemberName,
		MemberPhone:              booking.MemberPhone,
		MemberEmail:              booking.MemberEmail,
		Note:                     booking.Note.String,
		Status:                   booking.Status,
		Total:                    booking.Total,
		OnlineClassBookingDetail: onlineClassDetailResp,
		CreatedAt:                booking.CreatedAt,
		UpdatedAt:                booking.UpdatedAt,
	}

	return bookingResp
}

func CreateOnlineClassBookingListResponse(bookings *entity.OnlineClassBookingList) OnlineClassBookingListResponse {
	bookingsResp := OnlineClassBookingListResponse{}
	for _, p := range *bookings {
		booking := CreateOnlineClassBookingResponse(p)
		bookingsResp = append(bookingsResp, &booking)
	}
	return bookingsResp
}
