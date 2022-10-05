package stucts

import (
	"bytes"
	"fmt"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/formatters"
	"strconv"
	"strings"
	"text/template"
)

// TopQueriesTemplateInput Input for TopQueriesTextTemplate
type TopQueriesTemplateInput struct {
	Records string
}

func InitTopQueriesTemplateInput(records []TopQueriesTemplateInputRecord, temp *template.Template) TopQueriesTemplateInput {
	recordString := ""
	for _, record := range records {
		var bf bytes.Buffer
		err := temp.Execute(&bf, record)
		if err == nil {
			recordString = recordString + bf.String() + "\n"
		}
	}
	return TopQueriesTemplateInput{Records: recordString}
}

type TopQueriesTemplateInputRecord struct {
	Query                   string
	Pos                     string
	TotalDuration           string
	Count                   string
	TotalDurationPercentage string
	ResponseTimePerCall     string
}

func InitTopQueriesTemplateInputRecord(info *SimilarQueryInfo, queryInfoTemplateInput *QueryInfoTemplateInput, totalExecutionTime float64) TopQueriesTemplateInputRecord {
	queryLengthLimit := 80
	if len(info.Query) < queryLengthLimit {
		queryLengthLimit = len(info.Query)
	}
	return TopQueriesTemplateInputRecord{
		Query:                   info.Query[:queryLengthLimit],
		Pos:                     formatters.PrefixSpace(strings.TrimSpace(queryInfoTemplateInput.Pos), 4),
		TotalDuration:           formatters.PrefixSpace(formatters.Float64SecondsToString(info.GetTotalDuration()), 7),
		Count:                   formatters.PrefixSpace(strconv.Itoa(info.Count), 5),
		TotalDurationPercentage: formatters.PrefixSpace(fmt.Sprintf("%.2f%s", info.GetTotalDuration()*100/totalExecutionTime, "%"), 7),
		ResponseTimePerCall:     formatters.PrefixSpace(formatters.Float64SecondsToString(info.GetTotalDuration()/float64(info.Count)), 6),
	}
}
