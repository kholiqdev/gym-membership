package dto

import "gym/cmd/domain/member_order/entity"

func CreateMemberOrderResponse(booking *entity.MemberOrder) MemberOrderResponse {

	var classDetailResp []*MemberOrderDetailResponse
	for _, c := range booking.MemberOrderDetail {
		classDetailResp = append(classDetailResp, &MemberOrderDetailResponse{
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
	bookingResp := MemberOrderResponse{
		ID:                booking.ID,
		InvoiceNo:         booking.InvoiceNo,
		PaymentMethod:     booking.PaymentMethod,
		MemberName:        booking.MemberName,
		MemberPhone:       booking.MemberPhone,
		MemberEmail:       booking.MemberEmail,
		Note:              booking.Note.String,
		Status:            booking.Status,
		Total:             booking.Total,
		MemberOrderDetail: classDetailResp,
		CreatedAt:         booking.CreatedAt,
		UpdatedAt:         booking.UpdatedAt,
	}

	return bookingResp
}

func CreateOfflineListResponse(bookings *entity.MemberOrderList) MemberOrderListResponse {
	bookingsResp := MemberOrderListResponse{}
	for _, p := range *bookings {
		booking := CreateMemberOrderResponse(p)
		bookingsResp = append(bookingsResp, &booking)
	}
	return bookingsResp
}
