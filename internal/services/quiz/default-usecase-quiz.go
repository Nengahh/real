package quiz

import (
	"context"
	"real_nimi_project/internal/adapter/repository"
	"real_nimi_project/internal/domians"
)

type (
	QuizUseCase struct {
		quizRepo repository.QuizRepoInterface
	}

	QuizUseCaseInterface interface {
		CreateQuiz(
			ctx context.Context,
			payload CreateQuizParam,
		) (*domians.Quiz, error)

		GetAllQuizzes(
			ctx context.Context,
		) ([]domians.Quiz, error)

		UpdateQuiz(
			ctx context.Context,
			quiz *domians.Quiz,
		) (*domians.Quiz, error)

		DeleteQuizByID(
			ctx context.Context,
			id int,
		) error
	}
)

func (uc QuizUseCase) CreateQuiz(
	ctx context.Context,
	payload CreateQuizParam,
) (*domians.Quiz, error) {
	return uc.quizRepo.CreateQuiz(
		ctx,
		&domians.Quiz{
			Judul:        payload.Judul,
			Deskripsi:    payload.Deskripsi,
			WaktuMulai:   payload.WaktuMulai,
			WaktuSelesai: payload.WaktuSelesai,
		})
}

func (uc QuizUseCase) GetAllQuizzes(
	ctx context.Context,
) ([]domians.Quiz, error) {
	return uc.quizRepo.GetAllQuizzes(ctx)
}

func (uc QuizUseCase) UpdateQuiz(
	ctx context.Context,
	quiz *domians.Quiz,
) (*domians.Quiz, error) {
	return uc.quizRepo.UpdateQuiz(ctx, quiz)
}

func (uc QuizUseCase) DeleteQuizByID(
	ctx context.Context,
	id int,
) error {
	return uc.quizRepo.DeleteQuizByID(ctx, id)
}
