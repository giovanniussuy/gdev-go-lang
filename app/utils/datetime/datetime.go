package datetime

import (
	"fmt"
	"time"
)

// WARN: Numeros de pattern são especificos determinados pela linguagem.
const PatternDateBr = "02/01/2006"
const PatternHoursAndMinute = "15:04"
const PatternHoursAndMinuteAndSeconds = "15:04:05"
const PatternDateBrAndHoursAndMinuteAndSeconds = "02/01/2006 15:04:05"

// func initi executa em tempo de compilação
func init() {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("ERROR: loading location: ", err)
		return
	}
	time.Local = loc
}

func Parse(Pattern string, Time string) time.Time {
	parsed, _ := time.Parse(Pattern, Time)
	return parsed
}

func TimeNowPatternDateBr() string {
	return time.Now().Format(PatternDateBr)
}

func TimeNowPatternHoursAndMinute() string {
	return time.Now().Format(PatternHoursAndMinute)
}

func TimeNowPatternHoursAndMinuteAndSeconds() string {
	return time.Now().Format(PatternHoursAndMinuteAndSeconds)
}

func TimeNowPatternDateBrAndHoursAndMinuteAndSeconds() string {
	return time.Now().Format(PatternDateBrAndHoursAndMinuteAndSeconds)
}
