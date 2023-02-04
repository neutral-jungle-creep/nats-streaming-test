package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writer   []io.Writer
	LogLevel []logrus.Level
}

func (h *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range h.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (h *writerHook) Levels() []logrus.Level {
	return h.LogLevel
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func NewLogger() *Logger {
	return &Logger{e}
}

func init() {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", fileName, frame.Line)
		},
		FullTimestamp: true,
	}

	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(io.Discard)

	logger.AddHook(&writerHook{
		Writer:   []io.Writer{file, os.Stdout},
		LogLevel: logrus.AllLevels,
	})

	logger.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(logger)
}
