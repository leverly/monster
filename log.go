package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

type LogFormatter struct{}

func (m *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] [%s] [%s:%d] %s\n",
			timestamp, entry.Level, fName, entry.Caller.Line, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}
