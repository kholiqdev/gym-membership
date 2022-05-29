package dto

import "gym/cmd/domain/offline_class_booking/entity"

func CreateOfflineClassBookingResponse(booking *entity.OfflineClassBooking) OfflineClassBookingResponse {

	var offlineClassDetailResp []*OfflineClassBookingDetailResponse
	for _, c := range booking.OfflineClassBookingDetail {
		offlineClassDetailResp = append(offlineClassDetailResp, &OfflineClassBookingDetailResponse{
			ID:               c.ID,
			ClassName:        c.ClassName,
			ClassDescription: c.ClassDescription,
			ClassImage:       c.ClassImage,
			ClassPrice:       c.ClassPrice,
			ClassQuota:       c.ClassQuota,
			ClassDate:        c.ClassDate,
			ClassTime:        c.ClassTime,
			ClassTrainerName: c.ClassTrainerName,
			CreatedAt:        c.CreatedAt,
			UpdatedAt:        c.UpdatedAt,
		})
	}
	bookingResp := OfflineClassBookingResponse{
		ID:                        booking.ID,
		InvoiceNo:                 booking.InvoiceNo,
		PaymentMethod:             booking.PaymentMethod,
		MemberName:                booking.MemberName,
		MemberPhone:               booking.MemberPhone,
		MemberEmail:               booking.MemberEmail,
		Note:                      booking.Note.String,
		Status:                    booking.Status,
		Total:                     booking.Total,
		OfflineClassBookingDetail: offlineClassDetailResp,
		CreatedAt:                 booking.CreatedAt,
		UpdatedAt:                 booking.UpdatedAt,
	}

	return bookingResp
}

func CreateOfflineListResponse(bookings *entity.OfflineClassBookingList) OfflineClassBookingListResponse {
	bookingsResp := OfflineClassBookingListResponse{}
	for _, p := range *bookings {
		booking := CreateOfflineClassBookingResponse(p)
		bookingsResp = append(bookingsResp, &booking)
	}
	return bookingsResp
}
