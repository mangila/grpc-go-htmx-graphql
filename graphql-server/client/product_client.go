package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"shared/logger"
	"shared/service"
	"time"
)

type productClient struct {
	query   service.ProductQueryServiceClient
	command service.ProductCommandServiceClient
}

var (
	productClientInstance *productClient
)

// InitProductClient - init the product service grpc client
func InitProductClient(productServiceAddr string) {
	connParams := grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  1 * time.Second,  // Initial delay
			MaxDelay:   10 * time.Second, // Max delay for retries
			Multiplier: 1.6,              // Multiplier for exponential backoff
		},
	}
	keepaliveParams := keepalive.ClientParameters{
		Time:                10 * time.Second, // Ping server every 10 seconds if no activity
		Timeout:             5 * time.Second,  // Wait 5 seconds for ping ack before assuming connection is dead
		PermitWithoutStream: true,             // Send keepalive pings even without active RPCs
	}
	conn, err := grpc.NewClient(productServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(connParams),
		grpc.WithKeepaliveParams(keepaliveParams),
		grpc.WithUserAgent("graphql-server"))
	if err != nil {
		logger.Logger.Fatalf("connection failed: %v", err)
	}

	productClientInstance = &productClient{
		query:   service.NewProductQueryServiceClient(conn),
		command: service.NewProductCommandServiceClient(conn),
	}
	logger.Logger.Infof("connected to PRODUCT_SERVICE - %s", productServiceAddr)
}
