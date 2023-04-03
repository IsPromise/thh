package logger

import (
	"bytes"
	"fmt"
	"thh/arms"

	"github.com/sirupsen/logrus"
)

// TextFormatter formats logs into text
type TextFormatter struct {
}

// Format renders a single log entry
func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	trace := arms.MyTrace()
	b.WriteString(fmt.Sprintf("%v %v", trace.GetNextTrace(), entry.Message))
	b.WriteByte('\n')
	return b.Bytes(), nil
}
