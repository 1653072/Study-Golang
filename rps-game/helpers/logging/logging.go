package logging

import (
	"fmt"
	"golangdemo/rps-game/configs/log-conf"
	"golangdemo/rps-game/configs/system-path"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"time"
)

var SysLog *log.Logger

func InitializeLogging() {
	var logFileName = "app.log"
	var logFilePath = SystemPath.LogContainerFolderPath + "/" + logFileName

	res, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening \"logging\" file: %v", err)
		os.Exit(1)
	}
	defer res.Close()

	SysLog = log.New(res, "", log.Ldate|log.Ltime)
	SysLog.SetOutput(&lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    LogConf.MaxSize,    // Megabytes after which new file is created
		MaxBackups: LogConf.MaxBackups, // Number of backups
		MaxAge:     LogConf.MaxAge,     // Days
		Compress:   LogConf.Compress,
	})
}

func FormatResult(logType uint8, errCode error, msg string) string {
	var status, result string

	if errCode != nil && (logType == LogConf.Debug || logType == LogConf.Info || logType == LogConf.Warn) {
		logType = LogConf.Error
	}

	switch logType {
	case LogConf.Debug:
		status = "DEBUG"
		break
	case LogConf.Info:
		status = "INFO"
		break
	case LogConf.Warn:
		status = "WARN"
		break
	case LogConf.Error:
		status = "ERROR"
		break
	case LogConf.Fatal:
		status = "FATAL"
		break
	case LogConf.Panic:
		status = "PANIC"
		break
	default:
		status = "INFO"
	}

	if errCode != nil {
		result = fmt.Sprintf("[%v][%s] Message: %s", status, errCode.Error(), msg)
	} else {
		result = fmt.Sprintf("[%v] Message: %s", status, msg)
	}

	now := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("%s %s\n", now, result)

	return result
}
