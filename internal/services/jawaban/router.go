package jawaban

import (
	"real_nimi_project/internal/adapter/repository"
	"real_nimi_project/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	rh *RequestHandler
}

func NewRouter(db *gorm.DB) *Router {
	return &Router{
		rh: &RequestHandler{
			ctrl: &Controller{
				Uc: JawabanPesertaUseCase{
					jawabanPesertaRepo: repository.NewJawabanPesertaRepo(db),
				},
			},
		},
	}
}

func (r Router) Route(router *gin.RouterGroup) {
	jawabanPeserta := router.Group("/jawaban-peserta")
	jawabanPeserta.POST("/create",
		middleware.AuthMiddleware(),
		r.rh.CreateJawabanPeserta)

	jawabanPeserta.GET("/:id",
		middleware.AuthMiddlewareadmin(),
		r.rh.GetJawabanPesertaByID)

	jawabanPeserta.PUT("/update",
		middleware.AuthMiddleware(),
		r.rh.UpdateJawabanPeserta)

	jawabanPeserta.DELETE("/:id",
		middleware.AuthMiddlewareadmin(),
		r.rh.DeleteJawabanPeserta)
}
