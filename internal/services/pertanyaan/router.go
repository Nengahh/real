package pertanyaan

import (
	"real_nimi_project/internal/adapter/repository"
	"real_nimi_project/pkg/middleware"

	// "real_nimi_project/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	QuestionRouter struct {
		rq *QuestionRequestHandler
	}
)

func NewQuestionRouter(
	db *gorm.DB,
) *QuestionRouter {
	return &QuestionRouter{rq: &QuestionRequestHandler{
		ctrl: &QuestionController{
			questionUc: QuestionUseCase{
				questionRepo: repository.NewQuestionRepo(db),
			},
		},
	}}

}

func (r QuestionRouter) Route(router *gin.RouterGroup) {
	question := router.Group("/question")
	question.POST(
		"/create",
		middleware.AuthMiddlewareadmin(),
		r.rq.CreateQuestion,
	)

	question.GET(
		"/all",
		middleware.AuthMiddleware(),
		r.rq.GetAllQuestions,
	)

	question.PUT(
		"/update",
		middleware.AuthMiddlewareadmin(),
		r.rq.UpdateQuestion,
	)

	question.GET(
		"/:id",
		// middleware.AuthMiddlewareadmin(),
		r.rq.GetQuestionsByID,
	)

	question.DELETE(
		"/:id",
		middleware.AuthMiddlewareadmin(),
		r.rq.DeleteQuestion,
	)
}
