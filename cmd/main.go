package main

import (
	"context"

	"github.com/julienschmidt/httprouter"
	"github.com/quazar2000/ca-library-app/internal/composites"
	"github.com/quazar2000/ca-library-app/pkg/logging"
)

func main() {
	// entry point
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("config initializing")
	cfg := config.GetConfig()

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Ingo("create mongodb composite")
	mongoDBC, err := composites.NewMongoDBComposite(context.Background(), "", "", "", "", "", "")
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}

	logger.Info("author composite initializing")
	authorComposite, err := composites.NewAuthorComposite(mongoDBC)
	if err != nil {
		logger.Fatal("author composite failed")
	}
	authorComposite.Handler.Register(router)

	logger.Info("book composite initializing")
	bookComposite, err := composites.NewBookComposite(mongoDBC)
	if err != nil {
		logger.Fatal("book composite failed")
	}

	bookComposite.Handler.Register(router)

	// start
}
