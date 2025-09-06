package booking

import (
	bookingv1 "cloud-train-booking-lab/02-connect-rpc/gen/go/booking/v1"
	"context"
	"log"

	"connectrpc.com/connect"
)

type BookingServer struct{}

func (s *BookingServer) CreateBooking(
	ctx context.Context,
	req *connect.Request[bookingv1.CreateBookingRequest],
) (*connect.Response[bookingv1.CreateBookingResponse], error) {
	pnr := "PNR-12345"
	log.Printf("Received booking request for %s (%s -> %s)", req.Msg.Passenger, req.Msg.Departure, req.Msg.Destination)

	res := connect.NewResponse(&bookingv1.CreateBookingResponse{
		Pnr:    pnr,
		Status: "CONFIRMED",
	})

	return res, nil
}
