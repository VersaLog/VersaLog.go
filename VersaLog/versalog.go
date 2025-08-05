package versalog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/go-toast/toast"
)

type VersaLog struct {
	mode       string
	tag        string
	showFile   bool
	showTag    bool
	Notice     bool
	EnableAll  bool
	AllSave    bool
	SaveLevels []string
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

func NewVersaLog(mode string, showFile bool, showTag bool, tag string, enableAll bool, notice bool, allSave bool, saveLevels []string) *VersaLog {
	mode = strings.ToLower(mode)

	validModes := map[string]bool{"simple": true, "simple2": true, "detailed": true, "file": true}
	if !validModes[mode] {
		panic(fmt.Sprintf("Invalid mode '%s' specified. Valid modes are: simple, simple2, detailed, file", mode))
	}

	if enableAll {
		showFile = true
		showTag = true
		notice = true
		allSave = true
	}

	if mode == "file" {
		showFile = true
	}

	validSaveLevels := []string{"INFO", "ERROR", "WARNING", "DEBUG", "CRITICAL"}
	if allSave {
		if len(saveLevels) == 0 {
			saveLevels = append([]string{}, validSaveLevels...)
		} else {
			for _, l := range saveLevels {
				found := false
				for _, v := range validSaveLevels {
					if l == v {
						found = true
						break
					}
				}
				if !found {
					panic(fmt.Sprintf("Invalid saveLevels specified. Valid levels are: %v", validSaveLevels))
				}
			}
		}
	}

	return &VersaLog{
		mode:       mode,
		showFile:   showFile,
		showTag:    showTag,
		tag:        tag,
		Notice:     notice,
		EnableAll:  enableAll,
		AllSave:    allSave,
		SaveLevels: saveLevels,
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

func (v *VersaLog) saveLog(logText string, level string) {
	if !v.AllSave {
		return
	}
	found := false
	for _, l := range v.SaveLevels {
		if l == level {
			found = true
			break
		}
	}
	if !found {
		return
	}
	logDir := filepath.Join(filepath.Dir(filepath.Dir("./VersaLog/versalog.go")), "log")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, 0755)
	}
	logFile := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString(logText + "\n")
}

func (v *VersaLog) log(msg string, level string, tag ...string) {
	level = strings.ToUpper(level)
	color := COLORS[level]
	symbol := SYMBOLS[level]
	caller := ""
	finalTag := ""

	if len(tag) > 0 && tag[0] != "" {
		finalTag = tag[0]
	} else if v.showTag && v.tag != "" {
		finalTag = v.tag
	}

	if v.showFile || v.mode == "file" {
		caller = v.getCaller()
	}

	var output, plain string
	switch v.mode {
	case "simple":
		if v.showFile {
			if finalTag != "" {
				output = fmt.Sprintf("[%s][%s]%s%s%s %s", caller, finalTag, color, symbol, RESET, msg)
				plain = fmt.Sprintf("[%s][%s]%s %s", caller, finalTag, symbol, msg)
			} else {
				output = fmt.Sprintf("[%s]%s%s%s %s", caller, color, symbol, RESET, msg)
				plain = fmt.Sprintf("[%s]%s %s", caller, symbol, msg)
			}
		} else {
			if finalTag != "" {
				output = fmt.Sprintf("[%s]%s%s%s %s", finalTag, color, symbol, RESET, msg)
				plain = fmt.Sprintf("[%s]%s %s", finalTag, symbol, msg)
			} else {
				output = fmt.Sprintf("%s%s%s %s", color, symbol, RESET, msg)
				plain = fmt.Sprintf("%s %s", symbol, msg)
			}
		}
	case "simple2":
		timestamp := v.getTime()
		if v.showFile {
			if finalTag != "" {
				output = fmt.Sprintf("[%s] [%s][%s]%s%s%s %s", timestamp, caller, finalTag, color, symbol, RESET, msg)
				plain = fmt.Sprintf("[%s] [%s][%s]%s %s", timestamp, caller, finalTag, symbol, msg)
			} else {
				output = fmt.Sprintf("[%s] [%s]%s%s%s %s", timestamp, caller, color, symbol, RESET, msg)
				plain = fmt.Sprintf("[%s] [%s]%s %s", timestamp, caller, symbol, msg)
			}
		} else {
			output = fmt.Sprintf("[%s] %s%s%s %s", timestamp, color, symbol, RESET, msg)
			plain = fmt.Sprintf("[%s] %s %s", timestamp, symbol, msg)
		}
	case "file":
		output = fmt.Sprintf("[%s]%s[%s]%s %s", caller, color, level, RESET, msg)
		plain = fmt.Sprintf("[%s][%s] %s", caller, level, msg)
	default:
		timestamp := v.getTime()
		output = fmt.Sprintf("[%s]%s[%s]%s", timestamp, color, level, RESET)
		plain = fmt.Sprintf("[%s][%s]", timestamp, level)
		if finalTag != "" {
			output += fmt.Sprintf("[%s]", finalTag)
			plain += fmt.Sprintf("[%s]", finalTag)
		}
		if v.showFile {
			output += fmt.Sprintf("[%s]", caller)
			plain += fmt.Sprintf("[%s]", caller)
		}
		output += fmt.Sprintf(" : %s", msg)
		plain += fmt.Sprintf(" : %s", msg)
	}

	fmt.Println(output)
	v.saveLog(plain, level)

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
