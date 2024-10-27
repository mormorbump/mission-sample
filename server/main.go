package main

import (
	pb "com.graffity/mission-sample/pkg/grpc"
	"com.graffity/mission-sample/server/registry"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	log.Println("Welcome to Centray API! for PoC")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	s := grpc.NewServer()
	// DI用
	userRegistry := registry.NewUserRegistryImpl()
	// grpc service設定
	pb.RegisterUsersServiceServer(s, userRegistry.UserHandler())
	// heath checkの定義
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(s, healthServer)
	healthServer.SetServingStatus("health", healthpb.HealthCheckResponse_SERVING)
	reflection.Register(s)

	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
		log.Printf("Defaulting to port %s\n", port)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Printf("Listening at %v\n", listen.Addr())
	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
	return nil
}
