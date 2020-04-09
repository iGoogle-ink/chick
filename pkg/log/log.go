package log

import (
	"os"
	"sync"

	"github.com/op/go-logging"
)

var (
	log     *Logger
	lPrefix string
)

type Level int

const (
	DEBUG Level = iota + 1
	INFO
	NOTICE
	WARNING
	ERROR
	CRITICAL
	FATAL
	PANIC
)

type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

type Logger struct {
	logger  *logging.Logger
	format  logging.Formatter
	backend *logging.LogBackend
	once    sync.Once
}

func (l *Logger) setPrefix(prefix string) {
	l.backend = logging.NewLogBackend(os.Stdout, prefix, 0)
}

func (l *Logger) logOut(lv Level, format *string, args ...interface{}) {
	l.once.Do(func() {
		initLogger()
	})
	switch lv {
	case DEBUG:
		if format != nil {
			log.logger.Debugf(*format, args...)
			return
		}
		log.logger.Debug(args...)
	case INFO:
		if format != nil {
			log.logger.Infof(*format, args...)
			return
		}
		log.logger.Info(args...)
	case NOTICE:
		if format != nil {
			log.logger.Noticef(*format, args...)
			return
		}
		log.logger.Notice(args...)
	case WARNING:
		if format != nil {
			log.logger.Warningf(*format, args...)
			return
		}
		log.logger.Warning(args...)
	case ERROR:
		if format != nil {
			log.logger.Errorf(*format, args...)
			return
		}
		log.logger.Error(args...)
	case CRITICAL:
		if format != nil {
			log.logger.Criticalf(*format, args...)
			return
		}
		log.logger.Critical(args...)
	case FATAL:
		if format != nil {
			log.logger.Fatalf(*format, args...)
			return
		}
		log.logger.Fatal(args...)
	case PANIC:
		if format != nil {
			log.logger.Panicf(*format, args...)
			return
		}
		log.logger.Panic(args...)
	default:
		if format != nil {
			log.logger.Infof(*format, args...)
			return
		}
		log.logger.Info(args...)
	}
}

func SetPrefix(prefix string) {
	lPrefix = prefix
	if log != nil {
		log.setPrefix(prefix)
	}
}

func init() {
	log = &Logger{}
}

func initLogger() {
	log.logger = logging.MustGetLogger("chick")
	log.logger.ExtraCalldepth = 2
	// %{id:03x} %{level:-8s}
	log.format = logging.MustStringFormatter(`%{color}[%{time:2006-01-02 15:04:05.000}] ▶ %{level:.4s} [%{shortfunc}] ▶ %{color:reset}%{message}`)
	log.setPrefix(lPrefix)
	backendFormatter := logging.NewBackendFormatter(log.backend, log.format)
	logging.SetBackend(backendFormatter)
}

func Debug(args ...interface{}) {
	log.logOut(DEBUG, nil, args...)
}

func Debugf(format string, args ...interface{}) {
	log.logOut(DEBUG, &format, args...)
}

func Info(args ...interface{}) {
	log.logOut(INFO, nil, args...)
}

func Infof(format string, args ...interface{}) {
	log.logOut(INFO, &format, args...)
}

func Notice(args ...interface{}) {
	log.logOut(NOTICE, nil, args...)
}

func Noticef(format string, args ...interface{}) {
	log.logOut(NOTICE, &format, args...)
}

func Warning(args ...interface{}) {
	log.logOut(WARNING, nil, args...)
}

func Warningf(format string, args ...interface{}) {
	log.logOut(WARNING, &format, args...)
}

func Error(args ...interface{}) {
	log.logOut(ERROR, nil, args...)
}

func Errorf(format string, args ...interface{}) {
	log.logOut(ERROR, &format, args...)
}

func Critical(args ...interface{}) {
	log.logOut(CRITICAL, nil, args...)
}

func Criticalf(format string, args ...interface{}) {
	log.logOut(CRITICAL, &format, args...)
}

func Panic(args ...interface{}) {
	log.logOut(PANIC, nil, args...)
}

func Panicf(format string, args ...interface{}) {
	log.logOut(PANIC, &format, args...)
}

func Fatal(args ...interface{}) {
	log.logOut(FATAL, nil, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.logOut(FATAL, &format, args...)
}
