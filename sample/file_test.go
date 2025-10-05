package main

import versalog "github.com/VersaLog/VersaLog.go/VersaLog"

func main() {
	logger := versalog.NewVersaLog("file", false, false, "VersaLog", false, false, false, []string{}, false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")
}
