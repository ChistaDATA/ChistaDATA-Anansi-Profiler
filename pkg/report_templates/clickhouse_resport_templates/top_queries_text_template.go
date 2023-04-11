package clickhouse_resport_templates

const TopQueriesTemplate = `
# Profile
# Rank Response time   Calls R/Call Query
# ==== =============== ===== ====== =====
{{.Records}}
`

const TopQueryRecord = `# {{.Pos}} {{.TotalDuration}} {{.TotalDurationPercentage}} {{.Count}} {{.ResponseTimePerCall}} {{.Query}}`
