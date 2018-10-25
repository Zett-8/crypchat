package logger

import (
	"os"
)

func Init() *os.File {
	logfile, err := os.OpenFile("log/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("cannnot open log.log:" + err.Error())
	}

	return logfile
}
