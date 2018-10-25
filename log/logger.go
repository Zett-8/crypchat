package logger

import (
	"io"
	"log"
	"os"
)

func Init() {
	logfile, err := os.OpenFile("log/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("cannnot open log.log:" + err.Error())
	}
	defer logfile.Close()

	// io.MultiWriteで、
	// 標準出力とファイルの両方を束ねて、
	// logの出力先に設定する
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))

	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("last test")
}
