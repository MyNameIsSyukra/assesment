package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type Answer struct {
// 	Id   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
// 	QuestionId uuid.UUID `gorm:"type:uuid" json:"question_id"`
// 	ChoiceId   uuid.UUID `gorm:"type:uuid" json:"choice_id"`
// 	StudentId  uuid.UUID `gorm:"type:uuid" json:"student_id"`
// 	CreatedAt  time.Time `json:"created_at"`
// 	Question Question `gorm:"foreignKey:QuestionId;references:Id" json:"question"`
// 	Choice Choice `gorm:"foreignKey:ChoiceId;references:Id" json:"choice"`
// }

type Answer struct {
    ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"answer_id"`
    QuestionID uuid.UUID `gorm:"type:uuid" json:"question_id"`
    ChoiceID   uuid.UUID `gorm:"type:uuid" json:"choice_id"`
    SubmissionID uuid.UUID `gorm:"type:uuid" json:"submission_id"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
    DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
    
    // Relasi ke Question
    Question Question `gorm:"foreignKey:QuestionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"question,omitempty"`
    
    // Relasi ke Choice
    Choice Choice `gorm:"foreignKey:ChoiceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"choice,omitempty"`
    
    // Relasi ke Submission
    Submission Submission `gorm:"foreignKey:SubmissionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"submission,omitempty"`
}

// StudentID  uuid.UUID `gorm:"type:uuid" json:"student_id"`