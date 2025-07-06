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
	"INFO":     "\033[32m",
	"ERROR":    "\033[31m",
	"WARNING":  "\033[33m",
	"DEBUG":    "\033[36m",
	"CRITICAL": "\033[35m",
}

var SYMBOLS = map[string]string{
	"INFO":     "[+]",
	"ERROR":    "[-]",
	"WARNING":  "[!]",
	"DEBUG":    "[D]",
	"CRITICAL": "[C]",
}

const RESET = "\033[0m"

func NewVersaLog(mode string, showFile bool) *VersaLog {
	mode = strings.ToLower(mode)

	validModes := map[string]bool{"simple": true, "detailed": true, "file": true}
	if !validModes[mode] {
		panic(fmt.Sprintf("Invalid mode '%s' specified. Valid modes are: simple, detailed, file", mode))
	}

	if mode == "file" {
		showFile = true
	}

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

func (v *VersaLog) Error(msg string) {
	v.log(msg, "ERROR")
}

func (v *VersaLog) Warning(msg string) {
	v.log(msg, "WARNING")
}

func (v *VersaLog) Debug(msg string) {
	v.log(msg, "DEBUG")
}

func (v *VersaLog) Critical(msg string) {
	v.log(msg, "CRITICAL")
}
