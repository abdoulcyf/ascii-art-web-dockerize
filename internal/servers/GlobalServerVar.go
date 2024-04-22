package servers

import (
	"log/slog"
	"os"
)

// type Server struct {
// 	listener string
// }

type Error struct {
	Error string
}

var (
	errMsg string
	logMsg string
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
)

const (
	homePagePath            = "/"
	NotFound404          = "404 not found"
	getRequest              = "GET"
	methodNotAllowed        = "Method not allowed"
	badRequest              = "Bad Request"
	ErrorGeneratingAsciiArt = "Error generating ASCII art:"
	BannerNotFound          = "Banner not found"

	ascciArtPagePath = "/ascii-art"

	portAddressNumber = ":8080"
)
