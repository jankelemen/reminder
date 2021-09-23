// I didn't document any functions since I believe you can understand them from the tests
package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	LeftBracket      = "{"
	RightBracket     = "}"
	MsgFileNotFound  = "file with dates not found"
	MsgOnlyLinux     = "this program runs only on linux"
	NotifyCommandApp = "notify-send"
)

var (
	FilePath          string
	FilePath1         = os.Getenv("HOME") + "/.remind_me"
	FilePath2         = os.Getenv("HOME") + "/.config/reminder/remind_me"
	NotifyCommandArgs = []string{"-u", "critical"}
)

func init() {
	if runtime.GOOS != "linux" {
		log.Fatalln(MsgOnlyLinux)
	}

	if f, err := os.Stat(FilePath1); err == nil && !f.IsDir() {
		FilePath = FilePath1
	} else if f, err = os.Stat(FilePath2); err == nil && !f.IsDir() {
		FilePath = FilePath2
	} else {
		log.Fatalln(MsgFileNotFound)
	}
}

func main() {
	dat, err := os.ReadFile(FilePath)
	checkErr(err)

	for _, i := range strings.Split(string(dat), "\n") {
		if strings.Contains(i, ":") {
			split := strings.SplitN(i, ":", 2)
			times, text := UnpackAll(strings.TrimSpace(split[0])), strings.TrimSpace(split[1])

			cd := CurrentDate()
			for _, t := range times {
				if IsTheSameTime(cd, t) {
					cmd := append(NotifyCommandArgs, text)
					err = exec.Command(NotifyCommandApp, cmd...).Run()
					checkErr(err)
					break
				}
			}
		}
	}
}

func Unpack(toUnpack string) (res []string) {
	lPos := strings.Index(toUnpack, LeftBracket)
	rPos := strings.Index(toUnpack, RightBracket)
	replacements := strings.Split(toUnpack[lPos+1:rPos], ",")

	for _, replacement := range replacements {
		res = append(res, toUnpack[:lPos]+replacement+toUnpack[rPos+1:])
	}

	return
}

func UnpackAll(times string) (res []string) {
	if !strings.Contains(times, LeftBracket) {
		res = append(res, times)
		return res
	}

	res = Unpack(times)

	for strings.Contains(res[0], LeftBracket) {
		var temp []string
		for _, t := range res {
			temp = append(temp, Unpack(t)...)
		}
		res = temp
	}

	return res
}

func IsTheSameTime(time, timeReadFromFile string) bool {
	if len(time) != len(timeReadFromFile) {
		return false
	}

	for i := range time {
		if timeReadFromFile[i] != 'x' && time[i] != timeReadFromFile[i] {
			return false
		}
	}

	return true
}

func CurrentDate() string {
	cutFrom, cutTo := 2, 10
	return time.Now().String()[cutFrom:cutTo]
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
