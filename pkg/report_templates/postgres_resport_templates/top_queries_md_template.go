package postgres_resport_templates

const TopQueriesMDTemplate = `
# Top Queries
| Rank | Response time  | Calls | R/Call | Query |   
|------|----------------|-------|--------|-------|
{{.Records}}
`

const TopQueryMDRecord = `| {{.Pos}} | {{.TotalDuration}} {{.TotalDurationPercentage}} |  {{.Count}} | {{.ResponseTimePerCall}} | {{.Query}} |`
