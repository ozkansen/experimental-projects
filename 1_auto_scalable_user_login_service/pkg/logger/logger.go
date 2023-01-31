package logger

import (
	"log"
	"os"
)

const (
	Discard = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
)

type Logger interface {
	Debug(msg string)
	Debugf(format string, args ...any)
	Info(msg string)
	Infof(format string, args ...any)
	Warn(msg string)
	Warnf(format string, args ...any)
	Error(msg string)
	Errorf(format string, args ...any)
	Fatal(msg string)
	Fatalf(format string, args ...any)
	Panic(msg string)
	Panicf(format string, args ...any)
}

var _ Logger = (*StdLogger)(nil)

type StdLogger struct {
	logLevel int
	stdOut   *log.Logger
	errOut   *log.Logger
}

func NewStdLogger(level int) *StdLogger {
	return &StdLogger{
		logLevel: level,
		stdOut:   log.New(os.Stderr, "", log.LstdFlags),
		errOut:   log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (l *StdLogger) Debug(msg string) {
	if l.logLevel > DebugLevel || l.logLevel == Discard {
		return
	}
	l.stdOut.Println("DBG", msg)
}

func (l *StdLogger) Debugf(format string, args ...any) {
	if l.logLevel > DebugLevel || l.logLevel == Discard {
		return
	}
	l.stdOut.Printf("DBG "+format+"\n", args...)
}

func (l *StdLogger) Info(msg string) {
	if l.logLevel > InfoLevel || l.logLevel == Discard {
		return
	}
	l.stdOut.Println("INF", msg)
}

func (l *StdLogger) Infof(format string, args ...any) {
	if l.logLevel > InfoLevel || l.logLevel == Discard {
		return
	}
	l.stdOut.Printf("INF "+format+"\n", args...)
}

func (l *StdLogger) Warn(msg string) {
	if l.logLevel > WarnLevel || l.logLevel == Discard {
		return
	}
	l.stdOut.Println("WRN", msg)
}

func (l *StdLogger) Warnf(format string, args ...any) {
	if l.logLevel > WarnLevel || l.logLevel == Discard {
		return
	}
	l.stdOut.Printf("WRN "+format+"\n", args...)
}

func (l *StdLogger) Error(msg string) {
	if l.logLevel > ErrorLevel || l.logLevel == Discard {
		return
	}
	l.errOut.Println("ERR", msg)
}

func (l *StdLogger) Errorf(format string, args ...any) {
	if l.logLevel > ErrorLevel || l.logLevel == Discard {
		return
	}
	l.errOut.Printf("ERR "+format+"\n", args...)
}

func (l *StdLogger) Fatal(msg string) {
	if l.logLevel == Discard {
		return
	}
	l.errOut.Fatal("FTL", msg)
}

func (l *StdLogger) Fatalf(format string, args ...any) {
	if l.logLevel == Discard {
		return
	}
	l.errOut.Fatalf("FTL "+format+"\n", args...)
}

func (l *StdLogger) Panic(msg string) {
	if l.logLevel == Discard {
		return
	}
	l.errOut.Panicln("PNC", msg)
}

func (l *StdLogger) Panicf(format string, args ...any) {
	if l.logLevel == Discard {
		return
	}
	l.errOut.Panicf("PNC "+format+"\n", args...)
}
