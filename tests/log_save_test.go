package main

import versalog "github.com/kayu0514/VersaLog.go/VersaLog"

func main() {
	logger := versalog.NewVersaLog("detailed", false, false, "VersaLog", false, false, true)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")
}
