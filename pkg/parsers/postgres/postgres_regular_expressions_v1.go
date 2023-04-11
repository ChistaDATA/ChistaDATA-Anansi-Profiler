package postgres

import "regexp"

// LogMessageWithDurationRegEx is the regular expression
// eg: LOG:  duration: 247.351 ms
var LogMessageWithDurationRegEx *regexp.Regexp

// LogMessageWithNewQueryRegEx  eg: LOG:  statement: select * from salary where amount > rand order by event_id DESC limit 10
var LogMessageWithNewQueryRegEx *regexp.Regexp

// LogMessageWithEndQueryRegEx  eg: STATEMENT:  select * from salary where amount > rand order by event_id DESC limit 10
var LogMessageWithEndQueryRegEx *regexp.Regexp

// LogMessageWithErrorRegEx  eg: ERROR:  relation "salary" already exists
var LogMessageWithErrorRegEx *regexp.Regexp

func init() {
	LogMessageWithDurationRegEx = regexp.MustCompile("(?i)^LOG:  duration: ([0-9]*\\.[0-9]+) ms$")
	LogMessageWithNewQueryRegEx = regexp.MustCompile("(?i)^LOG:  statement: (.*)$")
	LogMessageWithEndQueryRegEx = regexp.MustCompile("(?i)^STATEMENT:  (.*)$")
	LogMessageWithErrorRegEx = regexp.MustCompile("(?i)^ERROR:  (.*)$")
}
