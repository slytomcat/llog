// llog - leveled logging.
// It is similar to standard log package but provides more flexibility in logging levels.
// llog uses standard logger of log package as backbone.
//
// Package provides 5 levels of logging:
// 		DEBUG (level 10) - outputs all messages
// 		INFO (level 20) - outputs all messages except debug messages
// 		WARNING (level 30) - outputs all messages except debug and info messages - it is default level
// 		ERROR (level 40) - outputs only error and critical messages
// 		CRITICAL (level 50) - outputs only critical messages
//
// Logging format&level can be set via call to
//   Setup(out io.Writer, prefix string, flag int, level int)
// Parameters out, prefix and flag are the same parameters as in log.New():
//   out - the io.Writer used as output for messages (usually it is os.Stderr)
//   prefix - log messages prefix
//   flags - time$info flags (see https://golang.org/pkg/log/#pkg-constants for details)
//   level - The logging level or -1 to set default level of logging (WARNING)
// It's better to use constants DEBUG, INFO, WARNING, ERROR or CRITICAL for setting the level
// instead of their numeric values.
// Call of Setup - is optional, if it is not called, the default format of log package is used.
//
// Logging level can changed via SetLevel(level int). If logging level is not set (either by
// Setup or by SetLevel) then logging level is WARNING.
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
// When logging level is greater than level of created message then function do nothing (functions
// Debug*, Info*, Warning*, Error*, can be redefined as empty function but Critical always do its
// job)
//
// Note: Setup() configures the default logger of log package. You may use other log package
// functions to make log messages if it requered.

package llog

import (
	"fmt"
	"io"
	"log"
)

const (
	_ int = 10 * iota
	DEBUG
	INFO
	WARNING // default
	ERROR
	CRITICAL
)

var (
	defaultDebug = func(v ...interface{}) {
		log.Output(2, "D: "+fmt.Sprintln(v...))
	}
	defaultDebugf = func(f string, v ...interface{}) {
		log.Output(2, "D: "+fmt.Sprintf(f, v...))
	}
	defaultInfo = func(v ...interface{}) {
		log.Output(2, "I: "+fmt.Sprintln(v...))
	}
	defaultInfof = func(f string, v ...interface{}) {
		log.Output(2, "I: "+fmt.Sprintf(f, v...))
	}
	defaultWarning = func(v ...interface{}) {
		log.Output(2, "W: "+fmt.Sprintln(v...))
	}
	defaultWarningf = func(f string, v ...interface{}) {
		log.Output(2, "W: "+fmt.Sprintf(f, v...))
	}
	defaultError = func(v ...interface{}) {
		log.Output(2, "E: "+fmt.Sprintln(v...))
	}
	defaultErrorf = func(f string, v ...interface{}) {
		log.Output(2, "E: "+fmt.Sprintf(f, v...))
	}

	defaultNop  = func(v ...interface{}) {}
	defaultNopf = func(f string, v ...interface{}) {}

	// initial settings for default logging level WARNING
	Debug    = defaultNop;			Debugf   = defaultNopf
	Info     = defaultNop;			Infof    = defaultNopf
	Warning  = defaultWarning;	Warningf = defaultWarningf
	Error    = defaultError;		Errorf   = defaultErrorf
)

func Setup(out io.Writer, prefix string, flag int, level int) {
	log.SetOutput(out)
	log.SetPrefix(prefix)
	log.SetFlags(flag)
	if level < 0 {
		level = WARNING
	}
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
	s := "C: "+fmt.Sprintln(v...)
	log.Output(2, s)
	panic(s)
}

func Criticalf(f string, v ...interface{}) {
	s := "C: "+fmt.Sprintf(f, v...)
	log.Output(2, s)
	panic(s)
}
