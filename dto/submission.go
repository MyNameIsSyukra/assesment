package dto

import (
	entities "assesment/entities"
	"time"

	"github.com/google/uuid"
)

type SubmissionCreateRequest struct {
	UserID	   uuid.UUID `json:"user_id" binding:"required"`
	AssessmentID uuid.UUID `json:"assessment_id" binding:"required"`
}

type SubmissionCreateResponse struct {
	ID           uuid.UUID      `gorm:"type:uuid" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid" json:"user_id"`
	AssessmentID uuid.UUID      `gorm:"type:uuid" json:"assessment_id"`
	EndedTime	time.Time      `json:"ended_time"`
	Question []entities.Question `json:"question"`
}


type GetSubmissionStudentResponse struct {
	ID        *uuid.UUID `gorm:"type:uuid" json:"id,omitempty"`
	Username  string    `json:"username"`
	Role  Role `json:"-"`
	User_userID uuid.UUID `gorm:"type:uuid" json:"user_user_id"`
	Kelas_kelasID uuid.UUID `gorm:"type:uuid" json:"kelas_kelas_id"`
	Status entities.ExamStatus `json:"status"`
	Score  float64       `json:"score"`
	TimeRemaining *int64        `json:"time_remaining,omitempty"`
}

type Role string

const (
	RoleStudent Role = "student"
	RoleTeacher Role = "teacher"
)
