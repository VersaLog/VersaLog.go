package main

import versalog "github.com/kayu0514/VersaLog.go/VersaLog"

// showFile false
func main() {
	logger := versalog.NewVersaLog("simple", false)

	logger.Info("Everything is fine.")
	logger.Err("Something went wrong.")
	logger.War("This is a warning.")
}

// showFile true
func main() {
	logger := versalog.NewVersaLog("simple", true)

	logger.Info("Everything is fine.")
	logger.Err("Something went wrong.")
	logger.War("This is a warning.")
}
