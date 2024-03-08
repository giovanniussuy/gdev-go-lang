package loggerMy

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/giovanniussuy/gdev-go-lang/app/utils/datetime"
)

type LogContent struct {
	LogInit string   `json:"logInit"`
	App     string   `json:"app"`
	File    string   `json:"file"`
	TraceId string   `json:"traceId"`
	IpTrace string   `json:"ipTrace"`
	Info    []string `json:"info"`
	Debug   []string `json:"debug"`
	Warn    []string `json:"warn"`
	Error   ErrorLog `json:"error"`
}

type ErrorLog struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	File    string `json:"file"`
	Time    string `json:"time"`
	Trace   string `json:"trace"`
}

var (
	VarLogContent  LogContent
	VarEnvironment string = os.Getenv("ENVIRONMENT")
	AppName        string = os.Getenv("APP_NAME")
)

func Init(traceid string, iptrace string) {
	_, path, line, _ := runtime.Caller(1)
	VarLogContent = LogContent{}
	VarLogContent.LogInit = datetime.TimeNowPatternDateBrAndHoursAndMinuteAndSeconds()
	VarLogContent.App = AppName
	VarLogContent.File = fmt.Sprintf("%s:%d", filepath.Base(path), line)
	VarLogContent.TraceId = traceid
	VarLogContent.IpTrace = iptrace
}

func Debug(message string) {
	if VarEnvironment == "dev" || VarEnvironment == "hom" {
		_, path, line, _ := runtime.Caller(1)
		logMessage := fmt.Sprintf("%s | %s | %s", fmt.Sprintf("%s:%d", filepath.Base(path), line), message, datetime.TimeNowPatternDateBrAndHoursAndMinuteAndSeconds())
		VarLogContent.Debug = append(VarLogContent.Debug, logMessage)
	}
}

func Info(message string) {
	_, path, line, _ := runtime.Caller(1)
	logMessage := fmt.Sprintf("%s | %s | %s", fmt.Sprintf("%s:%d", filepath.Base(path), line), message, datetime.TimeNowPatternDateBrAndHoursAndMinuteAndSeconds())
	VarLogContent.Info = append(VarLogContent.Info, logMessage)
}

func Warn(message string) {
	_, path, line, _ := runtime.Caller(1)
	logMessage := fmt.Sprintf("%s | %s | %s", fmt.Sprintf("%s:%d", filepath.Base(path), line), message, datetime.TimeNowPatternDateBrAndHoursAndMinuteAndSeconds())
	VarLogContent.Warn = append(VarLogContent.Warn, logMessage)
}

func Error(code string, message string, trace string) {
	_, path, line, _ := runtime.Caller(1)
	VarLogContent.Error.Code = code
	VarLogContent.Error.Message = message
	VarLogContent.File = fmt.Sprintf("%s:%d", filepath.Base(path), line)
	VarLogContent.Error.Time = datetime.TimeNowPatternDateBrAndHoursAndMinuteAndSeconds()
	VarLogContent.Error.Trace = trace
	Print()
}

func Print() {
	json, _ := json.Marshal(VarLogContent)
	if VarLogContent.Error.Code != "" {
		log.Printf("ERROR: %s", json)

	} else if len(VarLogContent.Warn) > 0 {
		log.Printf("WARN: %s", json)

	} else if len(VarLogContent.Debug) > 0 {
		log.Printf("DEBUG: %s", json)

	} else if len(VarLogContent.Info) > 0 {
		log.Printf("INFO: %s", json)

	}
}
