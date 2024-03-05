package quiz

import (
	"net/http"
	"real_nimi_project/internal/adapter/dto"
	"real_nimi_project/internal/domians"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	RequestHandler struct {
		ctrl ControllerInterface
	}
)

func (rh RequestHandler) CreateQuiz(ctx *gin.Context) {
	var payload = CreateQuizParam{}
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.CreateQuiz(ctx, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetAllQuizzes(ctx *gin.Context) {
	res, err := rh.ctrl.GetAllQuizzes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (rh RequestHandler) UpdateQuiz(ctx *gin.Context) {
	var payload = UpdateQuizParam{}
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.UpdateQuiz(ctx, &domians.Quiz{
		ID:           payload.ID,
		Judul:        payload.Judul,
		Deskripsi:    payload.Deskripsi,
		WaktuMulai:   payload.WaktuMulai,
		WaktuSelesai: payload.WaktuSelesai,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (rh *RequestHandler) DeleteQuiz(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}

	res, err := rh.ctrl.DeleteQuizByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
