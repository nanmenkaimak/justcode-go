package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/dbs/postgres"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/handlers"
	"log"
)

const portNumber = ":8080"

func main() {
	db, err := postgres.New()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)

	r := gin.Default()
	r.POST("/", handlers.Repo.CreateGist)
	r.GET("/all", handlers.Repo.GetAllGists)
	r.GET("/:id", handlers.Repo.GetGistByID)
	r.PUT("/:id", handlers.Repo.UpdateGistByID)
	r.DELETE("/:id", handlers.Repo.DeleteGistByID)

	r.Run(portNumber)
}
