package clickhouse_resport_templates

const TopQueriesTemplate = `
# Profile
# Rank Response time   Calls R/Call Query
# ==== =============== ===== ====== =====
{{range $record:=.}}
{{$record.Pos}}|{{$record.TotalDuration}}|{{$record.TotalDurationPercentage}}|{{$record.Count}}|{{$record.ResponseTimePerCall}}|{{$record.Query}}
{{end}}
`

const TopQueryRecord = `# {{.Pos}} {{.TotalDuration}} {{.TotalDurationPercentage}} {{.Count}} {{.ResponseTimePerCall}} {{.Query}}`
