package main

import (
	"log"
	"net"

	bookingv1 "cloud-train-booking-lab/01-simple-grpc/gen/go/booking/v1"
	"cloud-train-booking-lab/01-simple-grpc/internal/booking"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	bookingv1.RegisterBookingServiceServer(s, &booking.BookingServer{})

	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
