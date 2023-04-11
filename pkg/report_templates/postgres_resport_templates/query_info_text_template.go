package postgres_resport_templates

const QueryInfoTemplate = `
# Query {{.Pos}} : {{.QPS}} QPS
# Time range: From {{.FromTimestamp}} to {{.ToTimestamp}}
# ====================================================================
# Attribute      total     min     max     avg     95%  stddev  median
# ============ ======= ======= ======= ======= ======= ======= =======
# Count        {{.Count}} 
# Exec time    {{.Duration.Total}} {{.Duration.Min}} {{.Duration.Max}} {{.Duration.Avg}} {{.Duration.Percentile95}} {{.Duration.StdDev}} {{.Duration.Median}}
# ====================================================================
# Completion   {{.CompletedInfo}}
# Query_time distribution
# ====================================================================
#   1us  {{.QueryTimeDistribution.Under10us}}
#  10us  {{.QueryTimeDistribution.Over10usUnder100us}}
# 100us  {{.QueryTimeDistribution.Over100usUnder1ms}}
#   1ms  {{.QueryTimeDistribution.Over1msUnder10ms}}
#  10ms  {{.QueryTimeDistribution.Over10msUnder100ms}}
# 100ms  {{.QueryTimeDistribution.Over100msUnder1s}}
#    1s  {{.QueryTimeDistribution.Over1sUnder10s}}
#  10s+  {{.QueryTimeDistribution.Over10s}}
# ====================================================================
# Query
{{.Query}}
`
