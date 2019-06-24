package main

import (
	"context"
	"log"
	"time"
	"io"

	"google.golang.org/grpc"

	pb "task/proto"
)
func main() {
	// get configuration
	address := ":50051"

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewRestaurantServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call CreateRestaurant
	req1 := &pb.Restaurant{
			Name : "KFC",
			Rating : 3.5,
			Cuisines : "Chicken, Bucket",
			Address : "Udyog Vihar, Gurugram",
			Timing : "10:00 AM - 9:00 PM",
			Cft : 450,
		}
	res1, err := c.CreateRestaurant(ctx, req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	
	log.Printf("Create result: <%+v>\n\n", res1)

	rest_id := res1.RestId

	// Call ReadRestaurant
	req2 := &pb.RestaurantId{
		RestId : rest_id,
	}

	res2, err := c.ReadRestaurant(ctx, req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	
	log.Printf("Read result: <%+v>\n\n", res2)

	// call UpdateRestaurant
	req3 := &pb.Restaurant{
		RestId : rest_id,
		Name : "KFC",
		Rating : 4.5,
		Cuisines : "Chicken, Bucket",
		Address : "CyberHub, Gurugram",
		Timing : "11:00 AM - 10:00 PM",
		Cft : 450,
	}

	res3, err :=c.UpdateRestaurant(ctx, req3)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	
	log.Printf("Updated result: <%+v>\n\n", res3)

	// call DeleteRestaurant
	req4 := &pb.RestaurantId{
		RestId : rest_id,
	}

	res4, err := c.DeleteRestaurant(ctx, req4)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	
	log.Printf("Delete result: <%+v>\n\n", res4)

	// Call ReadAll
	req5 := &pb.Empty{}
	stream, err := c.GetAllRestaurants(ctx, req5)
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}

	for {
		  rest, err := stream.Recv()
		  if err == io.EOF {
		    break
		  }
		  if err != nil {
		    log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		  }
		  log.Println(rest)
		}
}
