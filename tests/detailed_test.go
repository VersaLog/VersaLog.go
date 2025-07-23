package main

import versalog "github.com/kayu0514/VersaLog.go/VersaLog"

// showFile false
func main() {
	logger := versalog.NewVersaLog("detailed", false, false, "VersaLog", false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")
}

// showFile true
func main() {
	logger := versalog.NewVersaLog("detailed", true, false, "VersaLog", false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// show_tag False
func main() {
	logger := versalog.NewVersaLog("detailed", false, false, "VersaLog", false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// show_tag true
func main() {
	logger := versalog.NewVersaLog("detailed", false, true, "VersaLog", false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// notice False
func main() {
	logger := versalog.NewVersaLog("detailed", false, false, "VersaLog", false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// notice True
func main() {
	logger := versalog.NewVersaLog("detailed", false, false, "VersaLog", false, true)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// enable_all true
func main() {
	logger := versalog.NewVersaLog("detailed", false, false, "VersaLog", true, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}
