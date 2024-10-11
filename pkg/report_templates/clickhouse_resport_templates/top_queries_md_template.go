package clickhouse_resport_templates

const TopQueriesMDTemplate = `
# Top Queries
| Rank | Response time  | Calls | R/Call | Query |   
|------|----------------|-------|--------|-------|
{{range $record:=.}}
|{{$record.Pos}}|{{$record.TotalDuration}}|{{$record.TotalDurationPercentage}}|{{$record.Count}}|{{$record.ResponseTimePerCall}}|{{$record.Query}}|
{{end}}
`

const TopQueryMDRecord = `| {{.Pos}} | {{.TotalDuration}} {{.TotalDurationPercentage}} |  {{.Count}} | {{.ResponseTimePerCall}} | {{.Query}} |`
