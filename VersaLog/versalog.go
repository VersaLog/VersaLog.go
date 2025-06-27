package versalog

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type VersaLog struct {
	mode     string
	showFile bool
}

var COLORS = map[string]string{
	"INFO":    "\033[32m",
	"ERROR":   "\033[31m",
	"WARNING": "\033[33m",
}

var SYMBOLS = map[string]string{
	"INFO":    "[+]",
	"ERROR":   "[-]",
	"WARNING": "[!]",
}

const RESET = "\033[0m"

func NewVersaLog(mode string, showFile bool) *VersaLog {
	mode = strings.ToLower(mode)
	return &VersaLog{
		mode:     mode,
		showFile: showFile,
	}
}

func (v *VersaLog) getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (v *VersaLog) getCaller() string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return "unknown:0"
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

func (v *VersaLog) log(msg string, level string) {
	level = strings.ToUpper(level)
	color := COLORS[level]
	symbol := SYMBOLS[level]
	caller := ""

	if v.showFile || v.mode == "file" {
		caller = v.getCaller()
	}

	var output string
	switch v.mode {
	case "simple":
		if v.showFile {
			output = fmt.Sprintf("[%s]%s%s%s %s", caller, color, symbol, RESET, msg)
		} else {
			output = fmt.Sprintf("%s%s%s %s", color, symbol, RESET, msg)
		}
	case "file":
		output = fmt.Sprintf("[%s]%s[%s]%s %s", caller, color, level, RESET, msg)
	default:
		timestamp := v.getTime()
		if v.showFile {
			output = fmt.Sprintf("[%s]%s[%s]%s[%s] : %s", timestamp, color, level, RESET, caller, msg)
		} else {
			output = fmt.Sprintf("[%s]%s[%s]%s : %s", timestamp, color, level, RESET, msg)
		}
	}

	fmt.Println(output)
}

func (v *VersaLog) Info(msg string) {
	v.log(msg, "INFO")
}

func (v *VersaLog) Err(msg string) {
	v.log(msg, "ERROR")
}

func (v *VersaLog) War(msg string) {
	v.log(msg, "WARNING")
}
