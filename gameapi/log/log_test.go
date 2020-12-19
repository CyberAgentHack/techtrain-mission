package log

import (
	"bytes"
	"fmt"
	"testing"
)

var buf bytes.Buffer

type testLogger struct{}

func (t *testLogger) Warnf(format string, v ...interface{}) {
	buf.WriteString(fmt.Sprintf(format, v...))
}

func (t *testLogger) Debugf(format string, v ...interface{}) {
	if debugOn {
		buf.WriteString(fmt.Sprintf(format, v...))
	}
}

func TestLogger(t *testing.T) {
	t.Parallel()

	SetLogger(&testLogger{})

	cases := []struct {
		name       string
		loggerFunc func(string, ...interface{})
		debugOn    bool
		input      string
		wanted     string
	}{
		{name: "warnf with debugOn", loggerFunc: logger.Warnf, debugOn: true, input: "warning", wanted: "warning"},
		{name: "warnf with debugOff", loggerFunc: logger.Warnf, debugOn: false, input: "warning", wanted: "warning"},
		{name: "debug with debugOn", loggerFunc: logger.Debugf, debugOn: true, input: "debug", wanted: "debug"},
		{name: "debug with debugOff", loggerFunc: logger.Debugf, debugOn: false, input: "debug", wanted: ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				buf.Reset()
			}()

			debugOn = tc.debugOn
			tc.loggerFunc(tc.input)

			actual := buf.String()
			if actual != tc.wanted {
				t.Errorf("actual: %s, wanted: %s", actual, tc.wanted)
			}

			buf.Reset()
		})
	}

}
