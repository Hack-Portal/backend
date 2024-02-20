package logger

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	OFF
)

type Logger interface {
	Debug(interface{})
	Debugf(string, ...interface{})
	Info(interface{})
	Infof(string, ...interface{})
	Warn(interface{})
	Warnf(string, ...interface{})
	Error(interface{})
	Errorf(string, ...interface{})
	Panic(interface{})
	Panicf(string, ...interface{})
}

type LoggerConfig struct{ LogLevel }

type logger struct{ LogLevel }

func NewLogger(lv LogLevel) Logger {
	return logger{lv}
}

func (l logger) Debug(v interface{}) {
	if l.LogLevel > DEBUG {
		return
	}
	fmt.Printf("%v [DEBUG] %v\n", l.formattedTime(), v)
}

func (l logger) Debugf(format string, v ...interface{}) {
	if l.LogLevel > DEBUG {
		return
	}
	fmt.Printf("%v [DEBUG] %s\n", l.formattedTime(), fmt.Sprintf(format, v...))
}

func (l logger) Info(v interface{}) {
	if l.LogLevel > INFO {
		return
	}
	fmt.Printf("%v [INFO] %v\n", l.formattedTime(), v)
}

func (l logger) Infof(format string, v ...interface{}) {
	if l.LogLevel > INFO {
		return
	}
	fmt.Printf("%v [INFO] %s\n", l.formattedTime(), fmt.Sprintf(format, v...))
}

func (l logger) Warn(v interface{}) {
	if l.LogLevel > WARN {
		return
	}
	fmt.Printf("%v [WARN] %v\n", l.formattedTime(), v)
}

func (l logger) Warnf(format string, v ...interface{}) {
	if l.LogLevel > WARN {
		return
	}
	fmt.Printf("%v [WARN] %s\n", l.formattedTime(), fmt.Sprintf(format, v...))
}

func (l logger) Error(v interface{}) {
	if l.LogLevel > ERROR {
		return
	}
	fmt.Printf("%v [ERROR] %v\n", l.formattedTime(), v)
}

func (l logger) Errorf(format string, v ...interface{}) {
	if l.LogLevel > ERROR {
		return
	}
	fmt.Printf("%v [ERROR] %s\n", l.formattedTime(), fmt.Sprintf(format, v...))
}

func (l logger) Panic(v interface{}) {
	l.Error(v)
	panic(v)
}

func (l logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Error(s)
	panic(s)
}

func (l logger) formattedTime() string {
	return time.Now().Format("2006-01-02 15:04:05.000000 +0900 UTC")
}
