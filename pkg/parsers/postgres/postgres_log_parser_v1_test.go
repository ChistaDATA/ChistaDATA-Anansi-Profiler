package postgres

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"testing"
	"time"
)

func TestParseLogV1(t *testing.T) {

	SetParseLogV1Params("")

	logPostgresLogMap := map[string]stucts.ExtractedLog{
		"2023-04-10 06:46:35.539 UTC [1] LOG:  database system is ready to accept connections":      {Timestamp: time.Date(2023, 04, 10, 06, 46, 35, 539000000, time.UTC), ProcessID: 1, Message: "LOG:  database system is ready to accept connections"},
		"executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.": {Message: "executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec."},
	}

	for log, want := range logPostgresLogMap {
		got, _ := ParseLogV1(log)
		if got.Timestamp != want.Timestamp || got.ThreadId != want.ThreadId || got.QueryId != want.QueryId || got.LogLevel != want.LogLevel || got.Message != want.Message {
			t.Errorf("for log %s: got %v, want %v", log, got, want)
		}
	}

	logErrorMap := map[string]string{}

	for log, want := range logErrorMap {
		_, got := ParseLogV1(log)
		if !(got == nil && want == "") && (got == nil || got.Error() != want) {
			t.Errorf("for log %s: got %s, want %s", log, got, want)
		}
	}
}
