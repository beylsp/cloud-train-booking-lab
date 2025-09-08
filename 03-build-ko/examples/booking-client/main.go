package main

import (
	"context"
	"log"
	"net/http"
	"time"

	bookingv1 "cloud-train-booking-lab/03-build-ko/gen/go/booking/v1"
	"cloud-train-booking-lab/03-build-ko/gen/go/booking/v1/bookingv1connect"

	"connectrpc.com/connect"
)

const (
	address = "http://127.0.0.1:50051"
)

func main() {
	c := bookingv1connect.NewBookingServiceClient(http.DefaultClient, address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := c.CreateBooking(
		ctx,
		connect.NewRequest(&bookingv1.CreateBookingRequest{
			Passenger:   "John Doe",
			Departure:   "Vienna",
			Destination: "Berlin"}),
	)
	if err != nil {
		log.Fatalf("could not create booking: %v", err)
	}

	log.Printf("Booking confirmed! PNR: %s, Status: %s", resp.Msg.Pnr, resp.Msg.Status)
}
