package booking

import (
	bookingv1 "cloud-train-booking-lab/01-simple-grpc/gen/go/booking/v1"
	"context"
	"log"
)

type BookingServer struct {
	bookingv1.UnimplementedBookingServiceServer
}

func (s *BookingServer) CreateBooking(
	ctx context.Context,
	req *bookingv1.CreateBookingRequest,
) (*bookingv1.CreateBookingResponse, error) {
	pnr := "PNR-12345"
	log.Printf("Received booking request: %s from %s to %s", req.Passenger, req.Departure, req.Destination)

	return &bookingv1.CreateBookingResponse{
		Pnr:    pnr,
		Status: "CONFIRMED",
	}, nil
}
