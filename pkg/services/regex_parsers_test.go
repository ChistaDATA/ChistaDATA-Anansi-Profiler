package services

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	"reflect"
	"testing"
	"time"
)

func TestParseClickHouseLog(t *testing.T) {

	logClickHouseLogMap := map[string]stucts.ClickHouseLog{
		"2022.09.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)":       {Timestamp: time.Date(2022, 9, 8, 5, 9, 25, 696359000, time.UTC), ThreadId: 46, QueryId: "add4a8af-c695-4378-9e55-5345bdea5998", LogLevel: "Debug", Message: "executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)"},
		"2022.09.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.": {Timestamp: time.Date(2022, 9, 8, 5, 9, 25, 696359000, time.UTC), ThreadId: 46, QueryId: "add4a8af-c695-4378-9e55-5345bdea5998", LogLevel: "Debug", Message: "executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec."},
		"2022.10.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.": {Timestamp: time.Date(2022, 10, 8, 5, 9, 25, 696359000, time.UTC), ThreadId: 46, QueryId: "add4a8af-c695-4378-9e55-5345bdea5998", LogLevel: "Debug", Message: "executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec."},
		"2022.10.08 05:09:25.696351 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.": {Timestamp: time.Date(2022, 10, 8, 5, 9, 25, 696351000, time.UTC), ThreadId: 46, QueryId: "add4a8af-c695-4378-9e55-5345bdea5998", LogLevel: "Debug", Message: "executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec."},
	}

	for log, want := range logClickHouseLogMap {
		got, _ := ParseClickHouseLog(log)
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
		_, got := ParseClickHouseLog(log)
		if !(got == nil && want == "") && (got == nil || got.Error() != want) {
			t.Errorf("for log %s: got %s, want %s", log, got, want)
		}
	}
}

func TestParseMessageWithQuery(t *testing.T) {

	messageQueryMap := map[string]stucts.Query{
		"executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)":                                                                    {ClientHost: "127.0.0.1", ClientPort: 57216, Query: "select * from system.tables;", User: "default"},
		"executeQuery: (from 127.0.0.1:57216, user: joe) select * from system.tables; (stage: Complete)":                                                         {ClientHost: "127.0.0.1", ClientPort: 57216, Query: "select * from system.tables;", User: "joe"},
		"executeQuery: (from 127.0.0.1:57216, user: joe, initial_query_id: add4a8af-c695-4378-9e55-5345bdea5998) select * from system.tables; (stage: Complete)": {ClientHost: "127.0.0.1", ClientPort: 57216, Query: "select * from system.tables;", User: "joe", InitialQueryId: "add4a8af-c695-4378-9e55-5345bdea5998"},
	}

	for message, want := range messageQueryMap {
		got := &stucts.Query{}
		err := ParseMessageWithQuery(message, got)
		if err != nil || got.ClientHost != want.ClientHost || got.ClientPort != want.ClientPort || got.Query != want.Query || got.User != want.User || got.InitialQueryId != want.InitialQueryId {
			t.Errorf("for message %s: error: %s, got %v, want %v", message, err, got, want)
		}
	}

	messageErrorMap := map[string]string{
		"executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)":                        "",
		"executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.":                  "error parsing message as ExecuteQueryLogMessageWithQueryRegEx",
		"executeQuery: (from 127.0.0.1:5721g) select * from system.tables; (stage: Complete)":                        "error parsing message as ExecuteQueryLogMessageWithQueryRegEx",
		"executeQuery: (from 684D:1111:222:3333:4444:5555:6:77:5721) select * from system.tables; (stage: Complete)": "",
	}

	for message, want := range messageErrorMap {
		got := ParseMessageWithQuery(message, &stucts.Query{})
		if !(got == nil && want == "") && (got == nil || got.Error() != want) {
			t.Errorf("for message %s: got %s, want %s", message, got, want)
		}
	}
}

func TestParseMessageWithDataInfo(t *testing.T) {

	messageQueryMap := map[string]stucts.Query{
		"executeQuery: Read 65 rows, 105.59 KiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.":          {ReadRows: 65, ReadBytes: 105.59 * 1024, Duration: 0.016169625},
		"executeQuery: Read 65 rows, 105 B in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.":               {ReadRows: 65, ReadBytes: 105, Duration: 0.016169625},
		"executeQuery: Read 65 rows, 105 MiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.":             {ReadRows: 65, ReadBytes: 105 * 1024 * 1024, Duration: 0.016169625},
		"executeQuery: Read 65 rows, 105 GiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.":             {ReadRows: 65, ReadBytes: 105 * 1024 * 1024 * 1024, Duration: 0.016169625},
		"executeQuery: Read 65 rows, 105000000000000 GiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.": {ReadRows: 65, ReadBytes: 105000000000000 * 1024 * 1024 * 1024, Duration: 0.016169625},
	}

	for message, want := range messageQueryMap {
		got := &stucts.Query{}
		err := ParseMessageWithDataInfo(message, got)
		if err != nil || got.ReadRows != want.ReadRows || got.ReadBytes != want.ReadBytes || got.Duration != want.Duration {
			t.Errorf("for message %s: error: %s, got %v, want %v", message, err, got, want)
		}
	}

	messageErrorMap := map[string]string{
		"executeQuery: Read 65 rows, 105.59 PiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.": "size format PiB not supported",
		"executeQuery: Read 65 rows, .59 PiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.":    "error parsing message as LogMessageWithDataRegEx",
	}

	for message, want := range messageErrorMap {
		got := ParseMessageWithDataInfo(message, &stucts.Query{})
		if !(got == nil && want == "") && (got == nil || got.Error() != want) {
			t.Errorf("for message %s: got %s, want %s", message, got, want)
		}
	}
}

func TestParseMessageWithPeakMemory(t *testing.T) {

	messageQueryMap := map[string]float64{
		"MemoryTracker: Peak memory usage (for query): 440.67 KiB.": 440.67 * 1024,
		"MemoryTracker: Peak memory usage: 440.67 KiB.":             440.67 * 1024,
		"MemoryTracker: Peak memory usage: 440 B.":                  440,
		"MemoryTracker: Peak memory usage: 440.67 MiB.":             440.67 * 1024 * 1024,
		"MemoryTracker: Peak memory usage: 440.67 GiB.":             440.67 * 1024 * 1024 * 1024,
	}

	for message, want := range messageQueryMap {
		got := &stucts.Query{}
		err := ParseMessageWithPeakMemory(message, got)
		if err != nil || got.PeakMemoryUsage != want {
			t.Errorf("for message %s: error: %s, got %v, want %v", message, err, got.PeakMemoryUsage, want)
		}
	}

	messageErrorMap := map[string]string{
		"MemoryTracker: Peak memory usage: 440. KiB.": "error parsing message as LogMessageWithPeakMemoryRegEx",
		"MemoryTracker: Peak memory usage: 440 PiB.":  "size format PiB not supported",
	}

	for message, want := range messageErrorMap {
		got := ParseMessageWithPeakMemory(message, &stucts.Query{})
		if !(got == nil && want == "") && (got == nil || got.Error() != want) {
			t.Errorf("for message %s: got %s, want %s", message, got, want)
		}
	}
}

func TestParseMessageWithErrorInfo(t *testing.T) {

	messageQueryMap := map[string]stucts.Query{
		"executeQuery: Code: 60. DB::Exception: Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build)) (from 127.0.0.1:57236) (in query: select * from default.table;), Stack trace (when copying this message, always include the lines below):": {ErrorCompleteText: "Code: 60. DB::Exception: Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build))", ErrorCode: "60", ErrorMessage: "Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build))", ErrorStackTrace: ", Stack trace (when copying this message, always include the lines below):"},
		"executeQuery: Code: 60. DB::Exception: Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build)) (from 127.0.0.1:57236) (in query: select * from default.table;)":                                                                           {ErrorCompleteText: "Code: 60. DB::Exception: Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build))", ErrorCode: "60", ErrorMessage: "Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build))", ErrorStackTrace: ""},
		"executeQuery: Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build)) (from 127.0.0.1:57236) (in query: select * from default.table;)":                                                                                                    {ErrorCompleteText: "Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build))", ErrorCode: "", ErrorMessage: "", ErrorStackTrace: ""},
	}

	for message, want := range messageQueryMap {
		got := &stucts.Query{}
		err := ParseMessageWithErrorInfo(message, got)
		if err != nil || got.ErrorMessage != want.ErrorMessage || got.ErrorCode != want.ErrorCode || got.ErrorCompleteText != want.ErrorCompleteText || got.ErrorStackTrace != got.ErrorStackTrace {
			t.Errorf("for message %s: error: %s, got %v, want %v", message, err, got, want)
		}
	}

	messageErrorMap := map[string]string{
		"executeQuery : Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build)) (from 127.0.0.1:57236) (in query: select * from default.table;)": "error parsing message as LogMessageWithErrorRegEx",
	}

	for message, want := range messageErrorMap {
		got := ParseMessageWithErrorInfo(message, &stucts.Query{})
		if !(got == nil && want == "") && (got == nil || got.Error() != want) {
			t.Errorf("for message %s: got %s, want %s", message, got, want)
		}
	}
}

func TestParseMessageWithAccessInfo(t *testing.T) {

	messageQueryMap := map[string]stucts.Query{
		"ContextAccess (default): Access granted: SHOW TABLES ON *.*":                                      {Databases: map[string]struct{}{"*": {}}, Tables: map[string]struct{}{"*": {}}},
		"ContextAccess (default): Access granted: INSERT(radio, mcc, net, created) ON default.cell_towers": {Databases: map[string]struct{}{"default": {}}, Tables: map[string]struct{}{"cell_towers": {}}},
	}

	for message, want := range messageQueryMap {
		got := &stucts.Query{Databases: types.InitStringSet(), Tables: types.InitStringSet()}
		err := ParseMessageWithAccessInfo(message, got)
		if err != nil || !reflect.DeepEqual(got.Tables, want.Tables) || !reflect.DeepEqual(got.Databases, want.Databases) {
			t.Errorf("for message %s: error: %s, got %v, want %v", message, err, got, want)
		}
	}

	messageErrorMap := map[string]string{
		"ContextAccess (default): Access granted: SHOW TABLES ON *.": "error parsing message as LogMessageWithAccessInfoRegEx",
	}

	for message, want := range messageErrorMap {
		got := ParseMessageWithAccessInfo(message, &stucts.Query{Databases: types.InitStringSet(), Tables: types.InitStringSet()})
		if !(got == nil && want == "") && (got == nil || got.Error() != want) {
			t.Errorf("for message %s: got %s, want %s", message, got, want)
		}
	}

}
