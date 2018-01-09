package llog

import (
	"fmt"
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
	Criticalf("%s-%s-%s\n", "It", "is", "CRITICALf")
}

func testset(level int) {
	CurrntLevel = level
	fmt.Printf("\n\nMessages in %d level\n", level)
	Debug("It", "is", "DEBUG")
	Debugf("%s-%s-%s\n", "It", "is", "DEBUGf")
	Info("It", "is", "INFO")
	Infof("%s-%s-%s\n", "It", "is", "INFOf")
	Warning("It", "is", "WARNING")
	Warningf("%s-%s-%s\n", "It", "is", "WARNINGf")
	Error("It", "is", "ERROR")
	Errorf("%s-%s-%s\n", "It", "is", "ERRORf")
	defer func() {
		if r := recover(); r != nil {
			if r != "It is CRITICAL\n" {
				T.Error("CRITICAL in level:", level)
			}
			testCriticalf(level)
		}
	}()
	Critical("It", "is", "CRITICAL")
}


var T *testing.T

func TestAll(t *testing.T) {
	T = t
	testset(DEBUG)
	testset(INFO)
	testset(WARNING)
	testset(ERROR)
	testset(CRITICAL)
}

