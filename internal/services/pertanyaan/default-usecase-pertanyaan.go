package pertanyaan

import (
	"context"
	"real_nimi_project/internal/adapter/repository"
	"real_nimi_project/internal/domians"
)

type (
	QuestionUseCase struct {
		questionRepo repository.QuestionRepoInterface
	}

	QuestionUseCaseInterface interface {
		CreateQuestion(
			ctx context.Context,
			payload CreateQuestionParam,
		) (*domians.Pertanyaan, error)

		GetAllQuestions(
			ctx context.Context,
		) ([]domians.Pertanyaan, error)

		GetQuestionsByID(
			ctx context.Context,
			id int,
		) ([]domians.Pertanyaan, error)

		UpdateQuestion(
			ctx context.Context,
			question *domians.Pertanyaan,
		) (*domians.Pertanyaan, error)

		DeleteQuestionByID(
			ctx context.Context,
			id int,
		) error
	}
)

func (uc QuestionUseCase) CreateQuestion(
	ctx context.Context,
	payload CreateQuestionParam,
) (*domians.Pertanyaan, error) {
	return uc.questionRepo.CreateQuestion(
		ctx,
		&domians.Pertanyaan{
			Pertanyaan:   payload.Pertanyaan,
			OpsiJawaban:  payload.OpsiJawaban,
			JawabanBenar: payload.JawabanBenar,
			IDQuiz:       payload.IDQuiz,
		})
}

func (uc QuestionUseCase) GetAllQuestions(
	ctx context.Context,
) ([]domians.Pertanyaan, error) {
	return uc.questionRepo.GetAllQuestions(ctx)
}

func (uc QuestionUseCase) UpdateQuestion(
	ctx context.Context,
	question *domians.Pertanyaan,
) (*domians.Pertanyaan, error) {
	return uc.questionRepo.UpdateQuestion(ctx, question)
}

func (uc QuestionUseCase) DeleteQuestionByID(
	ctx context.Context,
	id int,
) error {
	return uc.questionRepo.DeleteQuestionByID(ctx, id)
}

func (uc QuestionUseCase) GetQuestionsByID(
	ctx context.Context,
	id int,
) ([]domians.Pertanyaan, error) {
	return uc.questionRepo.GetQuestionsByID(ctx, id)
}
