package domain

import (
	pb "github.com/KinjiKawaguchi/text2manim/api/pkg/pb/text2manim/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Generation struct {
	gorm.Model
	RequestId    string `gorm:"column:request_id"`
	Prompt       string `gorm:"column:prompt"`
	Status       string `gorm:"column:status"`
	VideoUrl     string `gorm:"column:video_url"`
	ScriptUrl    string `gorm:"column:script_url"`
	ErrorMessage string `gorm:"column:error_message"`
	Email        string `gorm:"column:email;not null"`
}

func NewGeneration(email, prompt string) *Generation {
	return &Generation{
		Prompt: prompt,
		Status: pb.GenerationStatus_STATUS_PENDING.String(),
		Email:  email,
	}
}

func (g *Generation) ToProto() *pb.GenerationStatus {
	return &pb.GenerationStatus{
		RequestId:    g.RequestId,
		Prompt:       g.Prompt,
		Status:       pb.GenerationStatus_Status(pb.GenerationStatus_Status_value[g.Status]),
		VideoUrl:     g.VideoUrl,
		ScriptUrl:    g.ScriptUrl,
		ErrorMessage: g.ErrorMessage,
		CreatedAt:    timestamppb.New(g.CreatedAt),
		UpdatedAt:    timestamppb.New(g.UpdatedAt),
	}
}

func (g *Generation) FromProto(status *pb.GenerationStatus) {
	g.RequestId = status.RequestId
	g.Prompt = status.Prompt
	g.Status = status.Status.String()
	g.VideoUrl = status.VideoUrl
	g.ScriptUrl = status.ScriptUrl
	g.ErrorMessage = status.ErrorMessage
	// CreatedAtとUpdatedAtはgorm.Modelから自動的に管理されるため、ここでは設定しません
}
