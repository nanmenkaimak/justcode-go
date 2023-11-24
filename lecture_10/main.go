package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/dbs/postgres"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/dbs/redis"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/handlers"
	"log"
	"net/http"
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

	dbRedis, err := redis.New()
	if err != nil {
		log.Fatal()
	}

	repo := handlers.NewRepo(db, dbRedis)
	handlers.NewHandlers(repo)

	r := gin.Default()
	r.POST("/", handlers.Repo.CreateGist)
	r.GET("/all", handlers.Repo.GetAllGists)
	r.GET("/:id", handlers.Repo.GetGistByID)
	r.PUT("/:id", handlers.Repo.UpdateGistByID)
	r.DELETE("/:id", handlers.Repo.DeleteGistByID)

	go func() {
		r.Run(portNumber)
	}()
	chiRouter := chi.NewRouter()
	chiRouter.Mount("/debug", middleware.Profiler())
	go func() {
		http.ListenAndServe(":8081", chiRouter)
	}()
	select {}
}
