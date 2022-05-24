package app

import (
	"os"

	"github.com/PubApiADN/pkg/logger"
	"github.com/PubApiADN/pkg/mlhandlers"
)

func StartApp() {
	router := mlhandlers.DefaultRouter()

	MapUrls(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":" + "8080"
	}

	if err := router.Run(port); err != nil {
		logger.Errorf("error running server", err)
	}
}
