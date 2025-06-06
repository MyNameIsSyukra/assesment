package repository

import (
	"assesment/dto"
	entities "assesment/entities"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	QuestionRepository interface {
		CreateQuestion(ctx context.Context, tx *gorm.DB, question *entities.Question) (dto.QuestionResponse, error)
		GetQuestionByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Question, error)
		UpdateQuestion(ctx context.Context, tx *gorm.DB, question *entities.Question) (*entities.Question, error)
		DeleteQuestion(ctx context.Context, tx *gorm.DB, id uuid.UUID) error
		GetQuestionsByAssessmentID(ctx context.Context, tx *gorm.DB, assessmentID uuid.UUID) ([]entities.Question, error)
		CreateChoice(ctx context.Context, tx *gorm.DB, choice *entities.Choice) (*entities.Choice, error)
		// GetAllQuestions() (dto.GetAllQuestionsResponse, error)
	}
	questionRepository struct {
		Db *gorm.DB
	}
)

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{Db: db}
}

func (questionRepo *questionRepository) CreateQuestion(ctx context.Context, tx *gorm.DB, question *entities.Question) (dto.QuestionResponse, error) {	
	err := questionRepo.Db.Create(question).Error;
	if err != nil {
		return dto.QuestionResponse{}, err
	}
	
	return dto.QuestionResponse{
		ID:           question.ID,
		QuestionText: question.QuestionText,
		AssessmentID: question.AssessmentID,	
		CreatedAt: question.CreatedAt,
	}, nil
}

func (choiceRepo *questionRepository) CreateChoice(ctx context.Context, tx *gorm.DB, choice *entities.Choice) (*entities.Choice, error) {
	if err := choiceRepo.Db.Create(choice).Error; err != nil {
		return &entities.Choice{}, err
	}
	return choice, nil
}

func (questionRepo *questionRepository) GetQuestionByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*entities.Question, error) {
	var question entities.Question
	if err := questionRepo.Db.Where("id = ?", id).Preload("Choices").First(&question).Error; err != nil {
		return &entities.Question{}, err
	}
	return &question, nil
}

// func (questionRepo *questionRepository) GetAllQuestions() (dto.GetAllQuestionsResponse, error) {
// 	var questions []entities.Question
// 	if err := questionRepo.Db.Find(&questions).Error; err != nil {
// 		return dto.GetAllQuestionsResponse{}, err
// 	}
// 	return dto.GetAllQuestionsResponse{
// 		Questions: questions,
// 	}, nil
// }

func (questionRepo *questionRepository) UpdateQuestion(ctx context.Context, tx *gorm.DB, question *entities.Question) (*entities.Question, error) {
	if err := questionRepo.Db.Where("id = ?", question.ID).Updates(question).Error; err != nil {
		return &entities.Question{}, err
	}
	// fmt.Println("questionRepo", question)
	return question, nil
}

func (questionRepo *questionRepository) DeleteQuestion(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {

	if err := questionRepo.Db.Delete(&entities.Question{},"id", id).Error; err != nil {
		return err
	}
	return nil
}

func (questionRepo *questionRepository) GetQuestionsByAssessmentID(ctx context.Context, tx *gorm.DB, assessmentID uuid.UUID) ([]entities.Question, error) {
	var questions []entities.Question
	if err := questionRepo.Db.Where("assessment_id = ?", assessmentID).Preload("Choices").Find(&questions).Error; err != nil {
		return []entities.Question{}, err
	}
	return questions, nil
}



