// Package main implements a server for Restaurant service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	
	pb "task/proto"

	"google.golang.org/grpc"
)

const (
	port = "50051"
)

//RestaurantServiceServer Restaurant Service Struct
type RestaurantServiceServer struct {
	db *sql.DB
}

//NewRestaurantServiceServer New Restaurant Service Instance for the Server 
func NewRestaurantServiceServer(db *sql.DB) pb.RestaurantServiceServer {
	return &RestaurantServiceServer{db : db}
}

// connect returns SQL database connection from the pool
func (s *RestaurantServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

//CreateRestaurant Create a new restaurant.
func (s *RestaurantServiceServer) CreateRestaurant(ctx context.Context, req *pb.Restaurant) (*pb.RestaurantId, error){
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "INSERT INTO Restaurant(`name`, `rating`, `cuisines`, `address`, `timing`, `cft`) VALUES(?, ?, ?, ?, ?, ?)",
		req.Name, req.Rating, req.Cuisines, req.Address, req.Timing, req.Cft)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Restaurant-> "+err.Error())
	}

	id , err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Restaurant-> "+err.Error())
	}

	return &pb.RestaurantId{
		RestId : int32(id),
	}, nil
}

// ReadRestaurant Read a given restaurant by rest_id
func (s *RestaurantServiceServer) ReadRestaurant (ctx context.Context, req *pb.RestaurantId) (*pb.Restaurant, error) {
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Restaurant by ID
	rows, err := c.QueryContext(ctx, "SELECT `rest_id`, `name`, `rating`, `cuisines`, `address`, `timing`, `cft` FROM Restaurant WHERE `rest_id`=?",
		req.RestId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Restaurant-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Restaurant-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Restaurant with ID='%d' is not found",
			req.RestId))
	}

	var td pb.Restaurant
	if err := rows.Scan(&td.RestId, &td.Name, &td.Rating, &td.Cuisines, &td.Address, &td.Timing, &td.Cft); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Restaurant row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Restaurant rows with ID='%d'",
			req.RestId))
	}

	return &td, nil
}

//UpdateRestaurant Update the details of given Restaurant
func (s *RestaurantServiceServer) UpdateRestaurant (ctx context.Context, req *pb.Restaurant) (*pb.Status, error) {
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return &pb.Status{
			Success : false,
		}, err
	}
	defer c.Close()

	// update Restaurant
	res, err := c.ExecContext(ctx, "UPDATE Restaurant SET `name`=?, `rating`=?, `cuisines`=?, `address`=?, `timing`=?, `cft`=? WHERE `rest_id`=?",
	req.Name, req.Rating, req.Cuisines, req.Address, req.Timing, req.Cft, req.RestId)
	if err != nil {
		return &pb.Status{
			Success : false,
		}, status.Error(codes.Unknown, "failed to update Restaurant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return &pb.Status{
			Success : false,
		}, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return &pb.Status{
			Success : false,
		}, status.Error(codes.NotFound, fmt.Sprintf("Restaurant with ID='%d' is not found",
			req.RestId))
	}

	return &pb.Status{
		Success : true,
	}, nil
}

//DeleteRestaurant Delete Restaurant based on rest_id
func (s *RestaurantServiceServer) DeleteRestaurant (ctx context.Context, req *pb.RestaurantId) (*pb.Status, error) {
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return &pb.Status{
			Success : false,
		}, err
	}
	defer c.Close()

	// delete Restaurant
	res, err := c.ExecContext(ctx, "DELETE FROM Restaurant WHERE `rest_id`=?", req.RestId)
	if err != nil {
		return &pb.Status{
			Success : false,
		}, status.Error(codes.Unknown, "failed to delete Restaurant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return &pb.Status{
			Success : false,
		}, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return &pb.Status{
			Success : false,
		}, status.Error(codes.NotFound, fmt.Sprintf("Restaurant with ID='%d' is not found",
			req.RestId))
	}
	return &pb.Status{
		Success : true,
	}, nil
}

//GetAllRestaurants Get the list of all Restaurants
func (s *RestaurantServiceServer) GetAllRestaurants(req *pb.Empty, stream pb.RestaurantService_GetAllRestaurantsServer) error {
	ctx := context.Background()
	
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return  err
	}
	defer c.Close()

	// get Restaurant list
	rows, err := c.QueryContext(ctx, "SELECT `rest_id`, `name`, `rating`, `cuisines`, `address`, `timing`, `cft` FROM Restaurant")
	if err != nil {
		return  status.Error(codes.Unknown, "failed to select from Restaurant-> "+err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		td := new(pb.Restaurant)
		if err := rows.Scan(&td.RestId, &td.Name, &td.Rating, &td.Cuisines, &td.Address, &td.Timing, &td.Cft); err != nil {
			return  status.Error(codes.Unknown, "failed to retrieve field values from Restaurant row-> "+err.Error())
		}

		if err != nil {
			return  status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
		}

		if err := stream.Send(td); err != nil {
			return  err
		}
	}

	if err := rows.Err(); err != nil {
		return status.Error(codes.Unknown, "failed to retrieve data from Restaurant-> "+err.Error())
	}

	return nil
}

/*
CREATE TABLE `Restaurant` ( `rest_id` INT NOT NULL AUTO_INCREMENT , `name` VARCHAR(200) NULL DEFAULT NULL , `rating` FLOAT NULL DEFAULT NULL , `cuisines` VARCHAR(500) NULL DEFAULT NULL , `address` VARCHAR(500) NULL DEFAULT NULL , `timing` VARCHAR(200) NULL DEFAULT NULL , `cft` INT NULL DEFAULT NULL , PRIMARY KEY (`rest_id`));
*/

//CmdRunServer Connect to DB, set up GRPC Server
func CmdRunServer() error {
	ctx := context.Background()

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	db, err := sql.Open("mysql", "root:Nikhil1604*@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	server := grpc.NewServer()
	pb.RegisterRestaurantServiceServer(server, NewRestaurantServiceServer(db))

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()
	listen, err := net.Listen("tcp", ":" + port)
	if err != nil {
		return err
	}
	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}

func main() {
	if err := CmdRunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
