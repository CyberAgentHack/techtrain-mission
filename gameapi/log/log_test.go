package log_test

import (
	"testing"

	"github.com/task4233/techtrain-mission/gameapi/log"
)

var (
	logger = log.MyLogger
)

// TODO: テストが書けない......
func TestWarnf(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		debugOn     bool
		inputFormat string
		inputData   string
		want        string
	}{
		{
			name:        "test環境で実行",
			debugOn:     true,
			inputFormat: "Warnf: %s",
			inputData:   "TestWarnf",
			want:        "Warnf: TestWarnf",
		},
		{
			name:        "test環境以外で実行",
			debugOn:     false,
			inputFormat: "Warnf: %s",
			inputData:   "TestWarnf",
			want:        "Warnf: TestWarnf",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			log.SetDebugStatus(tc.debugOn)
			logger.Warnf(tc.inputFormat, tc.inputData)
			// TODO: 出力を見れるようにしたい
		})
	}
}

func TestDebugf(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		debugOn     bool
		inputFormat string
		inputData   string
		want        string
	}{
		{
			name:        "test環境で実行",
			debugOn:     true,
			inputFormat: "Debugf: %s",
			inputData:   "TestBebugf",
			want:        "[DEBUG] Warnf: TestDebugf",
		},
		{
			name:        "test環境以外で実行",
			debugOn:     false,
			inputFormat: "Debugf: %s",
			inputData:   "TestDebugf",
			want:        "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			log.SetDebugStatus(tc.debugOn)
			logger.Debugf(tc.inputFormat, tc.inputData)
			// TODO: 出力を見れるようにしたい
		})
	}
}
