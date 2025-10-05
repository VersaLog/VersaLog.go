package main

import versalog "github.com/VersaLog/VersaLog.go/VersaLog"

// showFile false
func main() {
	logger := versalog.NewVersaLog("simple", false, false, "VersaLog", false, false, false, []string{}, false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")
}

// showFile true
func main() {
	logger := versalog.NewVersaLog("simple", true, false, "VersaLog", false, false, false, []string{}, false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// show_tag False
func main() {
	logger := versalog.NewVersaLog("simple", false, false, "VersaLog", false, false, false, []string{}, false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// show_tag true
func main() {
	logger := versalog.NewVersaLog("simple", false, true, "VersaLog", false, false, false, []string{}, false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// notice False
func main() {
	logger := versalog.NewVersaLog("simple", false, false, "VersaLog", false, false, false, []string{}, false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// notice True
func main() {
	logger := versalog.NewVersaLog("simple", false, false, "VersaLog", false, true, false, []string{}, false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// silent False
func main() {
	logger := versalog.NewVersaLog("simple", false, false, "VersaLog", false, true, false, []string{}, false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// silent True
func main() {
	logger := versalog.NewVersaLog("simple", false, false, "VersaLog", false, true, false, []string{}, true, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}

// enable_all true
func main() {
	logger := versalog.NewVersaLog("simple", false, false, "VersaLog", true, false, false, []string{}, false, false)

	logger.Info("info")
	logger.Error("error")
	logger.Warning("warning.")
	logger.Debug("debug")
	logger.Critical("critical")

}
