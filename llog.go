package llog

import (
	"fmt"
	"io"
	lg "log"
)

const (
	_ int = 10 * iota
	DEBUG
	INFO
	WARNING // default
	ERROR
	CRITICAL
)

type Logger struct {
	level    int
	logger   *lg.Logger
	Debug    func(_ ...interface{})
	Debugf   func(_ string, _ ...interface{})
	Info     func(_ ...interface{})
	Infof    func(_ string, _ ...interface{})
	Warning  func(_ ...interface{})
	Warningf func(_ string, _ ...interface{})
	Error    func(_ ...interface{})
	Errorf   func(_ string, _ ...interface{})
}

func (l Logger) Critical(v ...interface{}) {
	l.logger.Println(append([]interface{}{"C:"}, v...)...)
	panic(fmt.Sprintln(v...))
}

func (l Logger) Criticalf(f string, v ...interface{}) {
	l.logger.Printf("%s "+f, append([]interface{}{"C:"}, v...)...)
	panic(fmt.Sprintf(f, v...))
}

func New(out io.Writer, prefix string, flag int) *Logger {
	ret := &Logger{WARNING, lg.New(out, prefix, flag), nil, nil, nil, nil, nil, nil, nil, nil}
	ret.SetLevel(WARNING)
	return ret
}

func (l *Logger) SetLevel(level int) {
	l.level = level
	none := func(_ ...interface{}) {}
	nonef := func(_ string, _ ...interface{}) {}
	l.Debug = func(v ...interface{}) { l.logger.Println(append([]interface{}{"D:"}, v...)...) }
	l.Debugf = func(f string, v ...interface{}) { l.logger.Printf("%s "+f, append([]interface{}{"D:"}, v...)...) }
	l.Info = func(v ...interface{}) { l.logger.Println(append([]interface{}{"I:"}, v...)...) }
	l.Infof = func(f string, v ...interface{}) { l.logger.Printf("%s "+f, append([]interface{}{"I:"}, v...)...) }
	l.Warning = func(v ...interface{}) { l.logger.Println(append([]interface{}{"W:"}, v...)...) }
	l.Warningf = func(f string, v ...interface{}) { l.logger.Printf("%s "+f, append([]interface{}{"W:"}, v...)...) }
	l.Error = func(v ...interface{}) { l.logger.Println(append([]interface{}{"E:"}, v...)...) }
	l.Errorf = func(f string, v ...interface{}) { l.logger.Printf("%s "+f, append([]interface{}{"E:"}, v...)...) }
	if level >= INFO { // INFO, WARNING, ERROR, CRITICAL
		l.Debug = none
		l.Debugf = nonef
	}
	if level >= WARNING { // WARNING, ERROR, CRITICAL
		l.Info = none
		l.Infof = nonef
	}
	if level >= ERROR { // ERROR, CRITICAL
		l.Warning = none
		l.Warningf = nonef
	}
	if level >= CRITICAL { // CRITICAL
		l.Error = none
		l.Errorf = nonef
	}
}
