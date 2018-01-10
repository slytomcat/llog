// llog - leveled logging.
// It is similar to standard log package but provides more flexibility in logging levels.
// llog uses standard log as backbone.
//
// Package provides 5 levels of logging:
// 		DEBUG (level 10) - outputs all messages
// 		INFO (level 20) - outputs all messages except debug messages
// 		WARNING (level 30) - outputs all messages except debug and info messages - it is default level
// 		ERROR (level 40) - outputs only error and critical messages
// 		CRITICAL (level 50) - outputs only critical messages
//
// Logging format&level can be set via Init(out io.Writer, prefix string, flag int, level int)
// Parameters out, prefix and flag are the same parameters as in log.New()
// It's better to use constants DEBUG, INFO, WARNING, ERROR or CRITICAL instead of their numeric
// value.
// Logging level can changed via SetLevel(level int).
//
// To create the log message you have to use one of the following methods:
//	Debug(v ...interface{}) - equal to log.Println() when logging level is DEBUG
//	Debugf(format string, v ...interface{}) - equal to log.Printf() when logging level is DEBUG
//	Info(v ...interface{}) - equal to log.Println() when logging level is DEBUG or INFO
//	Infof(format string, v ...interface{}) - equal to log.Printf() when logging level is DEBUG or INFO
//	Warning(v ...interface{}) - equal to log.Println() when logging level is DEBUG, INFO or WARNING
//	Warningf(format string, v ...interface{}) - equal to log.Printf() when logging level is DEBUG, INFO or WARNING
//	Error(v ...interface{}) - equal to log.Println() when logging level is DEBUG, INFO, WARNING or ERROR
//	Errorf(format string, v ...interface{}) - equal to log.Printf() when logging level is DEBUG, INFO, WARNING or ERROR
//	Critical(v ...interface{}) - equal to log.Panicln() in any logging level
//	Criticalf(format string, v ...interface{}) - equal to log.Panicf() in any logging level
//
// Note: Init configures the default logger of log lib. You may use other log functions to make
// log messages.

package llog

import (
	"log"
	"io"
)

const (
	_ int = 10 * iota
	DEBUG
	INFO
	WARNING // default
	ERROR
	CRITICAL
)

var(
defaultDebug = func(v ...interface{}) {
	log.Println(append([]interface{}{"D:"}, v...)...)
}
defaultDebugf = func(f string, v ...interface{}) {
	log.Printf("%s "+f, append([]interface{}{"D:"}, v...)...)
}
defaultInfo = func(v ...interface{}) {
	log.Println(append([]interface{}{"I:"}, v...)...)
}
defaultInfof = func(f string, v ...interface{}) {
	log.Printf("%s "+f, append([]interface{}{"I:"}, v...)...)
}
defaultWarning = func(v ...interface{}) {
	log.Println(append([]interface{}{"W:"}, v...)...)
}
defaultWarningf = func(f string, v ...interface{}) {
	log.Printf("%s "+f, append([]interface{}{"W:"}, v...)...)
}
defaultError = func(v ...interface{}) {
	log.Println(append([]interface{}{"E:"}, v...)...)
}
defaultErrorf = func(f string, v ...interface{}) {
	log.Printf("%s "+f, append([]interface{}{"E:"}, v...)...)
}

defaultNop = func(v ...interface{}) {}
defaultNopf = func(f string, v ...interface{}) {}

// initial settings for default logging level WARNING
Debug = defaultNop
Debugf = defaultNopf
Info = defaultNop
Infof = defaultNopf
Warning = defaultWarning
Warningf = defaultWarningf
Error = defaultError
Errorf = defaultErrorf
)

func Init(out io.Writer, prefix string, flag int, level int) {
	log.SetOutput(out)
	log.SetPrefix(prefix)
	log.SetFlags(flag)
	SetLevel(level)
}

func SetLevel(level int) {
	if level > DEBUG {
		Debug = defaultNop
		Debugf = defaultNopf
		if level > INFO {
			Info = defaultNop
			Infof = defaultNopf
			if level > WARNING {
				Warning = defaultNop
				Warningf = defaultNopf
				if level > ERROR {
					Error = defaultNop
					Errorf = defaultNopf
					return
				}
			}
		}
	}
	// here level <= ERROR
	Error = defaultError
	Errorf = defaultErrorf
	if level <= WARNING {
		Warning = defaultWarning
		Warningf = defaultWarningf
		if level <= INFO {
			Info = defaultInfo
			Infof = defaultInfof
			if level <= DEBUG {
				Debug = defaultDebug
				Debugf = defaultDebugf
			}
		}
	}
}


func Critical(v ...interface{}) {
	log.Panicln(append([]interface{}{"C:"}, v...)...)
}

func Criticalf(f string, v ...interface{}) {
	log.Panicf("%s "+f, append([]interface{}{"C:"}, v...)...)
}
