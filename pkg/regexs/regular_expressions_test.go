package regexs

import (
	"testing"
)

func TestClickHouseLogRegEx(t *testing.T) {

	logMatchMap := map[string]bool{
		"2022.09.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)":                true,
		"2022.09.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: (from 127.0.0.1:57216, user: default) select * from system.tables; (stage: Complete)": true,
		"2022.09.08 05:09:25.69635 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: (from 127.0.0.1:57216, user: default) select * from system.tables; (stage: Complete)":  false,
	}

	for log, want := range logMatchMap {
		got := ClickHouseLogRegEx.MatchString(log)
		if got != want {
			t.Errorf("for log %s: got %t, want %t", log, got, want)
		}
	}
}

func TestLogMessageWithQueryInfoRegEx(t *testing.T) {

	logMatchMap := map[string]bool{
		"executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)":                                   true,
		"executeQuery: (from 127.0.0.1:57216, user: default) select * from system.tables; (stage: Complete)":                    true,
		"executeQuery: (from 127.0.0.1:57216, user: default)(comment) select * from system.tables; (stage: Complete)":           true,
		"executeQuery: (from 127.0.0.1:57216, user: default)(comment)(comment2) select * from system.tables; (stage: Complete)": true,
		"(from 127.0.0.1:57216, user: default)(comment)(comment2) select * from system.tables; (stage: Complete)":               false,
	}

	for log, want := range logMatchMap {
		got := LogMessageWithQueryInfoRegEx.MatchString(log)
		if got != want {
			t.Errorf("for log %s: got %t, want %t", log, got, want)
		}
	}
}

func TestLogMessageWithPeakMemory(t *testing.T) {

	logMatchMap := map[string]bool{
		"MemoryTracker: Peak memory usage (for query): 440.67 KiB.": true,
		"MemoryTracker: Peak memory usage: 440.67 KiB.":             true,
		"MemoryTracker: Peak memory usage: .67 PiB.":                false,
		"MemoryTracker: Peak memory usage: 1. PiB.":                 false,
	}

	for log, want := range logMatchMap {
		got := LogMessageWithPeakMemoryRegEx.MatchString(log)
		if got != want {
			t.Errorf("for log %s: got %t, want %t", log, got, want)
		}
	}
}

func TestLogMessageWithErrorRegEx(t *testing.T) {

	logMatchMap := map[string]bool{
		"executeQuery: Code: 60. DB::Exception: Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build)) (from 127.0.0.1:57236) (in query: select * from default.table;), Stack trace (when copying this message, always include the lines below):": true,
		"executeQuery: Code: 60. DB::Exception: Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build)) (from 127.0.0.1:57236) (in query: select * from default.table;)":                                                                           true,
		"executeQuery: Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build)) (from 127.0.0.1:57236) (in query: select * from default.table;)":                                                                                                    true,
	}

	for log, want := range logMatchMap {
		got := LogMessageWithErrorRegEx.MatchString(log)
		if got != want {
			t.Errorf("for log %s: got %t, want %t", log, got, want)
		}
	}
}

func TestLogMessageWithAccessInfoRegEx(t *testing.T) {

	logMatchMap := map[string]bool{
		"ContextAccess (default): Access granted: SHOW TABLES ON *.*":                                      true,
		"ContextAccess (default): Access granted: INSERT(radio, mcc, net, created) ON default.cell_towers": true,
	}

	for log, want := range logMatchMap {
		got := LogMessageWithAccessInfoRegEx.MatchString(log)

		if got != want {
			t.Errorf("for log %s: got %t, want %t", log, got, want)
		}
	}
}
