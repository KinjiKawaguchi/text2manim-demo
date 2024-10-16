package domain

import (
	pb "github.com/KinjiKawaguchi/text2manim/api/pkg/pb/text2manim/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Generation struct {
	gorm.Model
	*pb.GenerationStatus
	Email string `gorm:"not null"`
}

func NewGeneration(email, prompt string) *Generation {
	return &Generation{
		GenerationStatus: &pb.GenerationStatus{
			RequestId:    "", // これは後で設定します
			Prompt:       prompt,
			Status:       pb.GenerationStatus_STATUS_PENDING,
			VideoUrl:     "",
			ScriptUrl:    "",
			ErrorMessage: "",
			CreatedAt:    timestamppb.Now(),
			UpdatedAt:    timestamppb.Now(),
		},
		Email: email,
	}
}

func (g *Generation) ToProto() *pb.GenerationStatus {
	return g.GenerationStatus
}

func (g *Generation) FromProto(status *pb.GenerationStatus) {
	g.GenerationStatus = status
}

func (g *Generation) BeforeCreate(tx *gorm.DB) error {
	g.Model.ID = 0 // Auto-increment
	return nil
}

func (g *Generation) AfterFind(tx *gorm.DB) error {
	// GORMのタイムスタンプをProtobufのタイムスタンプに変換
	g.GenerationStatus.CreatedAt = timestamppb.New(g.Model.CreatedAt)
	g.GenerationStatus.UpdatedAt = timestamppb.New(g.Model.UpdatedAt)
	return nil
}

func (g *Generation) BeforeSave(tx *gorm.DB) error {
	g.GenerationStatus.UpdatedAt = timestamppb.Now()
	return nil
}
