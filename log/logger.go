package logger

import (
	"os"

	"github.com/rhysd/abspath"
)

func Init() *os.File {

	home, err := abspath.HomeDir()
	if err != nil {
		panic(err)
	}

	logfile, err := os.OpenFile(home.String()+"/go/src/github.com/Zett-8/crypchat/log/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("cannnot open log.log:" + err.Error())
	}

	return logfile
}
