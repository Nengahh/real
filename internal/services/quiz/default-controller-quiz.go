package quiz

import (
	"fmt"
	"log"
	"real_nimi_project/internal/adapter/dto"
	"real_nimi_project/internal/domians"
	"time"

	"golang.org/x/net/context"
)

type (
	Controller struct {
		Uc QuizUseCaseInterface
	}

	ControllerInterface interface {
		CreateQuiz(
			ctx context.Context,
			payload CreateQuizParam,
		) (*dto.Response, error)

		GetAllQuizzes(
			ctx context.Context,
		) (*dto.Response, error)

		UpdateQuiz(
			ctx context.Context,
			quiz *domians.Quiz,
		) (*dto.Response, error)

		DeleteQuizByID(
			ctx context.Context,
			id int,
		) (*dto.Response, error)
	}
)

func (ctrl Controller) CreateQuiz(
	ctx context.Context,
	payload CreateQuizParam,
) (*dto.Response, error) {
	start := time.Now()
	result, err := ctrl.Uc.CreateQuiz(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		result,
		"Quiz created successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) GetAllQuizzes(
	ctx context.Context,
) (*dto.Response, error) {
	start := time.Now()
	res, err := ctrl.Uc.GetAllQuizzes(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		res,
		"List of quizzes",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) UpdateQuiz(
	ctx context.Context,
	quiz *domians.Quiz,
) (*dto.Response, error) {
	start := time.Now()
	updatedQuiz, err := ctrl.Uc.UpdateQuiz(ctx, quiz)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		updatedQuiz,
		"Quiz updated successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) DeleteQuizByID(
	ctx context.Context,
	id int,
) (*dto.Response, error) {
	start := time.Now()
	err := ctrl.Uc.DeleteQuizByID(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		nil,
		"Quiz deleted successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}
