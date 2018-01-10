// llog - leveled logging.
// It is similar to standard log package but provides more flexibility in logging levels.
// llog uses standard log as backbone.
//
// Package provides 5 levels of logging:
// 		DEBUG (level 10) - outputs all messages
// 		INFO (level 20) - outputs all messages except debug messages
// 		WARNING (level 30) - outputs all messages except debug and info messages
// 		ERROR (level 40) - outputs only error and critical messages
// 		CRITICAL (level 50) - outputs only critical messages
//
// Loggging level must be set ONLY ONCE (thre is no way to decrease logging level).
// Logging level is set via SetLevel(level int). It's better to use constatns DEBUG, INFO,
// WARNING, ERROR or CRITICAL instead of their numeric value.
//
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

package llog

import (
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

func SetLevel(level int) {
	if level > DEBUG {
		Debug = func(v ...interface{}) {}
		Debugf = func(f string, v ...interface{}) {}
	}
	if level > INFO {
		Info = Debug
		Infof = Debugf
	}
	if level > WARNING {
		Warning = Debug
		Warningf = Debugf
	}
	if level > ERROR {
		Error = Debug
		Errorf = Debugf
	}
}

var Debug = func(v ...interface{}) {
	log.Println(append([]interface{}{"D:"}, v...)...)
}

var Debugf = func(f string, v ...interface{}) {
	log.Printf("%s "+f, append([]interface{}{"D:"}, v...)...)
}

var Info = func(v ...interface{}) {
	log.Println(append([]interface{}{"I:"}, v...)...)
}

var Infof = func(f string, v ...interface{}) {
	log.Printf("%s "+f, append([]interface{}{"I:"}, v...)...)
}

var Warning = func(v ...interface{}) {
	log.Println(append([]interface{}{"W:"}, v...)...)
}

var Warningf = func(f string, v ...interface{}) {
	log.Printf("%s "+f, append([]interface{}{"W:"}, v...)...)
}

var Error = func(v ...interface{}) {
	log.Println(append([]interface{}{"E:"}, v...)...)
}

var Errorf = func(f string, v ...interface{}) {
	log.Printf("%s "+f, append([]interface{}{"E:"}, v...)...)
}

func Critical(v ...interface{}) {
	log.Panicln(append([]interface{}{"C:"}, v...)...)
}

func Criticalf(f string, v ...interface{}) {
	log.Panicf("%s "+f, append([]interface{}{"C:"}, v...)...)
}
