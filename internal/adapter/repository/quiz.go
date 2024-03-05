package repository

import (
	"context"
	"real_nimi_project/internal/domians"

	"gorm.io/gorm"
)

type (
	QuizRepo struct {
		db *gorm.DB
	}

	QuizRepoInterface interface {
		CreateQuiz(
			ctx context.Context,
			quiz *domians.Quiz,
		) (*domians.Quiz, error)

		GetAllQuizzes(
			ctx context.Context,
		) ([]domians.Quiz, error)

		DeleteQuizByID(
			ctx context.Context,
			id int,
		) error

		UpdateQuiz(
			ctx context.Context,
			quiz *domians.Quiz,
		) (*domians.Quiz, error)
	}
)

func NewQuizRepo(db *gorm.DB) QuizRepo {
	return QuizRepo{db: db}
}

func (repo QuizRepo) CreateQuiz(
	ctx context.Context,
	quiz *domians.Quiz,
) (*domians.Quiz, error) {
	err := repo.db.WithContext(ctx).
		Create(&quiz).
		Error
	return quiz, err
}

func (repo QuizRepo) GetAllQuizzes(
	ctx context.Context,
) ([]domians.Quiz, error) {
	var quizzes []domians.Quiz
	err := repo.db.WithContext(ctx).Find(&quizzes).
		Error
	return quizzes, err
}

func (repo QuizRepo) DeleteQuizByID(
	ctx context.Context,
	id int,
) error {
	quiz := &domians.Quiz{}

	// Find the quiz by ID
	err := repo.db.WithContext(ctx).
		Where("ID = ?", id).
		First(quiz).
		Error

	if err != nil {
		return err
	}

	// Delete the quiz
	err = repo.db.WithContext(ctx).Delete(quiz).Error

	return err
}

func (repo QuizRepo) UpdateQuiz(
	ctx context.Context,
	quiz *domians.Quiz,
) (*domians.Quiz, error) {
	err := repo.db.WithContext(ctx).
		Model(&domians.Quiz{}).
		Where("ID = ?", quiz.ID).
		Updates(domians.Quiz{
			Judul:        quiz.Judul,
			Deskripsi:    quiz.Deskripsi,
			WaktuMulai:   quiz.WaktuMulai,
			WaktuSelesai: quiz.WaktuSelesai,
		}).
		Error

	if err != nil {
		return nil, err
	}

	return quiz, nil
}
