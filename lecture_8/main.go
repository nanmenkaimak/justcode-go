package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/nanmenkaimak/justcode-go/lecture_8/internal/dbs/postgres"
	"github.com/nanmenkaimak/justcode-go/lecture_8/internal/handlers"
	"github.com/nanmenkaimak/justcode-go/lecture_8/internal/middleware"
	"log"
)

const portNumber = ":8080"

func main() {
	db, err := postgres.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)

	r := gin.Default()
	auth := r.Group("/user")
	{
		auth.POST("/signup", handlers.Repo.SignUp)
		auth.POST("/login", handlers.Repo.Login)
	}

	user := r.Group("/user/v2")
	{
		user.Use(middleware.AuthMiddleware())
		user.PUT("/update/:id", handlers.Repo.UpdateUserByID)
		user.DELETE("/delete/:id", handlers.Repo.DeleteUserByID)
	}
	r.Run(portNumber)
}
