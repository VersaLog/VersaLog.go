package main

import versalog "github.com/VersaLog/VersaLog.go/VersaLog"

func main() {
	logger := versalog.NewVersaLog("detailed", false, false, "VersaLog", false, false, true, nil, false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")
}
