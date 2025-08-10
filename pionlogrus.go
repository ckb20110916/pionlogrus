package pionlogrus

import (
	"github.com/ckb20110916/loguselogrus"
	"github.com/pion/logging"
	"github.com/sirupsen/logrus"
	"sync"
)

type logger struct {
	entry *logrus.Entry
}

func (slf *logger) Trace(msg string) {
	slf.entry.Trace(msg)
}

func (slf *logger) Tracef(format string, args ...interface{}) {
	slf.entry.Tracef(format, args...)
}

func (slf *logger) Debug(msg string) {
	slf.entry.Debug(msg)
}

func (slf *logger) Debugf(format string, args ...interface{}) {
	slf.entry.Debugf(format, args...)
}

func (slf *logger) Info(msg string) {
	slf.entry.Info(msg)
}

func (slf *logger) Infof(format string, args ...interface{}) {
	slf.entry.Infof(format, args...)
}

func (slf *logger) Warn(msg string) {
	slf.entry.Warn(msg)
}

func (slf *logger) Warnf(format string, args ...interface{}) {
	slf.entry.Warnf(format, args...)
}

func (slf *logger) Error(msg string) {
	slf.entry.Error(msg)
}

func (slf *logger) Errorf(format string, args ...interface{}) {
	slf.entry.Errorf(format, args...)
}

var (
	Factory = &factory{}
)

type factory struct {
	mutex   sync.Mutex
	loggers []*logger
}

func (f *factory) NewLogger(scope string) logging.LeveledLogger {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	l := &logger{
		entry: loguselogrus.Logger.WithField("scope", scope),
	}
	f.loggers = append(f.loggers, l)
	return l
}
