package mapper

import (
	"text2manim-demo-server/internal/domain/ent"

	pb "github.com/KinjiKawaguchi/text2manim/api/pkg/pb/text2manim/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromProto(generation *ent.Generation, protoStatus *pb.GenerationStatus) {
	generation.RequestID = protoStatus.RequestId
	generation.Prompt = protoStatus.Prompt
	generation.Status = protoStatus.Status.String()
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
		Status:       pb.GenerationStatus_Status(pb.GenerationStatus_Status_value[generation.Status]),
		VideoUrl:     generation.VideoURL,
		ScriptUrl:    generation.ScriptURL,
		ErrorMessage: generation.ErrorMessage,
		CreatedAt:    timestamppb.New(generation.CreatedAt),
		UpdatedAt:    timestamppb.New(generation.UpdatedAt),
	}
}
