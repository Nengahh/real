package repository

import (
	"context"
	"real_nimi_project/internal/domians"

	"gorm.io/gorm"
)

type (
	QuestionRepo struct {
		db *gorm.DB
	}

	QuestionRepoInterface interface {
		CreateQuestion(
			ctx context.Context,
			question *domians.Pertanyaan,
		) (*domians.Pertanyaan, error)

		GetAllQuestions(
			ctx context.Context,
		) ([]domians.Pertanyaan, error)

		DeleteQuestionByID(
			ctx context.Context,
			id int,
		) error

		GetQuestionsByID(
			ctx context.Context,
			id int,
		) ([]domians.Pertanyaan, error)

		UpdateQuestion(
			ctx context.Context,
			question *domians.Pertanyaan,
		) (*domians.Pertanyaan, error)
	}
)

func NewQuestionRepo(db *gorm.DB) QuestionRepo {
	return QuestionRepo{db: db}
}

func (repo QuestionRepo) CreateQuestion(
	ctx context.Context,
	question *domians.Pertanyaan,
) (*domians.Pertanyaan, error) {
	err := repo.db.WithContext(ctx).
		Create(&question).
		Error
	return question, err
}

func (repo QuestionRepo) GetAllQuestions(
	ctx context.Context,
) ([]domians.Pertanyaan, error) {
	var questions []domians.Pertanyaan
	err := repo.db.WithContext(ctx).Find(&questions).
		Error
	return questions, err
}

func (repo QuestionRepo) DeleteQuestionByID(
	ctx context.Context,
	id int,
) error {
	question := &domians.Pertanyaan{}

	// Find the question by ID
	err := repo.db.WithContext(ctx).
		Where("ID = ?", id).
		First(question).
		Error

	if err != nil {
		return err
	}

	// Delete the question
	err = repo.db.WithContext(ctx).Delete(question).Error

	return err
}

func (repo QuestionRepo) GetQuestionsByID(
	ctx context.Context,
	id int,
) ([]domians.Pertanyaan, error) {
	var questions []domians.Pertanyaan
	err := repo.db.WithContext(ctx).
		Where("id_quiz = ?", id).
		Find(&questions).
		Error
	return questions, err
}

func (repo QuestionRepo) UpdateQuestion(
	ctx context.Context,
	question *domians.Pertanyaan,
) (*domians.Pertanyaan, error) {
	err := repo.db.WithContext(ctx).
		Model(&domians.Pertanyaan{}).
		Where("ID = ?", question.ID).
		Updates(domians.Pertanyaan{
			Pertanyaan:   question.Pertanyaan,
			OpsiJawaban:  question.OpsiJawaban,
			JawabanBenar: question.JawabanBenar,
			IDQuiz:       question.IDQuiz,
		}).
		Error

	if err != nil {
		return nil, err
	}

	return question, nil
}
