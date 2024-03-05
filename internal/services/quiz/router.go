package quiz

import (
	"real_nimi_project/internal/adapter/repository"
	"real_nimi_project/pkg/middleware"

	// "real_nimi_project/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Router struct {
		rq *RequestHandler
	}
)

func NewRoute(
	db *gorm.DB,
) *Router {
	return &Router{rq: &RequestHandler{
		ctrl: &Controller{
			Uc: QuizUseCase{
				quizRepo: repository.NewQuizRepo(db),
			},
		},
	}}

}

func (r Router) Route(router *gin.RouterGroup) {
	quiz := router.Group("/quiz")
	quiz.POST(
		"/create",
		middleware.AuthMiddlewareadmin(),
		r.rq.CreateQuiz,
	)

	quiz.GET(
		"/all",
		// middleware.AuthMiddleware(),
		r.rq.GetAllQuizzes,
	)

	quiz.PUT(
		"/update",
		middleware.AuthMiddlewareadmin(),
		r.rq.UpdateQuiz,
	)

	quiz.DELETE(
		"/:id",
		middleware.AuthMiddlewareadmin(),
		r.rq.DeleteQuiz,
	)
}
