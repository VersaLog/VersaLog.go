package main

import versalog "github.com/kayu0514/VersaLog.go/VersaLog"

func main() {
	logger := versalog.NewVersaLog("file", false, false, "VersaLog", false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")
}
