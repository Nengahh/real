package pertanyaan

import (
	"net/http"
	"real_nimi_project/internal/adapter/dto"
	"real_nimi_project/internal/domians"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	QuestionRequestHandler struct {
		ctrl QuestionControllerInterface
	}
)

func (rh QuestionRequestHandler) CreateQuestion(ctx *gin.Context) {
	var payload = CreateQuestionParam{}
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.CreateQuestion(ctx, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (rh QuestionRequestHandler) GetAllQuestions(ctx *gin.Context) {
	res, err := rh.ctrl.GetAllQuestions(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (rh QuestionRequestHandler) UpdateQuestion(ctx *gin.Context) {
	var payload = UpdateQuestionParam{}
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.UpdateQuestion(ctx, &domians.Pertanyaan{
		ID:           payload.ID,
		Pertanyaan:   payload.Pertanyaan,
		OpsiJawaban:  payload.OpsiJawaban,
		JawabanBenar: payload.JawabanBenar,
		IDQuiz:       payload.IDQuiz,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (rh *QuestionRequestHandler) DeleteQuestion(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}

	res, err := rh.ctrl.DeleteQuestionByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (rh QuestionRequestHandler) GetQuestionsByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.GetQuestionsByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
