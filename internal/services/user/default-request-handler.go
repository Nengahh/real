package user

import (
	"net/http"
	"real_nimi_project/internal/adapter/dto"

	"github.com/gin-gonic/gin"
)

type (
	RequestHandler struct {
		ctrl ContrllerInterface
	}
)

func (rh RequestHandler) Register(ctx *gin.Context) {
	var payload = RegisterUserParam{}
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.Register(ctx, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)

}

func (rh RequestHandler) GetAll(ctx *gin.Context) {
	res, err := rh.ctrl.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (rh RequestHandler) LoginUser(ctx *gin.Context) {
	var payload = LoginParam{}
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.LoginUser(ctx, payload.Email, payload.Password, payload.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.SetCookie(
		"token",
		res.AccessToken,
		3600,
		"/",
		"",
		false,
		true,
	)

	ctx.JSON(http.StatusOK, res)
}

func (rh *RequestHandler) DeleteUser(c *gin.Context) {
	name := c.Param("name")

	res, err := rh.ctrl.DeleteUserByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
