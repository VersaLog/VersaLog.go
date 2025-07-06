package main

import versalog "github.com/kayu0514/VersaLog.go/VersaLog"

// showFile false
func main() {
	logger := versalog.NewVersaLog("detailed", false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")
}

// showFile true
func main() {
	logger := versalog.NewVersaLog("detailed", true)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}
