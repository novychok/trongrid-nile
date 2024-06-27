package log

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return l
}
