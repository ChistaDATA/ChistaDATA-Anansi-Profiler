package postgres

import "regexp"

// PostgresLogMessageWithDurationRegEx is the regular expression
// eg: LOG:  duration: 247.351 ms
var PostgresLogMessageWithDurationRegEx *regexp.Regexp

// PostgresLogMessageWithNewQueryRegEx  eg: LOG:  statement: select * from salary where amount > rand order by event_id DESC limit 10
var PostgresLogMessageWithNewQueryRegEx *regexp.Regexp

// PostgresLogMessageWithEndQueryRegEx  eg: STATEMENT:  select * from salary where amount > rand order by event_id DESC limit 10
var PostgresLogMessageWithEndQueryRegEx *regexp.Regexp

// PostgresLogMessageWithErrorRegEx  eg: ERROR:  relation "salary" already exists
var PostgresLogMessageWithErrorRegEx *regexp.Regexp

// PostgresLogMessageWithSystemUsageRegEx
// DETAIL:  ! system usage stats:
// !	0.005556 s user, 0.000000 s system, 0.006151 s elapsed
// !	[0.005556 s user, 0.001778 s system total]
// !	20100 kB max resident size
// !	0/0 [0/0] filesystem blocks in/out
// !	0/208 [0/777] page faults/reclaims, 0 [0] swaps
// !	0 [0] signals rcvd, 0/0 [0/0] messages rcvd/sent
// !	1/0 [8/1] voluntary/involuntary context switches
var PostgresLogMessageWithSystemUsageRegEx *regexp.Regexp

// PostgresLogMessageWithDurationAndQueryRegEx eg: LOG:  duration: 1.922 ms  statement: select * from shipments ;
var PostgresLogMessageWithDurationAndQueryRegEx *regexp.Regexp

func init() {
	PostgresLogMessageWithDurationRegEx = regexp.MustCompile("^LOG:  duration: ([0-9]*\\.[0-9]+) ([a-zA-Z]*)$")
	PostgresLogMessageWithNewQueryRegEx = regexp.MustCompile("^LOG:  statement: (.*)$")
	PostgresLogMessageWithEndQueryRegEx = regexp.MustCompile("^STATEMENT:  (.*)$")
	PostgresLogMessageWithErrorRegEx = regexp.MustCompile("^ERROR:  (.*)$")
	PostgresLogMessageWithSystemUsageRegEx = regexp.MustCompile(`(?s)DETAIL:  ! system usage stats:.*!	(\d+) ([a-zA-Z]+) max resident size.*`)
	PostgresLogMessageWithDurationAndQueryRegEx = regexp.MustCompile("(?s)LOG:  duration: ([0-9]*\\.[0-9]+) ([a-zA-Z]*)  (?:statement:|execute [a-zA-Z_][a-zA-Z0-9_]*:)(.*)")
}
