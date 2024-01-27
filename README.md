# llog
[![GithubGo](https://github.com/slytomcat/llog/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/slytomcat/llog/actions/workflows/go.yml)

Package llog implements level-restricted logging.
It is similar and based to standard log package but provides additional flexibility
in management of logging messages.

Package provides 5 levels of logging:

    DEBUG			- output all messages (lowest level)
    INFO			- output all messages except debug messages
    WARNING			- output all messages except debug and info messages - it is default level
    ERROR			- output only error and critical messages
    CRITICAL		- output only critical messages (highest level)

There is a standard logger that initialized with os.Stderr as output, "" as prefix,
log.LstdFlags (see https://golang.org/pkg/log/#pkg-constants for details) as flags, and
WARNING as logging level.

Logging level of standard logger can be changed via SetLevel(level int).
Use constants DEBUG, INFO, WARNING, ERROR or CRITICAL for setting the level value.
You can also change output: SetOutput(w io.Writer), log message prefix: SetPrefix(prefix string),
and message format flags: SetFlags(flag int) (see https://golang.org/pkg/log/#pkg-constants
for details about flags).

To create a new log message of the required logging level you have to use one of the following functions:

    Debug(v ...interface{})			// equal to log.Println() when logging level is DEBUG
    Info(v ...interface{}) 			// equal to log.Println() when logging level is INFO or less
    Warning(v ...interface{})		// equal to log.Println() when logging level is WARNING or less
    Error(v ...interface{})			// equal to log.Println() when logging level is ERROR or less
    Critical(v ...interface{})		// equal to log.Panicln() in any logging level
    Debugf(format string, v ...interface{})			// equal to log.Printf() when logging level is DEBUG
    Infof(format string, v ...interface{})			// equal to log.Printf() when logging level is INFO or less
    Warningf(format string, v ...interface{})		// equal to log.Printf() when logging level is WARNING or less
    Errorf(format string, v ...interface{})			// equal to log.Printf() when logging level is ERROR or less
    Criticalf(format string, v ...interface{})		// equal to log.Panicf() in any logging level

Output of the functions will be additionally tagged with one letter identifier of logging level
of the message: D(EBUG), I(NFO), W(ARNING), E(RROR) or C(RITICAL). For example:

    Warning("Some message") // Output: W: Some message

When the current logging level is greater than level of created message then function do nothing.

New() allows to create the new Logger that has the same methods as standard logger.

NOTE: as llog.Logger is just extension of log.Logger the full set of log package
methods/functions/constants are available to use together with llog methods. But standard logger
has only llog declared methods.
