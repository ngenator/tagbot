package lumber

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//Log levels
const (
	TRACE = iota
	DEBUG
	INFO
	QUIET
	SILENT
)

//Log Simple log struct to wrap for colors
type Log struct {
	*log.Logger
	colorCode string
}

//Log.println Mimic the internal println function of the log module so we can wrap in colors
func (l *Log) println(v ...interface{}) {
	s := fmt.Sprintf("\033[1;%v%v\033[0m\n", l.colorCode, fmt.Sprint(v...))
	l.Logger.Output(2, s)
}

//newLog Create a new simple wrapped log object
func newLog(out io.Writer, prefix string, flag int, colorCode string) *Log {
	return &Log{
		log.New(out, prefix, flag),
		colorCode,
	}
}

//Logger The logger object containing each of the wrapped log objects
type Logger struct {
	TraceLog   Log
	DebugLog   Log
	SuccessLog Log
	InfoLog    Log
	WarningLog Log
	ErrorLog   Log
}

//Trace Forward to the trace logger
func (l *Logger) Trace(v ...interface{}) {
	l.TraceLog.println(v)
}

//Debug Forward to the debug logger
func (l *Logger) Debug(v ...interface{}) {
	l.DebugLog.println(v)
}

//Success Forward to the success logger
func (l *Logger) Success(v ...interface{}) {
	l.SuccessLog.println(v)
}

//Info Forward to the info logger
func (l *Logger) Info(v ...interface{}) {
	l.InfoLog.println(v)
}

//Warning ..
// Forward to the warning logger
func (l *Logger) Warning(v ...interface{}) {
	l.WarningLog.println(v)
}

//Error ..
// Error Forward to the error logger
func (l *Logger) Error(v ...interface{}) {
	l.ErrorLog.println(v)
}

//NewLogger ..
// Create a new logger object and direct outputs to various io.Writers
// Levels are TRACE, DEBUG, INFO, QUIET, SILENT
func NewLogger(loglevel int) *Logger {

	var (
		traceLog   *Log
		debugLog   *Log
		successLog *Log
		infoLog    *Log
		warningLog *Log
		errorLog   *Log
	)

	var (
		traceHandle   io.Writer
		debugHandle   io.Writer
		successHandle io.Writer
		infoHandle    io.Writer
		warningHandle io.Writer
		errorHandle   io.Writer
	)

	switch loglevel {
	case TRACE:
		traceHandle = os.Stdout
		debugHandle = os.Stdout
		successHandle = os.Stdout
		infoHandle = os.Stdout
		warningHandle = os.Stdout
		errorHandle = os.Stderr
	case DEBUG:
		traceHandle = ioutil.Discard
		debugHandle = os.Stdout
		successHandle = os.Stdout
		infoHandle = os.Stdout
		warningHandle = os.Stdout
		errorHandle = os.Stderr
	case INFO:
		traceHandle = ioutil.Discard
		debugHandle = ioutil.Discard
		successHandle = os.Stdout
		infoHandle = os.Stdout
		warningHandle = os.Stdout
		errorHandle = os.Stderr
	case QUIET:
		traceHandle = ioutil.Discard
		debugHandle = ioutil.Discard
		successHandle = os.Stdout
		infoHandle = os.Stdout
		warningHandle = ioutil.Discard
		errorHandle = os.Stderr
	case SILENT:
		traceHandle = ioutil.Discard
		debugHandle = ioutil.Discard
		successHandle = ioutil.Discard
		infoHandle = ioutil.Discard
		warningHandle = ioutil.Discard
		errorHandle = ioutil.Discard
	}

	traceLog = newLog(traceHandle,
		"TRACE: ",
		log.Ltime|log.Lshortfile,
		"32m")

	debugLog = newLog(debugHandle,
		"DEBUG: ",
		log.Ltime,
		"35m")

	successLog = newLog(successHandle,
		"SUCCESS: ",
		log.Ltime,
		"32m")

	infoLog = newLog(infoHandle,
		"INFO: ",
		log.Ltime,
		"36m")

	warningLog = newLog(warningHandle,
		"WARNING: ",
		log.Ltime,
		"31m")

	errorLog = newLog(errorHandle,
		"ERROR: ",
		log.Ltime|log.Lshortfile,
		"31m")

	return &Logger{
		*traceLog,
		*debugLog,
		*successLog,
		*infoLog,
		*warningLog,
		*errorLog,
	}
}
