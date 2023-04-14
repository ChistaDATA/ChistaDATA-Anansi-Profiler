package postgres

import (
	"testing"
)

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

func TestPostgresLogMessageWithSystemUsageRegEx(t *testing.T) {

	logMatchMap := map[string]bool{
		//"executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)": false,
		`DETAIL:  ! system usage stats:
	!	0.005556 s user, 0.000000 s system, 0.006151 s elapsed
	!	[0.005556 s user, 0.001778 s system total]
	!	20100 kB max resident size
	!	0/0 [0/0] filesystem blocks in/out
	!	0/208 [0/777] page faults/reclaims, 0 [0] swaps
	!	0 [0] signals rcvd, 0/0 [0/0] messages rcvd/sent
	!	1/0 [8/1] voluntary/involuntary context switches`: true,
	}

	for log, want := range logMatchMap {
		got := PostgresLogMessageWithSystemUsageRegEx.MatchString(log)
		if got != want {
			t.Errorf("for log %s: got %t, want %t", log, got, want)
		}
	}
}

func TestPostgresLogMessageWithDurationAndQueryRegEx(t *testing.T) {

	logMatchMap := map[string]bool{
		"executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)": false,
		`LOG:  duration: 1.922 ms  statement: select * from shipments ;`:                      true,
	}

	for log, want := range logMatchMap {
		got := PostgresLogMessageWithDurationAndQueryRegEx.MatchString(log)
		if got != want {
			t.Errorf("for log %s: got %t, want %t", log, got, want)
		}
	}
}
