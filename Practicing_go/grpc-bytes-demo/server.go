package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// Замените путь на фактический путь к сгенерированному пакету
	pb "github.com/yourusername/project/byteservice"
)

type server struct {
	pb.UnimplementedByteProcessorServer
}

// ProcessBytes реализует метод из proto-файла
func (s *server) ProcessBytes(ctx context.Context, req *pb.ProcessRequest) (*pb.ProcessResponse, error) {
	// req.GetPayload() возвращает обычный []byte
	inputBytes := req.GetPayload()
	length := len(inputBytes)

	// Разворачиваем слайс байт
	reversed := make([]byte, length)
	for i, b := range inputBytes {
		reversed[length-1-i] = b
	}

	return &pb.ProcessResponse{
		Result: reversed,
		Size:   int32(length),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterByteProcessorServer(s, &server{})
	reflection.Register(s) // Полезно для дебага через утилиты типа grpcurl

	fmt.Println("gRPC Server running on :50051")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
