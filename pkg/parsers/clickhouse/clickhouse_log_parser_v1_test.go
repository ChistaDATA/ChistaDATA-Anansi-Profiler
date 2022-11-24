package clickhouse

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"testing"
	"time"
)

func TestParseClickHouseLog(t *testing.T) {

	logClickHouseLogMap := map[string]stucts.ExtractedLog{
		"2022.09.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)":       {Timestamp: time.Date(2022, 9, 8, 5, 9, 25, 696359000, time.UTC), ThreadId: 46, QueryId: "add4a8af-c695-4378-9e55-5345bdea5998", LogLevel: "Debug", Message: "executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)"},
		"2022.09.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.": {Timestamp: time.Date(2022, 9, 8, 5, 9, 25, 696359000, time.UTC), ThreadId: 46, QueryId: "add4a8af-c695-4378-9e55-5345bdea5998", LogLevel: "Debug", Message: "executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec."},
		"2022.10.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.": {Timestamp: time.Date(2022, 10, 8, 5, 9, 25, 696359000, time.UTC), ThreadId: 46, QueryId: "add4a8af-c695-4378-9e55-5345bdea5998", LogLevel: "Debug", Message: "executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec."},
		"2022.10.08 05:09:25.696351 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.": {Timestamp: time.Date(2022, 10, 8, 5, 9, 25, 696351000, time.UTC), ThreadId: 46, QueryId: "add4a8af-c695-4378-9e55-5345bdea5998", LogLevel: "Debug", Message: "executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec."},
	}

	for log, want := range logClickHouseLogMap {
		got, _ := ParseLogV1(log)
		if got.Timestamp != want.Timestamp || got.ThreadId != want.ThreadId || got.QueryId != want.QueryId || got.LogLevel != want.LogLevel || got.Message != want.Message {
			t.Errorf("for log %s: got %v, want %v", log, got, want)
		}
	}

	logErrorMap := map[string]string{
		"executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.":     "error parsing ClickHouse log, part size not 6",
		"2022.99.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> invalid time": "parsing time \"2022.99.08 05:09:25.696359\": month out of range",
		"2022.09.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea599g} <Debug> invalid uuid": "error parsing ClickHouse log, part size not 6",
	}

	for log, want := range logErrorMap {
		_, got := ParseLogV1(log)
		if !(got == nil && want == "") && (got == nil || got.Error() != want) {
			t.Errorf("for log %s: got %s, want %s", log, got, want)
		}
	}
}
