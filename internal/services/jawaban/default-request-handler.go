package jawaban

import (
	"net/http"
	"real_nimi_project/internal/adapter/dto"
	"real_nimi_project/internal/domians"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	ctrl ControllerInterface
}

func (rh RequestHandler) CreateJawabanPeserta(ctx *gin.Context) {
	var payload = ParticipantAnswerDTO{}
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.CreateJawabanPeserta(ctx, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetJawabanPesertaByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.GetJawabanPesertaByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (rh RequestHandler) UpdateJawabanPeserta(ctx *gin.Context) {
	var payload = ParticipantAnswerDTO{}
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.UpdateJawabanPeserta(ctx, &domians.ParticipantAnswer{
		ID:             payload.ID,
		IDUser:         payload.IDUser,
		IDQuiz:         payload.IDQuiz,
		IDPertanyaan:   payload.IDPertanyaan,
		JawabanPeserta: payload.JawabanPeserta,
		Skor:           payload.Skor,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (rh *RequestHandler) DeleteJawabanPeserta(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}

	res, err := rh.ctrl.DeleteJawabanPeserta(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
