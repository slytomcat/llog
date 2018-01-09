// llog - leveled logging.
// It is similar to standard log package but provides more flexibility in logging levels.
//
// Package provides 5 levels of logging:
// 		DEBUG (level 10) - outputs all messages
// 		INFO (level 20) - outputs all messages except debug messages
// 		WARNING (level 30) - outputs all messages except debug and info messages
// 		ERROR (level 40) - outputs only error and critical messages
// 		CRITICAL (level 50) - outputs only critical messages
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
//	Critical(v ...interface{}) - equal to log.Fatalln() in any logging level
//	Criticalf(format string, v ...interface{}) - equal to log.Fatalf() in any logging level
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

var CurrntLevel = 30

func Debug(v ...interface{}) {
	if CurrntLevel >= DEBUG {
		log.Println(append([]interface{}{"D:"}, v...)...)
	}
}

func Debugf(f string, v ...interface{}) {
	if CurrntLevel >= DEBUG {
		log.Printf("%s "+f, append([]interface{}{"D:"}, v...)...)
	}
}

func Info(v ...interface{}) {
	if CurrntLevel >= INFO {
		log.Println(append([]interface{}{"I:"}, v...)...)
	}
}

func Infof(f string, v ...interface{}) {
	if CurrntLevel >= INFO {
		log.Printf("%s "+f, append([]interface{}{"I:"}, v...)...)
	}
}

func Warning(v ...interface{}) {
	if CurrntLevel >= WARNING {
		log.Println(append([]interface{}{"W:"}, v...)...)
	}
}

func Warningf(f string, v ...interface{}) {
	if CurrntLevel >= INFO {
		log.Printf("%s "+f, append([]interface{}{"W:"}, v...)...)
	}
}

func Error(v ...interface{}) {
	if CurrntLevel >= WARNING {
		log.Println(append([]interface{}{"E:"}, v...)...)
	}
}

func Errorf(f string, v ...interface{}) {
	if CurrntLevel >= INFO {
		log.Printf("%s "+f, append([]interface{}{"E:"}, v...)...)
	}
}

func Critical(v ...interface{}) {
	log.Fatalln(append([]interface{}{"C:"}, v...)...)
}

func Criticalf(f string, v ...interface{}) {
	log.Fatalf("%s "+f, append([]interface{}{"C:"}, v...)...)
}

