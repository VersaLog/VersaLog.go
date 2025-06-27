package main

import versalog "github.com/kayu0514/VersaLog.go/VersaLog"

func main() {
	logger := versalog.NewVersaLog("file")

	logger.Info("Everything is fine.")
	logger.Err("Something went wrong.")
	logger.War("This is a warning.")
}
