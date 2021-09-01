package logplus

import (
	"strings"
	"time"

	"github.com/fatih/color"
)

var fString = "%v/%v/%v %v:%v:%v > %v"

func Info(text ...string) {
	fText := strings.Join(text, "")
	fText += "\n"

	info := color.New(color.FgHiGreen)

	nUtc := time.Now().UTC()
	info.Printf(fString, nUtc.Year(), nUtc.Month(), nUtc.Day(), nUtc.Hour(), nUtc.Minute(), nUtc.Second())
}

func Debug(text ...string) {
	fText := strings.Join(text, "")
	fText += "\n"

	debug := color.New(color.FgHiCyan)

	nUtc := time.Now().UTC()
	debug.Printf(fString, nUtc.Year(), nUtc.Month(), nUtc.Day(), nUtc.Hour(), nUtc.Minute(), nUtc.Second())
}

func Warn(text ...string) {
	fText := strings.Join(text, "")
	fText += "\n"

	warn := color.New(color.FgHiYellow)

	nUtc := time.Now().UTC()
	warn.Printf(fString, nUtc.Year(), nUtc.Month(), nUtc.Day(), nUtc.Hour(), nUtc.Minute(), nUtc.Second())	
}

func Error(text ...string) {
	fText := strings.Join(text, "")
	fText += "\n"

	err := color.New(color.FgHiRed)

	nUtc := time.Now().UTC()
	err.Printf(fString, nUtc.Year(), nUtc.Month(), nUtc.Day(), nUtc.Hour(), nUtc.Minute(), nUtc.Second())
}

func Fatal(text ...string) {
	fText := strings.Join(text, "")
	fText += "\n"

	fatal := color.New(color.FgRed)

	nUtc := time.Now().UTC()
	fatal.Printf(fString, nUtc.Year(), nUtc.Month(), nUtc.Day(), nUtc.Hour(), nUtc.Minute(), nUtc.Second())
}