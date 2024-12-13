package postgres_resport_templates

const QueryInfoTemplate = `
# Query {{.Pos}} : {{.QPS}} QPS
# Time range: From {{.FromTimestamp}} to {{.ToTimestamp}}
# ====================================================================
# Attribute      total     min     max     avg     95%  stddev  median
# ============ ======= ======= ======= ======= ======= ======= =======
# Count        {{.Count}} 
# Exec time    {{.Duration.Total}} {{.Duration.Min}} {{.Duration.Max}} {{.Duration.Avg}} {{.Duration.Percentile95}} {{.Duration.StdDev}} {{.Duration.Median}}
# Peak Memory          {{.PeakMemoryUsage.Min}} {{.PeakMemoryUsage.Max}} {{.PeakMemoryUsage.Avg}} {{.PeakMemoryUsage.Percentile95}} {{.PeakMemoryUsage.StdDev}} {{.PeakMemoryUsage.Median}}
# ====================================================================
# Databases    {{.DatabaseInfo}}
# Hosts        {{.HostInfo}}
# Users        {{.UserInfo}}
# Completion   {{.CompletedInfo}}
# Errors       {{.ErrorInfo}}
# Query_time distribution
# ====================================================================
#   1us  {{.QueryTimeDistribution.TimeDistString.Under10us}}
#  10us  {{.QueryTimeDistribution.TimeDistString.Over10usUnder100us}}
# 100us  {{.QueryTimeDistribution.TimeDistString.Over100usUnder1ms}}
#   1ms  {{.QueryTimeDistribution.TimeDistString.Over1msUnder10ms}}
#  10ms  {{.QueryTimeDistribution.TimeDistString.Over10msUnder100ms}}
# 100ms  {{.QueryTimeDistribution.TimeDistString.Over100msUnder1s}}
#    1s  {{.QueryTimeDistribution.TimeDistString.Over1sUnder10s}}
#  10s+  {{.QueryTimeDistribution.TimeDistString.Over10s}}
# ====================================================================
# Query
{{.Query}}
`
