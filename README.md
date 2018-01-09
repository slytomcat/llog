# llog
Leveled logging for GO

It is similar to standard log package but provides more flexibility in logging levels.

Package provides 5 levels of logging:
  - DEBUG (level 10) - outputs all messages
  - INFO (level 20) - outputs all messages except debug messages
  - WARNING (level 30) - outputs all messages except debug and info messages
  - ERROR (level 40) - outputs only error and critical messages
  - CRITICAL (level 50) - outputs only critical messages

New logger is created by call New() that accepts the same parameters as log.New()

By default the logger created with WARNING logging level. You can change logging level
by call of Logger.SetLevel(level int). You may use constants DEBUG, INFO, WARNING, ERROR
and CRITICAL instead of digital level values.

To create the log message you have to use one of the following methods:

    Logger.Debug(v ...interface{}) - equal to log.Println() when logging level is DEBUG
    Logger.Debugf(format string, v ...interface{}) - equal to log.Printf() when logging level is DEBUG
    Logger.Info(v ...interface{}) - equal to log.Println() when logging level is DEBUG or INFO
    Logger.Infof(format string, v ...interface{}) - equal to log.Printf() when logging level is DEBUG or INFO
    Logger.Warning(v ...interface{}) - equal to log.Println() when logging level is DEBUG, INFO or WARNING
    Logger.Warningf(format string, v ...interface{}) - equal to log.Printf() when logging level is DEBUG, INFO or WARNING
    Logger.Error(v ...interface{}) - equal to log.Println() when logging level is DEBUG, INFO, WARNING or ERROR
    Logger.Errorf(format string, v ...interface{}) - equal to log.Printf() when logging level is DEBUG, INFO, WARNING or ERROR
    Logger.Critical(v ...interface{}) - equal to log.Println() + panic() in any logging level
    Logger.Criticalf(format string, v ...interface{}) - equal to log.Printf() + panic() in any logging level



