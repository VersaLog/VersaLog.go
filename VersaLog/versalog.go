package versalog

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/go-toast/toast"
)

type VersaLog struct {
	mode      string
	tag       string
	showFile  bool
	showtag   bool
	Notice    bool
	EnableAll bool
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

func NewVersaLog(mode string, showFile bool, showtag bool, tag string, enableAll bool, notice bool) *VersaLog {
	mode = strings.ToLower(mode)

	validModes := map[string]bool{"simple": true, "detailed": true, "file": true}
	if !validModes[mode] {
		panic(fmt.Sprintf("Invalid mode '%s' specified. Valid modes are: simple, detailed, file", mode))
	}

	if enableAll {
		showFile = true
		showtag = true
		notice = true
	}

	if mode == "file" {
		showFile = true
	}

	return &VersaLog{
		mode:      mode,
		showFile:  showFile,
		showtag:   showtag,
		tag:       tag,
		Notice:    notice,
		EnableAll: enableAll,
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

func (v *VersaLog) log(msg string, level string, tag ...string) {
	level = strings.ToUpper(level)
	color := COLORS[level]
	symbol := SYMBOLS[level]
	caller := ""
	finalTag := ""

	if len(tag) > 0 && tag[0] != "" {
		finalTag = tag[0]
	} else if v.showtag && v.tag != "" {
		finalTag = v.tag
	}

	if v.showFile || v.mode == "file" {
		caller = v.getCaller()
	}

	var output string
	switch v.mode {
	case "simple":
		if v.showFile {
			if finalTag != "" {
				output = fmt.Sprintf("[%s][%s]%s%s%s %s", caller, finalTag, color, symbol, RESET, msg)
			} else {
				output = fmt.Sprintf("[%s]%s%s%s %s", caller, color, symbol, RESET, msg)
			}
		} else {
			if finalTag != "" {
				output = fmt.Sprintf("[%s]%s%s%s %s", finalTag, color, symbol, RESET, msg)
			} else {
				output = fmt.Sprintf("%s%s%s %s", color, symbol, RESET, msg)
			}
		}
	case "file":
		output = fmt.Sprintf("[%s]%s[%s]%s %s", caller, color, level, RESET, msg)
	default:
		timestamp := v.getTime()
		output = fmt.Sprintf("[%s]%s[%s]%s", timestamp, color, level, RESET)
		if finalTag != "" {
			output += fmt.Sprintf("[%s]", finalTag)
		}
		if v.showFile {
			output += fmt.Sprintf("[%s]", caller)
		}
		output += fmt.Sprintf(" : %s", msg)
	}

	fmt.Println(output)

	if v.Notice && (level == "ERROR" || level == "CRITICAL") {
		toastMessage := toast.Notification{
			AppID:   "VersaLog",
			Title:   fmt.Sprintf("%s Log notice", level),
			Message: msg,
		}
		toastMessage.Push()
	}
}

func (v *VersaLog) Info(msg string, tag ...string) {
	v.log(msg, "INFO", tag...)
}

func (v *VersaLog) Error(msg string, tag ...string) {
	v.log(msg, "ERROR", tag...)
}

func (v *VersaLog) Warning(msg string, tag ...string) {
	v.log(msg, "WARNING", tag...)
}

func (v *VersaLog) Debug(msg string, tag ...string) {
	v.log(msg, "DEBUG", tag...)
}

func (v *VersaLog) Critical(msg string, tag ...string) {
	v.log(msg, "CRITICAL", tag...)
}
