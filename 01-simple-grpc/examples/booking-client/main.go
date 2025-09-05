package main

import (
	"context"
	"log"
	"time"

	pb "cloud-train-booking-lab/01-simple-grpc/gen/go/booking/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "127.0.0.1:50051"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBookingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := c.CreateBooking(ctx, &pb.CreateBookingRequest{
		Passenger:   "John Doe",
		Departure:   "Vienna",
		Destination: "Berlin",
	})
	if err != nil {
		log.Fatalf("could not create booking: %v", err)
	}

	log.Printf("Booking confirmed! PNR: %s, Status: %s", resp.Pnr, resp.Status)
}
