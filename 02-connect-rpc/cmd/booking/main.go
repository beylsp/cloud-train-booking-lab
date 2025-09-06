package main

import (
	"log"
	"net/http"

	"cloud-train-booking-lab/02-connect-rpc/gen/go/booking/v1/bookingv1connect"
	"cloud-train-booking-lab/02-connect-rpc/internal/booking"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	port = ":50051"
)

func main() {
	booking := &booking.BookingServer{}
	mux := http.NewServeMux()
	path, handler := bookingv1connect.NewBookingServiceHandler(booking)
	mux.Handle(path, handler)

	log.Printf("Starting connect rpc listener on port " + port)
	if err := http.ListenAndServe(port, h2c.NewHandler(mux, &http2.Server{})); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
