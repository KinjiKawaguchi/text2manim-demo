package mapper

import (
	"text2manim-demo-server/internal/domain/ent"

	entGeneration "text2manim-demo-server/internal/domain/ent/generation"

	pb "github.com/KinjiKawaguchi/text2manim/api/pkg/pb/text2manim/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProtoStatusToEntStatus(protoStatus pb.GenerationStatus_Status) entGeneration.Status {
	switch protoStatus {
	case pb.GenerationStatus_STATUS_UNSPECIFIED:
		return entGeneration.StatusUnspecified
	case pb.GenerationStatus_STATUS_PENDING:
		return entGeneration.StatusPending
	case pb.GenerationStatus_STATUS_PROCESSING:
		return entGeneration.StatusProcessing
	case pb.GenerationStatus_STATUS_COMPLETED:
		return entGeneration.StatusCompleted
	case pb.GenerationStatus_STATUS_FAILED:
		return entGeneration.StatusFailed
	default:
		return entGeneration.StatusUnspecified
	}
}

func EntStatusToProtoStatus(entStatus entGeneration.Status) pb.GenerationStatus_Status {
	switch entStatus {
	case entGeneration.StatusUnspecified:
		return pb.GenerationStatus_STATUS_UNSPECIFIED
	case entGeneration.StatusPending:
		return pb.GenerationStatus_STATUS_PENDING
	case entGeneration.StatusProcessing:
		return pb.GenerationStatus_STATUS_PROCESSING
	case entGeneration.StatusCompleted:
		return pb.GenerationStatus_STATUS_COMPLETED
	case entGeneration.StatusFailed:
		return pb.GenerationStatus_STATUS_FAILED
	default:
		return pb.GenerationStatus_STATUS_UNSPECIFIED
	}
}

func FromProto(generation *ent.Generation, protoStatus *pb.GenerationStatus) {
	generation.RequestID = protoStatus.RequestId
	generation.Prompt = protoStatus.Prompt
	generation.Status = ProtoStatusToEntStatus(protoStatus.Status)
	generation.VideoURL = protoStatus.VideoUrl
	generation.ScriptURL = protoStatus.ScriptUrl
	generation.ErrorMessage = protoStatus.ErrorMessage
	generation.CreatedAt = protoStatus.CreatedAt.AsTime()
	generation.UpdatedAt = protoStatus.UpdatedAt.AsTime()
}

func ToProto(generation *ent.Generation) *pb.GenerationStatus {
	return &pb.GenerationStatus{
		RequestId:    generation.RequestID,
		Prompt:       generation.Prompt,
		Status:       EntStatusToProtoStatus(generation.Status),
		VideoUrl:     generation.VideoURL,
		ScriptUrl:    generation.ScriptURL,
		ErrorMessage: generation.ErrorMessage,
		CreatedAt:    timestamppb.New(generation.CreatedAt),
		UpdatedAt:    timestamppb.New(generation.UpdatedAt),
	}
}
