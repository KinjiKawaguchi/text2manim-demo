package usecase

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"text2manim-demo-server/internal/domain"
	"text2manim-demo-server/internal/repository"
	"text2manim-demo-server/pkg/ratelimiter"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/KinjiKawaguchi/text2manim/api/pkg/pb/text2manim/v1"
)

type VideoGenerationUseCase interface {
	HealthCheck(ctx context.Context) error
	CreateGeneration(ctx context.Context, email, prompt string) (*pb.CreateGenerationResponse, error)
	GetGenerationStatus(ctx context.Context, requestID string) (*pb.GetGenerationStatusResponse, error)
	StreamGenerationStatus(ctx context.Context, requestID string, stream pb.Text2ManimService_StreamGenerationStatusServer) error
	CheckDatabaseConnection(ctx context.Context) error
}

type videoGenerationUseCase struct {
	repo             repository.GenerationRepository
	rateLimiter      *ratelimiter.RateLimiter
	text2ManimClient pb.Text2ManimServiceClient
	logger           *slog.Logger
}

func NewVideoGenerationUseCase(repo repository.GenerationRepository, limit int, interval time.Duration, text2ManimClient pb.Text2ManimServiceClient, logger *slog.Logger) VideoGenerationUseCase {
	return &videoGenerationUseCase{
		repo:             repo,
		rateLimiter:      ratelimiter.NewRateLimiter(limit, interval),
		text2ManimClient: text2ManimClient,
		logger:           logger,
	}
}

func (uc *videoGenerationUseCase) HealthCheck(ctx context.Context) error {
	_, err := uc.text2ManimClient.HealthCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return status.Errorf(codes.Internal, "failed to check health: %v", err)
	}
	return nil
}

func (uc *videoGenerationUseCase) CreateGeneration(ctx context.Context, email, prompt string) (*pb.CreateGenerationResponse, error) {
	if !uc.rateLimiter.Allow() {
		uc.logger.Warn("Rate limit exceeded", "email", email)
		return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
	}

	generation := domain.NewGeneration(email, prompt)

	if err := uc.repo.Create(generation); err != nil {
		uc.logger.Error("Failed to create generation record", "error", err, "id", generation.ID)
		return nil, status.Errorf(codes.Internal, "failed to create generation record: %v", err)
	}

	resp, err := uc.text2ManimClient.CreateGeneration(ctx, &pb.CreateGenerationRequest{Prompt: prompt})
	if err != nil {
		uc.logger.Error("Failed to initiate video generation", "error", err, "id", generation.ID)
		generation.Status = pb.GenerationStatus_STATUS_FAILED.String()
		if updateErr := uc.repo.Update(generation.ID, generation); updateErr != nil {
			uc.logger.Error("Failed to update generation status", "error", updateErr, "id", generation.ID)
		}
		return nil, status.Errorf(codes.Internal, "failed to initiate video generation: %v", err)
	}

	generation.RequestId = resp.RequestId
	if updateErr := uc.repo.Update(generation.ID, generation); updateErr != nil {
		uc.logger.Error("Failed to update generation with request ID", "error", updateErr, "id", generation.ID)
	}

	return resp, nil
}

func (uc *videoGenerationUseCase) GetGenerationStatus(ctx context.Context, requestID string) (*pb.GetGenerationStatusResponse, error) {
	generation, err := uc.repo.FindByRequestID(requestID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "generation not found")
		}
		uc.logger.Error("Failed to get generation status", "error", err, "requestID", requestID)
		return nil, status.Errorf(codes.Internal, "failed to get generation status: %v", err)
	}

	if generation.Status == pb.GenerationStatus_STATUS_UNSPECIFIED.String() || generation.Status == pb.GenerationStatus_STATUS_COMPLETED.String() || generation.Status == pb.GenerationStatus_STATUS_FAILED.String() {
		return &pb.GetGenerationStatusResponse{
			GenerationStatus: generation.ToProto(),
		}, nil
	}
	resp, err := uc.text2ManimClient.GetGenerationStatus(ctx, &pb.GetGenerationStatusRequest{RequestId: requestID})
	if err != nil {
		uc.logger.Error("Failed to fetch video generation status", "error", err, "requestID", requestID)
		return nil, status.Errorf(codes.Internal, "failed to fetch video generation status: %v", err)
	}

	// Update local record with API response
	generation.FromProto(resp.GenerationStatus)

	if err := uc.repo.Update(generation.ID, generation); err != nil {
		uc.logger.Error("Failed to update generation record", "error", err, "requestID", requestID)
	}

	return &pb.GetGenerationStatusResponse{
		GenerationStatus: generation.ToProto(),
	}, nil
}

func (uc *videoGenerationUseCase) CheckDatabaseConnection(ctx context.Context) error {
	return uc.repo.Ping()
}

func (uc *videoGenerationUseCase) StreamGenerationStatus(ctx context.Context, requestID string, stream pb.Text2ManimService_StreamGenerationStatusServer) error {
	uc.logger.Warn("StreamGenerationStatus is not implemented")
	return status.Errorf(codes.Unimplemented, "method not implemented")
}
