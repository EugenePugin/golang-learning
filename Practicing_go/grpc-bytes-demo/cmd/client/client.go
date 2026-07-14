package grpc_bytes_demo

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// Импортируем наш сгенерированный пакет
	pb "grpc-bytes-demo/api/byteservice"

	_ "google.golang.org/grpc/encoding/gzip"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
		return
	}
	defer conn.Close()

	c := pb.NewByteProcessorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dataToSend := []byte("hello everybody!")

	r, err := c.ProcessBytes(ctx, &pb.ProcessRequest{Payload: dataToSend}, grpc.UseCompressor("gzip"))
	if err != nil {
		fmt.Printf("could not process bytes: %v\n", err)
		return
	}

	fmt.Printf("Original: %s\n", string(dataToSend))
	fmt.Printf("Received: %s\n", string(r.GetResult()))
	fmt.Printf("Received size: %d bytes\n", r.GetSize())
}
