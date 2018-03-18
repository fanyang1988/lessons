package logger

//
// logger Just an seelog help funcs
//

import (
	"github.com/cihub/seelog"
)

// Interface interface for logger
type Interface interface {
	Tracef(format string, params ...interface{})
	Debugf(format string, params ...interface{})
	Infof(format string, params ...interface{})
	Warnf(format string, params ...interface{}) error
	Errorf(format string, params ...interface{}) error
	Criticalf(format string, params ...interface{}) error

	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{}) error
	Error(v ...interface{}) error
	Critical(v ...interface{}) error
}

// NullLogger an null logger for not log anything
var NullLogger Interface = new(nullLogger)

// NullLogger an null logger do nothing
type nullLogger struct {
}

// funcs by do nothing
func (n *nullLogger) Tracef(format string, params ...interface{}) {
}

func (n *nullLogger) Debugf(format string, params ...interface{}) {
}

func (n *nullLogger) Infof(format string, params ...interface{}) {
}

func (n *nullLogger) Warnf(format string, params ...interface{}) error {
	return nil
}

func (n *nullLogger) Errorf(format string, params ...interface{}) error {
	return nil
}

func (n *nullLogger) Criticalf(format string, params ...interface{}) error {
	return nil
}

func (n *nullLogger) Trace(v ...interface{}) {
}

func (n *nullLogger) Debug(v ...interface{}) {
}

func (n *nullLogger) Info(v ...interface{}) {
}

func (n *nullLogger) Warn(v ...interface{}) error {
	return nil
}

func (n *nullLogger) Error(v ...interface{}) error {
	return nil
}

func (n *nullLogger) Critical(v ...interface{}) error {
	return nil
}

// InitSeelogCfg change seelog cfg
func InitSeelogCfg(path string, logFileName string) {
	logger, err := seelog.LoggerFromConfigAsBytes(
		[]byte(makeLoggerCfg(path, logFileName)))
	if err != nil {
		// just panic
		panic(err)
	}
	seelog.ReplaceLogger(logger)
}

// SeeLogLogger get logger imp by seelog
func SeeLogLogger() seelog.LoggerInterface {
	return seelog.Current
}
