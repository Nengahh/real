package pertanyaan

import (
	"fmt"
	"log"
	"real_nimi_project/internal/adapter/dto"
	"real_nimi_project/internal/domians"
	"time"

	"golang.org/x/net/context"
)

type (
	QuestionController struct {
		questionUc QuestionUseCaseInterface
	}

	QuestionControllerInterface interface {
		CreateQuestion(
			ctx context.Context,
			payload CreateQuestionParam,
		) (*dto.Response, error)

		GetAllQuestions(
			ctx context.Context,
		) (*dto.Response, error)

		GetQuestionsByID(
			ctx context.Context,
			id int) (*dto.Response, error)

		UpdateQuestion(
			ctx context.Context,
			question *domians.Pertanyaan,
		) (*dto.Response, error)

		DeleteQuestionByID(
			ctx context.Context,
			id int,
		) (*dto.Response, error)
	}
)

func (ctrl QuestionController) CreateQuestion(
	ctx context.Context,
	payload CreateQuestionParam,
) (*dto.Response, error) {
	start := time.Now()
	result, err := ctrl.questionUc.CreateQuestion(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		result,
		"Question created successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl QuestionController) GetAllQuestions(
	ctx context.Context,
) (*dto.Response, error) {
	start := time.Now()
	res, err := ctrl.questionUc.GetAllQuestions(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		res,
		"List of questions",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl QuestionController) UpdateQuestion(
	ctx context.Context,
	question *domians.Pertanyaan,
) (*dto.Response, error) {
	start := time.Now()
	updatedQuestion, err := ctrl.questionUc.UpdateQuestion(ctx, question)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		updatedQuestion,
		"Question updated successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl QuestionController) GetQuestionsByID(
	ctx context.Context,
	id int,
) (*dto.Response, error) {
	start := time.Now()
	questions, err := ctrl.questionUc.GetQuestionsByID(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		questions,
		"List of questions for the quiz",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl QuestionController) DeleteQuestionByID(
	ctx context.Context,
	id int,
) (*dto.Response, error) {
	start := time.Now()
	err := ctrl.questionUc.DeleteQuestionByID(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		nil,
		"Question deleted successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}
