package llog

import (
	"fmt"
	lg "log"
	"os"
	"testing"
)

func testCriticalf(level int) {
	defer func() {
		if r := recover(); r != nil {
			if r != "It-is-CRITICALf\n" {
				T.Error("CRITICALf in level:", level)
			}
		}
	}()
	log.Criticalf("%s-%s-%s\n", "It", "is", "CRITICALf")
}

func testset(level int) {
	log.SetLevel(level)
	fmt.Printf("\n\nMessages in %d level\n", level)
	log.Debug("It", "is", "DEBUG")
	log.Debugf("%s-%s-%s\n", "It", "is", "DEBUGf")
	log.Info("It", "is", "INFO")
	log.Infof("%s-%s-%s\n", "It", "is", "INFOf")
	log.Warning("It", "is", "WARNING")
	log.Warningf("%s-%s-%s\n", "It", "is", "WARNINGf")
	log.Error("It", "is", "ERROR")
	log.Errorf("%s-%s-%s\n", "It", "is", "ERRORf")
	defer func() {
		if r := recover(); r != nil {
			if r != "It is CRITICAL\n" {
				T.Error("CRITICAL in level:", level)
			}
			testCriticalf(level)
		}
	}()
	log.Critical("It", "is", "CRITICAL")
}

var log *Logger

var T *testing.T

func TestAll(t *testing.T) {
	T = t
	log = New(os.Stderr, "", lg.Lshortfile|lg.Lmicroseconds)
	testset(DEBUG)
	testset(INFO)
	testset(WARNING)
	testset(ERROR)
	testset(CRITICAL)
}
