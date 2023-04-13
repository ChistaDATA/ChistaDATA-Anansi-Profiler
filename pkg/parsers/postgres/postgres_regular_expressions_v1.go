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

func init() {
	PostgresLogMessageWithDurationRegEx = regexp.MustCompile("^LOG:  duration: ([0-9]*\\.[0-9]+) ([a-zA-Z]*)$")
	PostgresLogMessageWithNewQueryRegEx = regexp.MustCompile("^LOG:  statement: (.*)$")
	PostgresLogMessageWithEndQueryRegEx = regexp.MustCompile("^STATEMENT:  (.*)$")
	PostgresLogMessageWithErrorRegEx = regexp.MustCompile("^ERROR:  (.*)$")
}
