package booking

import pb "cloud-train-booking-lab/01-simple-grpc/gen/go/booking/v1"

type BookingServer struct {
	pb.UnimplementedBookingServiceServer
}
