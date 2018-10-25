package logger

import (
	"os"
)

func Init(f string) *os.File {
	logfile, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("cannnot open log.log:" + err.Error())
	}

	return logfile
}
