package infrastructure

import (
	"context"

	pb "github.com/KinjiKawaguchi/text2manim/api/pkg/pb/text2manim/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func NewText2ManimClient(endpoint, apiKey, environment string) (pb.Text2ManimServiceClient, error) {
	var conn *grpc.ClientConn
	var err error
	if environment == "development" {
		conn, err = grpc.NewClient(endpoint,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(apiKeyInterceptor(apiKey)),
		)
	} else {
		creds := credentials.NewClientTLSFromCert(nil, "")
		conn, err = grpc.NewClient(
			endpoint,
			grpc.WithTransportCredentials(creds),
			grpc.WithUnaryInterceptor(apiKeyInterceptor(apiKey)),
		)
	}
	if err != nil {
		return nil, err
	}
	return pb.NewText2ManimServiceClient(conn), nil
}

func apiKeyInterceptor(apiKey string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md := metadata.Pairs("x-api-key", apiKey)
		newCtx := metadata.NewOutgoingContext(ctx, md)
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}
