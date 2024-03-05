package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

// // igunakan untuk mengelola server dan router-rotor yang terkait dengan API
// type (
// 	Api struct {
// 		server  *gin.Engine
// 		routers []Router
// 	}

// 	Router interface {
// 		Router(handler *gin.RouterGroup)
// 	}
// )

// func (a Api) Start() error {
// 	//mengelompokkan grup router di server dengan path root ("/")
// 	root := a.server.Group("/")
// 	for _, router := range a.routers {
// 		router.Router(root)
// 	}

// 	err := a.server.Run(":8000")
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	return err

// }

// type (
//     Api struct {
//         server  *gin.Engine
//         routers []Router
//     }

//     Router interface {
//         Route(handler *gin.RouterGroup)
//     }
// )

// func NewApi(db *gorm.DB) *Api {
//     server := gin.Default()
//     userRouter := user.NewRoute(db)

//     return &Api{
//         server:  server,
//         routers: []Router{userRouter},
//     }
// }

// func (a Api) Start() error {
//     root := a.server.Group("/")
//     for _, router := range a.routers {
//         router.Route(root)
//     }

//     err := a.server.Run(":8000")
//     if err != nil {
//         log.Println(err)
//         return err
//     }
//     return err
// }

type (
	Api struct {
		server  *gin.Engine
		routers []Router
	}

	Router interface {
		Route(handler *gin.RouterGroup)
	}
)

func (a Api) Start() error {
	root := a.server.Group("/")
	for _, router := range a.routers {
		router.Route(root)
	}

	err := a.server.Run(":8000")
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
