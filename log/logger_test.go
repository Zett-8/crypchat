package logger

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	os.Remove("log.log")
	os.Create("log.log")

	logfile := Init()
	defer logfile.Close()

	log.SetOutput(logfile)
	log.SetFlags(0)
	log.Println("this is test log output")

	bs, err := ioutil.ReadFile("log.log")
	if err != nil {
		t.Errorf("can not read log file")
	}

	if string(bs) != "this is test log output\n" {
		t.Errorf("log text is not correct")
	}

	os.Remove("log.log")
}
