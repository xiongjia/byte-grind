package main

import (
	"log/slog"

	"byte-grind/wheat/pkg/util"
)

func main() {
	slog.SetDefault(util.NewLogger(
		util.LoggerWithLevel(util.LOG_LEVEL_DEBUG),
		util.LoggerWithSource(true),
	))
	slog.Debug("test", slog.String("test", "t"))
}
