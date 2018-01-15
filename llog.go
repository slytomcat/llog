// Package llog implements level-restricted logging.
// It is similar and based to standard log package but provides additional flexibility
// in management of logging messages.
//
// Package provides 5 levels of logging:
// 	DEBUG			- output all messages (lowest level)
// 	INFO			- output all messages except debug messages
// 	WARNING			- output all messages except debug and info messages - it is default level
// 	ERROR			- output only error and critical messages
// 	CRITICAL		- output only critical messages (highest level)
//
// There is a standard logger that initialized with os.Stderr as output, "" as prefix,
// log.LstdFlags (see https://golang.org/pkg/log/#pkg-constants for details) as flags, and
// WARNING as logging level.
//
// Logging level of standard logger can be changed via SetLevel(level int).
// Use constants DEBUG, INFO, WARNING, ERROR or CRITICAL for setting the level value.
// You can also change output: SetOutput(w io.Writer), log message prefix: SetPrefix(prefix string),
// and message format flags: SetFlags(flag int) (see https://golang.org/pkg/log/#pkg-constants
// for details about flags).
//
// To create the log message (via default logger) of the required logging level you have to
// use one of the following functions:
//  Debug(v ...interface{})			// equal to log.Println() when logging level is DEBUG
//  Info(v ...interface{}) 			// equal to log.Println() when logging level is INFO or less
//  Warning(v ...interface{})		// equal to log.Println() when logging level is WARNING or less
//  Error(v ...interface{})			// equal to log.Println() when logging level is ERROR or less
//  Critical(v ...interface{})		// equal to log.Panicln() in any logging level
//  Debugf(format string, v ...interface{})			// equal to log.Printf() when logging level is DEBUG
//  Infof(format string, v ...interface{})			// equal to log.Printf() when logging level is INFO or less
//  Warningf(format string, v ...interface{})		// equal to log.Printf() when logging level is WARNING or less
//  Errorf(format string, v ...interface{})			// equal to log.Printf() when logging level is ERROR or less
//  Criticalf(format string, v ...interface{})		// equal to log.Panicf() in any logging level
//
// Output of the functions will be additionally tagged with one letter identifier of logging level
// of message: D(EBUG), I(NFO), W(ARNING), E(RROR) or C(RITICAL). For example
//	Warning("Some message") // Output: W: Some message
//
// When the current logging level is greater than level of created message then function do nothing.
//
// New() allows to create the new Logger that has the same methods as standard logger.
//
// NOTE: as llog.Logger is just extension of log.Logger the full set of log package
// methods/functions/constants are available to use together with llog methods. But standard logger
// has only llog declared methods.
package llog

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	_ int = 10 * iota
	DEBUG        // 10
	INFO         // 20
	WARNING      // 30 default
	ERROR        // 40
	CRITICAL     // 50
)

type Logger struct {
	log.Logger		// extend the Logger of standard log package
	level int			// Current logging level
}

var std = New(os.Stderr, "", log.LstdFlags, WARNING)

// New returns new Logger.
// Parameters out, prefix and flag are the same parameters as in log.New() from log package
//  out - the io.Writer used as output for messages (usually it is os.Stderr)
//  prefix - log messages prefix
//  flags - time&info flags (see https://golang.org/pkg/log/#pkg-constants for details)
//  level - The logging level or -1 to set default level of logging (WARNING)
// It's better to use constants DEBUG, INFO, WARNING, ERROR or CRITICAL for setting the level
// instead of their numeric values.
func New(out io.Writer, prefix string, flag int, level int) *Logger {
	if level < 0 {
		level = WARNING
	}
	ret := Logger{
		*log.New(out, prefix, flag),
		level,
	}
	return &ret
}

// SetLevel sets logging level. Call of SetLevel is optional, if logging level is
// not set (either by Setup() or by SetLevel() then logging level is WARNING.
func (l *Logger) SetLevel(level int) {
	if level < 0 {
		level = WARNING
	}
	l.level = level
}

// Debug prints debug level message to output if current logging level is DEBUG.
// Debug is equal to log.Println()
func (l *Logger) Debug(v ...interface{}) {
	if l.level == DEBUG {
		l.Output(2, "D: "+fmt.Sprintln(v...))
	}
}

// Debugf prints DEBUG level message to output if current logging level is DEBUG.
// Debugf is equal to log.Printf()
func (l *Logger) Debugf(f string, v ...interface{}) {
	if l.level == DEBUG {
		l.Output(2, "D: "+fmt.Sprintf(f, v...))
	}
}

// Info prints INFO level message to output if current logging level is INFO or less.
// Info is equal to log.Println()
func (l *Logger) Info(v ...interface{}) {
	if l.level <= INFO {
		l.Output(2, "I: "+fmt.Sprintln(v...))
	}
}

// Infof prints INFO level message to output if current logging level is INFO or less.
// Infof is equal to log.Printf()
func (l *Logger) Infof(f string, v ...interface{}) {
	if l.level <= INFO {
		l.Output(2, "I: "+fmt.Sprintf(f, v...))
	}
}

// Warning prints WARNING level message to output if current logging level is WARNING or less
// Warning is equal to log.Println()
func (l *Logger) Warning(v ...interface{}) {
	if l.level <= WARNING {
		l.Output(2, "W: "+fmt.Sprintln(v...))
	}
}

// Warningf prints WARNING level message to output if current logging level is WARNING or less
// Warningf is equal to log.Printf()
func (l *Logger) Warningf(f string, v ...interface{}) {
	if l.level <= WARNING {
		l.Output(2, "W: "+fmt.Sprintf(f, v...))
	}
}

// Error prints ERROR level message to output if current logging level is ERROR or less
// Error is equal to log.Println()
func (l *Logger) Error(v ...interface{}) {
	if l.level <= ERROR {
		l.Output(2, "E: "+fmt.Sprintln(v...))
	}
}

// Errorf prints ERROR level message to output if current logging level is ERROR or less
// Errorf is equal to log.Printf()
func (l *Logger) Errorf(f string, v ...interface{}) {
	if l.level <= ERROR {
		l.Output(2, "E: "+fmt.Sprintf(f, v...))
	}
}

// Critical always prints CRITICAL level message to output and then call panic()
// Critical is equal to log.Panicln()
func (l *Logger) Critical(v ...interface{}) {
	s := "C: " + fmt.Sprintln(v...)
	l.Output(2, s)
	panic(s)
}

// Criticalf always prints CRITICAL level message to output and then call panic()
// Criticalf is equal to log.Panicf()
func (l *Logger) Criticalf(f string, v ...interface{}) {
	s := "C: " + fmt.Sprintf(f, v...)
	l.Output(2, s)
	panic(s)
}

// Debug prints DEBUG level message to output if current logging level is DEBUG
// Debug is equal to log.Println()
func Debug(v ...interface{}) {
	if std.level == DEBUG {
		std.Output(2, "D: "+fmt.Sprintln(v...))
	}
}

// Debugf prints DEBUG level message to output if current logging level is DEBUG
// Debugf is equal to log.Printf()
func Debugf(f string, v ...interface{}) {
	if std.level == DEBUG {
		std.Output(2, "D: "+fmt.Sprintf(f, v...))
	}
}

// Info prints INFO level message to output if current logging level is INFO or less
// Info is equal to log.Println()
func Info(v ...interface{}) {
	if std.level <= INFO {
		std.Output(2, "I: "+fmt.Sprintln(v...))
	}
}

// Infof prints INFO level message to output if current logging level is INFO or less
// Infof is equal to log.Printf()
func Infof(f string, v ...interface{}) {
	if std.level <= INFO {
		std.Output(2, "I: "+fmt.Sprintf(f, v...))
	}
}

// Warning prints WARNING level message to output if current logging level is WARNING or less
// Warning is equal to log.Println()
func Warning(v ...interface{}) {
	if std.level <= WARNING {
		std.Output(2, "W: "+fmt.Sprintln(v...))
	}
}

// Warningf prints WARNING level message to output if current logging level is WARNING or less
// Warningf is equal to log.Printf()
func Warningf(f string, v ...interface{}) {
	if std.level <= WARNING {
		std.Output(2, "W: "+fmt.Sprintf(f, v...))
	}
}

// Error prints ERROR level message to output if current logging level is ERROR or less
// Error is equal to log.Println()
func Error(v ...interface{}) {
	if std.level <= ERROR {
		std.Output(2, "E: "+fmt.Sprintln(v...))
	}
}

// Errorf prints ERROR level message to output if current logging level is ERROR or less
// Errorf is equal to log.Printf()
func Errorf(f string, v ...interface{}) {
	if std.level <= ERROR {
		std.Output(2, "E: "+fmt.Sprintf(f, v...))
	}
}

// Critical always prints CRITICAL level message to output and then call panic()
// Critical is equal to log.Panicln()
func Critical(v ...interface{}) {
	s := "C: " + fmt.Sprintln(v...)
	std.Output(2, s)
	panic(s)
}

// Criticalf always prints CRITICAL level message to output and then call panic()
// Criticalf is equal to log.Panicf()
func Criticalf(f string, v ...interface{}) {
	s := "C: " + fmt.Sprintf(f, v...)
	std.Output(2, s)
	panic(s)
}

// SetOutput sets the default logger output. The same method is applicable to Logger.
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// SetPrefix sets the messages prefix for default logger. The same method is applicable to Logger.
func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

// SetFlags sets messages flags for default logger. The same method is applicable to Logger.
func SetFlags(flag int) {
	std.SetFlags(flag)
}

// SetLevel sets the current logging level for default logger. The same method is applicable to Logger.
func SetLevel(level int) {
	std.SetLevel(level)
}
