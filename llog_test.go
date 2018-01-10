package llog

import (
	"log"
	"os"
	"testing"
)

func testCriticalf() {
	defer func() {
		_ = recover()
	}()
	Criticalf("%s-%s-%s\n", "It", "is", "CRITICALf")
}

func testset(level int) {
	SetLevel(level)
	Debug("It", "is", "DEBUG")
	Debugf("%s-%s-%s\n", "It", "is", "DEBUGf")
	Info("It", "is", "INFO")
	Infof("%s-%s-%s\n", "It", "is", "INFOf")
	Warning("It", "is", "WARNING")
	Warningf("%s-%s-%s\n", "It", "is", "WARNINGf")
	Error("It", "is", "ERROR")
	Errorf("%s-%s-%s\n", "It", "is", "ERRORf")
	defer func() {
		_ = recover()
		testCriticalf()
	}()
	Critical("It", "is", "CRITICAL")
}

func TestMain(m *testing.M) {
	Init(os.Stdout, "", 0, DEBUG)
	os.Exit(m.Run())
}

func ExampleDebug() {
	log.SetOutput(os.Stdout)
	testset(DEBUG)
	// Output:
	// D: It is DEBUG
	// D: It-is-DEBUGf
	// I: It is INFO
	// I: It-is-INFOf
	// W: It is WARNING
	// W: It-is-WARNINGf
	// E: It is ERROR
	// E: It-is-ERRORf
	// C: It is CRITICAL
	// C: It-is-CRITICALf
}

func ExampleInfo() {
	log.SetOutput(os.Stdout)
	testset(INFO)
	// Output:
	// I: It is INFO
	// I: It-is-INFOf
	// W: It is WARNING
	// W: It-is-WARNINGf
	// E: It is ERROR
	// E: It-is-ERRORf
	// C: It is CRITICAL
	// C: It-is-CRITICALf
}

func ExampleWarning() {
	log.SetOutput(os.Stdout)
	testset(WARNING)
	// Output:
	// W: It is WARNING
	// W: It-is-WARNINGf
	// E: It is ERROR
	// E: It-is-ERRORf
	// C: It is CRITICAL
	// C: It-is-CRITICALf
}

func ExampleError() {
	log.SetOutput(os.Stdout)
	testset(ERROR)
	// Output:
	// E: It is ERROR
	// E: It-is-ERRORf
	// C: It is CRITICAL
	// C: It-is-CRITICALf
}

func ExampleCritical() {
	log.SetOutput(os.Stdout)
	testset(CRITICAL)
	// Output:
	// C: It is CRITICAL
	// C: It-is-CRITICALf
}

func ExampleError1() {
	log.SetOutput(os.Stdout)
	testset(ERROR)
	// Output:
	// E: It is ERROR
	// E: It-is-ERRORf
	// C: It is CRITICAL
	// C: It-is-CRITICALf
}

func ExampleWarning1() {
	log.SetOutput(os.Stdout)
	testset(WARNING)
	// Output:
	// W: It is WARNING
	// W: It-is-WARNINGf
	// E: It is ERROR
	// E: It-is-ERRORf
	// C: It is CRITICAL
	// C: It-is-CRITICALf
}

func ExampleInfo1() {
	log.SetOutput(os.Stdout)
	testset(INFO)
	// Output:
	// I: It is INFO
	// I: It-is-INFOf
	// W: It is WARNING
	// W: It-is-WARNINGf
	// E: It is ERROR
	// E: It-is-ERRORf
	// C: It is CRITICAL
	// C: It-is-CRITICALf
}

func ExampleDebug1() {
	log.SetOutput(os.Stdout)
	testset(DEBUG)
	// Output:
	// D: It is DEBUG
	// D: It-is-DEBUGf
	// I: It is INFO
	// I: It-is-INFOf
	// W: It is WARNING
	// W: It-is-WARNINGf
	// E: It is ERROR
	// E: It-is-ERRORf
	// C: It is CRITICAL
	// C: It-is-CRITICALf
}
