package postgres

import "testing"

func TestLogMessageWithDurationRegEx(t *testing.T) {

	logMatchMap := map[string]bool{
		"executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)": false,
		"LOG:  duration: 247.351 ms": true,
	}

	for log, want := range logMatchMap {
		got := PostgresLogMessageWithDurationRegEx.MatchString(log)
		if got != want {
			t.Errorf("for log %s: got %t, want %t", log, got, want)
		}
	}
}

func TestLogMessageWithNewQueryRegEx(t *testing.T) {

	logMatchMap := map[string]bool{
		"executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)":       false,
		"LOG:  statement: select * from salary where amount > rand order by event_id DESC limit 10": true,
	}

	for log, want := range logMatchMap {
		got := PostgresLogMessageWithNewQueryRegEx.MatchString(log)
		if got != want {
			t.Errorf("for log %s: got %t, want %t", log, got, want)
		}
	}
}

func TestLogMessageWithEndQueryRegEx(t *testing.T) {

	logMatchMap := map[string]bool{
		"executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)":  false,
		"STATEMENT:  select * from salary where amount > rand order by event_id DESC limit 10": true,
	}

	for log, want := range logMatchMap {
		got := PostgresLogMessageWithEndQueryRegEx.MatchString(log)
		if got != want {
			t.Errorf("for log %s: got %t, want %t", log, got, want)
		}
	}
}

func TestLogMessageWithErrorRegEx(t *testing.T) {

	logMatchMap := map[string]bool{
		"executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)": false,
		"ERROR:  relation \"salary\" already exists":                                          true,
	}

	for log, want := range logMatchMap {
		got := PostgresLogMessageWithErrorRegEx.MatchString(log)
		if got != want {
			t.Errorf("for log %s: got %t, want %t", log, got, want)
		}
	}
}
