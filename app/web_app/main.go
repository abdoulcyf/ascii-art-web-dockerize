package main

import (
	// "log/slog"
	// "os"

	//"errors"
	"log/slog"
	"os"

	"github.com/ediallocyf/asciiartweb/internal/servers"
)

var (
	opts   = &slog.HandlerOptions{Level: slog.LevelDebug}
	logger = slog.New(slog.NewTextHandler(os.Stderr, opts))

	errMsg string
	logMsg string
)


func main() {
	errLoadS := servers.LoadingServerS()
	if errLoadS != nil {
		errMsg = "----<-LoadingServerS------<--ListenAndServe----" + errLoadS.Error()

		logMsg = " Application failed to start:"
		logger.Error(logMsg + errMsg)
	}

	logger.Info("Server running at port 8080")
}
