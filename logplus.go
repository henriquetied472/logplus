package logplus

import (
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var infoString = "%v/%v/%v %v:%v:%v [INFO] > %v"
var debugString = "%v/%v/%v %v:%v:%v [DEBUG] > %v"
var warnString = "%v/%v/%v %v:%v:%v [WARNING] > %v"
var errorString = "%v/%v/%v %v:%v:%v [ERROR] > %v"
var fatalString = "%v/%v/%v %v:%v:%v [FATAL] > %v"

func checkMonth(m time.Month) string {
	sm := m.String()
	var res string

	switch sm {
	case "January":
		res = "01"
	case "February":
		res = "02"
	case "March":
		res = "03"
	case "April":
		res = "04"
	case "May":
		res = "05"
	case "June":
		res = "06"
	case "July":
		res = "07"
	case "August":
		res = "08"
	case "September":
		res = "09"
	case "October":
		res = "10"
	case "November":
		res = "11"
	case "December":
		res = "12"
	}

	return res
}

func checkDigits(digit int) string {
	sDigit := strconv.Itoa(digit)

	if len(sDigit) == 1 {
		sDigit = "0" + sDigit
	}

	return sDigit
}

func Info(text ...string) {
	fText := strings.Join(text, "")
	fText += "\n"

	info := color.New(color.FgHiGreen)

	ct := time.Now()
	y := ct.Year()
	mo := checkMonth(ct.Month())
	d := checkDigits(int(ct.Day()))
	h := checkDigits(int(ct.Day()))
	mi := checkDigits(int(ct.Minute()))
	s := checkDigits(int(ct.Second()))

	info.Printf(infoString, y, mo, d, h, mi, s, fText)
}

func Debug(text ...string) {
	fText := strings.Join(text, "")
	fText += "\n"

	debug := color.New(color.FgHiCyan)

	ct := time.Now()
	y := ct.Year()
	mo := checkMonth(ct.Month())
	d := checkDigits(int(ct.Day()))
	h := checkDigits(int(ct.Day()))
	mi := checkDigits(int(ct.Minute()))
	s := checkDigits(int(ct.Second()))
	debug.Printf(debugString, y, mo, d, h, mi, s, fText)
}

func Warn(text ...string) {
	fText := strings.Join(text, "")
	fText += "\n"

	warn := color.New(color.FgHiYellow)

	ct := time.Now()
	y := ct.Year()
	mo := checkMonth(ct.Month())
	d := checkDigits(int(ct.Day()))
	h := checkDigits(int(ct.Day()))
	mi := checkDigits(int(ct.Minute()))
	s := checkDigits(int(ct.Second()))
	warn.Printf(warnString, y, mo, d, h, mi, s, fText)	
}

func Error(text ...string) {
	fText := strings.Join(text, "")
	fText += "\n"

	err := color.New(color.FgHiRed)

	ct := time.Now()
	y := ct.Year()
	mo := checkMonth(ct.Month())
	d := checkDigits(int(ct.Day()))
	h := checkDigits(int(ct.Day()))
	mi := checkDigits(int(ct.Minute()))
	s := checkDigits(int(ct.Second()))
	err.Printf(errorString, y, mo, d, h, mi, s, fText)
}

func Fatal(text ...string) {
	fText := strings.Join(text, "")
	fText += "\n"

	fatal := color.New(color.FgRed)

	ct := time.Now()
	y := ct.Year()
	mo := checkMonth(ct.Month())
	d := checkDigits(int(ct.Day()))
	h := checkDigits(int(ct.Day()))
	mi := checkDigits(int(ct.Minute()))
	s := checkDigits(int(ct.Second()))
	fatal.Printf(fatalString, y, mo, d, h, mi, s, fText)
}