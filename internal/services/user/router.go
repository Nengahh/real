// package user

// import (
// 	"real_nimi_project/internal/adapter/repository"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type (
// 	Router struct {
// 		rq *RequestHandler
// 	}
// )

// func NewRoute(
// 	db *gorm.DB,
// ) *Router {
// 	return &Router{rq: &RequestHandler{
// 		ctrl: &Controller{
// 			Uc: UserUseCase{
// 				userRepo: repository.NewUserRepo(db),
// 			},
// 		},
// 	}}

// }

// func (r Router) Router(router *gin.RouterGroup) {
// 	user := router.Group("/user")
// 	user.POST(
// 		"/register",
// 		r.rq.Register,
// 	)

// 	user.GET(
// 		"/all",
// 		r.rq.GetAll,
// 	)

// 	user.POST(
// 		"/login",
// 		r.rq.LoginUser,
// 	)

// }

package user

import (
	"real_nimi_project/internal/adapter/repository"
	"real_nimi_project/pkg/middleware"

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
			Uc: UserUseCase{
				userRepo: repository.NewUserRepo(db),
			},
		},
	}}

}

func (r Router) Route(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.POST(
		"/register",
		r.rq.Register,
	)

	user.GET(
		"/all",
		middleware.AuthMiddlewareadmin(),
		r.rq.GetAll,
	)

	user.POST(
		"/login",
		// middleware.JWTMiddleware(),
		r.rq.LoginUser,
	)

	user.DELETE(
		"/:name",
		middleware.AuthMiddlewareadmin(),
		r.rq.DeleteUser,
	)

	user.GET(
		"/logout",
		Logout,
	)
}
