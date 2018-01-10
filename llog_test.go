package llog

import (
	"log"
	"os"
	"testing"
)

func testCriticalf() {  // NOTE: # of lines importat for predictable output
	defer func() {
		_ = recover()
	}()
	Criticalf("%s-%s-%s\n", "It", "is", "CRITICALf")
}

func testset(level int) {  // NOTE:  # of lines importat for predictable output
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
	Setup(os.Stdout, "", log.Lshortfile, -1)
	os.Exit(m.Run())
}

func ExampleDebug() {
	log.SetOutput(os.Stdout)
	testset(DEBUG)
	// Output:
	// llog_test.go:18: D: It is DEBUG
	// llog_test.go:19: D: It-is-DEBUGf
	// llog_test.go:20: I: It is INFO
	// llog_test.go:21: I: It-is-INFOf
	// llog_test.go:22: W: It is WARNING
	// llog_test.go:23: W: It-is-WARNINGf
	// llog_test.go:24: E: It is ERROR
	// llog_test.go:25: E: It-is-ERRORf
	// llog_test.go:30: C: It is CRITICAL
	// llog_test.go:13: C: It-is-CRITICALf
}

func ExampleInfo() {
	log.SetOutput(os.Stdout)
	testset(INFO)
	// Output:
	// llog_test.go:20: I: It is INFO
	// llog_test.go:21: I: It-is-INFOf
	// llog_test.go:22: W: It is WARNING
	// llog_test.go:23: W: It-is-WARNINGf
	// llog_test.go:24: E: It is ERROR
	// llog_test.go:25: E: It-is-ERRORf
	// llog_test.go:30: C: It is CRITICAL
	// llog_test.go:13: C: It-is-CRITICALf
}

func ExampleWarning() {
	log.SetOutput(os.Stdout)
	testset(WARNING)
	// Output:
	// llog_test.go:22: W: It is WARNING
	// llog_test.go:23: W: It-is-WARNINGf
	// llog_test.go:24: E: It is ERROR
	// llog_test.go:25: E: It-is-ERRORf
	// llog_test.go:30: C: It is CRITICAL
	// llog_test.go:13: C: It-is-CRITICALf
}

func ExampleError() {
	log.SetOutput(os.Stdout)
	testset(ERROR)
	// Output:
	// llog_test.go:24: E: It is ERROR
	// llog_test.go:25: E: It-is-ERRORf
	// llog_test.go:30: C: It is CRITICAL
	// llog_test.go:13: C: It-is-CRITICALf
}

func ExampleCritical() {
	log.SetOutput(os.Stdout)
	testset(CRITICAL)
	// Output:
	// llog_test.go:30: C: It is CRITICAL
	// llog_test.go:13: C: It-is-CRITICALf
}

func ExampleError1() {
	log.SetOutput(os.Stdout)
	testset(ERROR)
	// Output:
	// llog_test.go:24: E: It is ERROR
	// llog_test.go:25: E: It-is-ERRORf
	// llog_test.go:30: C: It is CRITICAL
	// llog_test.go:13: C: It-is-CRITICALf
}

func ExampleWarning1() {
	log.SetOutput(os.Stdout)
	testset(WARNING)
	// Output:
	// llog_test.go:22: W: It is WARNING
	// llog_test.go:23: W: It-is-WARNINGf
	// llog_test.go:24: E: It is ERROR
	// llog_test.go:25: E: It-is-ERRORf
	// llog_test.go:30: C: It is CRITICAL
	// llog_test.go:13: C: It-is-CRITICALf
}

func ExampleInfo1() {
	log.SetOutput(os.Stdout)
	testset(INFO)
	// Output:
	// llog_test.go:20: I: It is INFO
	// llog_test.go:21: I: It-is-INFOf
	// llog_test.go:22: W: It is WARNING
	// llog_test.go:23: W: It-is-WARNINGf
	// llog_test.go:24: E: It is ERROR
	// llog_test.go:25: E: It-is-ERRORf
	// llog_test.go:30: C: It is CRITICAL
	// llog_test.go:13: C: It-is-CRITICALf
}

func ExampleDebug1() {
	log.SetOutput(os.Stdout)
	testset(DEBUG)
	// Output:
	// llog_test.go:18: D: It is DEBUG
	// llog_test.go:19: D: It-is-DEBUGf
	// llog_test.go:20: I: It is INFO
	// llog_test.go:21: I: It-is-INFOf
	// llog_test.go:22: W: It is WARNING
	// llog_test.go:23: W: It-is-WARNINGf
	// llog_test.go:24: E: It is ERROR
	// llog_test.go:25: E: It-is-ERRORf
	// llog_test.go:30: C: It is CRITICAL
	// llog_test.go:13: C: It-is-CRITICALf
}
