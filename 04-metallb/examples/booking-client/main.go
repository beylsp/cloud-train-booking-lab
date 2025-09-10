package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	bookingv1 "cloud-train-booking-lab/04-metallb/gen/go/booking/v1"
	"cloud-train-booking-lab/04-metallb/gen/go/booking/v1/bookingv1connect"

	"connectrpc.com/connect"
)

const (
	port = "50051"
)

func main() {
	host := flag.String("H", "127.0.0.1", "Service host")
	flag.Parse()

	address := fmt.Sprintf("http://%s:%s", *host, port)
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
