package main

import (
	"context"

	author3 "github.com/quazar2000/ca-library-app/internal/adapters/api/author"
	author2 "github.com/quazar2000/ca-library-app/internal/adapters/db/author"
	"github.com/quazar2000/ca-library-app/internal/composites"

	"github.com/quazar2000/ca-library-app/internal/domain/author"
)

func main() {
	// entry point
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("config initializing")
	cfg := config.GetConfig()

	logger.Ingo("create mongodb composite")
	mongoDBC, err := composites.NewMongoDBComposite(context.Background(), "", "", "", "", "", "")
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}

	logger.Info("author composite initializing")
	authorComposite, err := composites.NewAuthorComposite()
	if err != nil {
		logger.Fatal("author composite failed")
	}

	logger.Info("book composite initializing")
	bookComposite, err := composites.NewBookComposite(mongoDBC)
	if err != nil {
		logger.Fatal("book composite failed")
	}

	authorStorage := author2.NewStorage()
	authorService := author.NewService(authorStorage)
	authorHanlder := author3.NewHandler(authorService)

}
