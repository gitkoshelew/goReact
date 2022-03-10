package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

// Logger ...
type Logger struct {
	*logrus.Entry
}

// GetLogger ...
func GetLogger() *Logger {
	return &Logger{e}
}

// GetLoggerWithField ...
func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

// Init ...
func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			return fmt.Sprintf(f.Function), fmt.Sprintf("%s:%d", path.Base(f.File), f.Line) // return File name (path.Base(f.File)) and File line
		},
		ForceColors:   true,
		FullTimestamp: true,
	}

	err := os.MkdirAll("../../logs", 0644)

	if err != nil || os.IsExist(err) {
		panic("can't create log dir. no configured logging to files")
	} else {
		allFile, err := os.OpenFile("../../logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			panic(fmt.Sprintf("[Message]: %s", err))
		}

		l.SetOutput(allFile)

		l.AddHook(&writerHook{
			Writer:    []io.Writer{allFile, os.Stdout},
			LogLevels: logrus.AllLevels,
		})
	}

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
