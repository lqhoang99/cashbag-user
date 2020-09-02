package grpcserver

import (
	"context"
	"errors"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"

	"demo-user/config"
	"demo-user/dao"
	userpb "demo-user/proto"
)

type server struct{}

// GetUserBriefByID
func (s *server) GetUserBriefByID(ctx context.Context, req *userpb.GetUserBriefByIDRequest) (*userpb.GetUserBriefByIDResponse, error) {
	// Convert to ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		err := errors.New("Khong the convert to ObjectID")
		return nil, err
	}

	// Find User
	user, err := dao.UserFindByID(objectID)
	if err != nil {
		err := errors.New("Khong tim thay User")
		return nil, err
	}

	response := &userpb.GetUserBriefByIDResponse{
		UserBrief: &userpb.UserBrief{
			Id:   objectID.Hex(),
			Name: user.Name,
		},
	}

	return response, nil
}

// Start ...
func Start() {
	envVars := config.GetEnv()
	lis, err := net.Listen("tcp", envVars.GRPCUri)
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})

	log.Println(" gRPC server started on port 9000")

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while %v", err)
	}
}
