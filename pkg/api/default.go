package api

import (
	"fmt"
	"log"
	"real_nimi_project/internal/services/jawaban"
	"real_nimi_project/internal/services/pertanyaan"
	"real_nimi_project/internal/services/quiz"
	"real_nimi_project/internal/services/user"
	"real_nimi_project/pkg/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func Default() *Api {
// 	server := gin.Default()

// 	//register database
// 	db, err := db.Default() //"_" menunjukan variable db di skip
// 	if err != nil {
// 		log.Println(err)
// 		panic(fmt.Sprintf("panic at db connection: %s", err.Error()))
// 	}

// 	userRouter := user.NewRoute(db)

// 	return &Api{
// 		server:  server,
// 		routers: nil,
// 	}
// }

// func Default() *Api {
// 	server := gin.Default()
// 	sqlConn, err := db.Default()
// 	if err != nil {
// 		log.Println(err)
// 		panic(fmt.Sprintf("panic at db connection: %s", err.Error()))
// 	}

// 	var routers = []Router{
// 		user.NewRoute(sqlConn),
// 	}

// 	return &Api{
// 		server:  server,
// 		routers: routers,
// 	}
// }

func Default() *Api {
	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = false
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	server.Use(cors.New(config))

	db, err := db.Default()
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("panic at db connection: %s", err.Error()))
	}

	userRouter := user.NewRoute(db)
	quizRouter := quiz.NewRoute(db)
	pertanyaanRouter := pertanyaan.NewQuestionRouter(db)
	jawabanRouter := jawaban.NewRouter(db)

	return &Api{
		server:  server,
		routers: []Router{userRouter, quizRouter, pertanyaanRouter, jawabanRouter},
	}
}
