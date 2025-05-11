package main

import (
	logger "github.com/yourpov/logrite"
)

func main() {
	logger.SetConfig(logger.Config{
		ShowIcons:    true,
		UppercaseTag: true,
		UseColors:    true,
	})

	// ── General Logs ────────────────────────────────────────

	logger.Info("This is an info log")
	logger.Warn("This is a warning log")
	logger.Error("This is an error log")
	logger.Success("This is a success log")
	logger.Debug("This is a debug log\n")

	// ── Predefined Tags ────────────────────────────────────────

	logger.WebStart(8080)
	logger.Cmd("clear", "15ms", "root")
	logger.Login("admin", "127.0.0.1\n")

	// ── Custom emoji + tag ────────────────────────────────────────

	logger.Custom("🕸️", "custom", "This is a custom log with a custom emoji", logger.Green, logger.BgBlack)

	// ── Fully manual ────────────────────────────────────────

	logger.Log("manual", " This is a manually made log\n", logger.BgWhite, logger.White)
}
