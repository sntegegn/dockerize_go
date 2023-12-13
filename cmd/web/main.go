package main

import (
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"github.com/sntegegn/dockerize_go/models"
)

type application struct {
	logger      *slog.Logger
	redisClient *redis.Client
	user        models.UserModel
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := openDB()
	if err != nil {
		logger.Error("error accessing db", "db", err)
		os.Exit(1)
	}
	defer db.Close()

	user := models.UserModel{DB: db}

	app := &application{
		logger:      logger,
		redisClient: getRedisClient(),
		user:        user,
	}

	err = user.CreateTable()
	if err != nil {
		logger.Error("Error creating a new table", "err", err)
		os.Exit(1)
	}

	app.logger.Info("db stats are: ", "stats", db.Stats())

	srv := &http.Server{
		Addr:    ":8081",
		Handler: app.routes(),
	}

	app.logger.Info("Listening on localhost:8080")
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
