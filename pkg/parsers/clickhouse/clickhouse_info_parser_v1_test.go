package clickhouse

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"reflect"
	"testing"
)

func TestParseMessageWithQuery(t *testing.T) {

	messageQueryMap := map[string]stucts.Query{
		"executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)":                                                                    {ClientHost: "127.0.0.1", ClientPort: 57216, Query: "select * from system.tables;", User: "default"},
		"executeQuery: (from 127.0.0.1:57216, user: joe) select * from system.tables; (stage: Complete)":                                                         {ClientHost: "127.0.0.1", ClientPort: 57216, Query: "select * from system.tables;", User: "joe"},
		"executeQuery: (from 127.0.0.1:57216, user: joe, initial_query_id: add4a8af-c695-4378-9e55-5345bdea5998) select * from system.tables; (stage: Complete)": {ClientHost: "127.0.0.1", ClientPort: 57216, Query: "select * from system.tables;", User: "joe", InitialQueryId: "add4a8af-c695-4378-9e55-5345bdea5998"},
	}

	for message, want := range messageQueryMap {
		got := stucts.InitDBPerfInfoRepository()
		err := ParseMessageWithQueryV1(stucts.ExtractedLog{Message: message}, got)
		if err != nil || got.Queries.GetQuery(message).ClientHost != want.ClientHost || got.Queries.GetQuery(message).ClientPort != want.ClientPort || got.Queries.GetQuery(message).Query != want.Query || got.Queries.GetQuery(message).User != want.User || got.Queries.GetQuery(message).InitialQueryId != want.InitialQueryId {
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
		got := ParseMessageWithQueryV1(stucts.ExtractedLog{Message: message}, stucts.InitDBPerfInfoRepository())
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
		got := stucts.InitDBPerfInfoRepository()
		err := ParseMessageWithDataInfoV1(stucts.ExtractedLog{Message: message}, got)
		if err != nil || got.Queries.GetQuery(message).ReadRows != want.ReadRows || got.Queries.GetQuery(message).ReadBytes != want.ReadBytes || got.Queries.GetQuery(message).Duration != want.Duration {
			t.Errorf("for message %s: error: %s, got %v, want %v", message, err, got, want)
		}
	}

	messageErrorMap := map[string]string{
		"executeQuery: Read 65 rows, 105.59 PiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.": "size format PiB not supported",
		"executeQuery: Read 65 rows, .59 PiB in 0.016169625 sec., 4019 rows/sec., 6.38 MiB/sec.":    "error parsing message as LogMessageWithDataRegEx",
	}

	for message, want := range messageErrorMap {
		got := ParseMessageWithDataInfoV1(stucts.ExtractedLog{Message: message}, stucts.InitDBPerfInfoRepository())
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
		got := stucts.InitDBPerfInfoRepository()
		err := ParseMessageWithPeakMemoryV1(stucts.ExtractedLog{Message: message}, got)
		if err != nil || got.Queries.GetQuery(message).PeakMemoryUsage != want {
			t.Errorf("for message %s: error: %s, got %v, want %v", message, err, got.Queries.GetQuery(message).PeakMemoryUsage, want)
		}
	}

	messageErrorMap := map[string]string{
		"MemoryTracker: Peak memory usage: 440. KiB.": "error parsing message as LogMessageWithPeakMemoryRegEx",
		"MemoryTracker: Peak memory usage: 440 PiB.":  "size format PiB not supported",
	}

	for message, want := range messageErrorMap {
		got := ParseMessageWithPeakMemoryV1(stucts.ExtractedLog{Message: message}, stucts.InitDBPerfInfoRepository())
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
		got := stucts.InitDBPerfInfoRepository()
		err := ParseMessageWithErrorInfoV1(stucts.ExtractedLog{Message: message}, got)
		if err != nil || got.Queries.GetQuery(message).ErrorMessage != want.ErrorMessage || got.Queries.GetQuery(message).ErrorCode != want.ErrorCode || got.Queries.GetQuery(message).ErrorCompleteText != want.ErrorCompleteText || got.Queries.GetQuery(message).ErrorStackTrace != want.ErrorStackTrace {
			t.Errorf("for message %s: error: %s, got %v, want %v", message, err, got, want)
		}
	}

	messageErrorMap := map[string]string{
		"executeQuery : Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build)) (from 127.0.0.1:57236) (in query: select * from default.table;)": "error parsing message as LogMessageWithErrorRegEx",
	}

	for message, want := range messageErrorMap {
		got := ParseMessageWithErrorInfoV1(stucts.ExtractedLog{Message: message}, stucts.InitDBPerfInfoRepository())
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
		got := stucts.InitDBPerfInfoRepository()
		err := ParseMessageWithAccessInfoV1(stucts.ExtractedLog{Message: message}, got)
		if err != nil || !reflect.DeepEqual(got.Queries.GetQuery(message).Tables, want.Tables) || !reflect.DeepEqual(got.Queries.GetQuery(message).Databases, want.Databases) {
			t.Errorf("for message %s: error: %s, got %v, want %v", message, err, got, want)
		}
	}

	messageErrorMap := map[string]string{
		"ContextAccess (default): Access granted: SHOW TABLES ON *.": "error parsing message as LogMessageWithAccessInfoRegEx",
	}

	for message, want := range messageErrorMap {
		got := ParseMessageWithAccessInfoV1(stucts.ExtractedLog{Message: message}, stucts.InitDBPerfInfoRepository())
		if !(got == nil && want == "") && (got == nil || got.Error() != want) {
			t.Errorf("for message %s: got %s, want %s", message, got, want)
		}
	}

}

func TestParseMessageWithDBInfo(t *testing.T) {
	messageQueryMap := map[string]stucts.Query{
		"poc.wide_item_daily (368dfb93-a0dc-4220-80d3-501a873db5e6) (SelectExecutor): Key condition: unknown": {Databases: map[string]struct{}{"poc": {}}, Tables: map[string]struct{}{"wide_item_daily": {}}},
	}

	for message, want := range messageQueryMap {
		got := stucts.InitDBPerfInfoRepository()
		err := ParseMessageWithDBInfo(stucts.ExtractedLog{Message: message}, got)
		tes := got.Queries.GetQuery(message)
		if err != nil || !reflect.DeepEqual(tes, want.Tables) || !reflect.DeepEqual(got.Queries.GetQuery(message).Databases, want.Databases) {
			t.Errorf("for message %s: error: %s, got %v, want %v", message, err, got, want)
		}
	}

	//messageErrorMap := map[string]string{
	//	"ContextAccess (default): Access granted: SHOW TABLES ON *.": "error parsing message as LogMessageWithAccessInfoRegEx",
	//}
	//
	//for message, want := range messageErrorMap {
	//	got := ParseMessageWithAccessInfoV1(stucts.ExtractedLog{Message: message}, stucts.InitDBPerfInfoRepository())
	//	if !(got == nil && want == "") && (got == nil || got.Error() != want) {
	//		t.Errorf("for message %s: got %s, want %s", message, got, want)
	//	}
	//}
}
