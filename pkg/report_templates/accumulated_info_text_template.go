package report_templates

const AccumulatedInfoTemplate = `
# Current date: {{.CurrentDate}}
# Hostname: {{.Hostname}}
# Files:
{{.Files}}
# Overall: {{.TotalQueryCount}}, Unique: {{.TotalUniqueQueryCount}} , QPS: {{.TotalQPS}}
# Time range: {{.FromTimestamp}} to {{.ToTimestamp}}
# Attribute          total     min     max     avg     95%  stddev  median
# ============     ======= ======= ======= ======= ======= ======= =======
# Exec time        {{.Duration.Total}} {{.Duration.Min}} {{.Duration.Max}} {{.Duration.Avg}} {{.Duration.Percentile95}} {{.Duration.StdDev}} {{.Duration.Median}}
# Rows read        {{.ReadRows.Total}} {{.ReadRows.Min}} {{.ReadRows.Max}} {{.ReadRows.Avg}} {{.ReadRows.Percentile95}} {{.ReadRows.StdDev}} {{.ReadRows.Median}}
# Bytes read       {{.ReadBytes.Total}} {{.ReadBytes.Min}} {{.ReadBytes.Max}} {{.ReadBytes.Avg}} {{.ReadBytes.Percentile95}} {{.ReadBytes.StdDev}} {{.ReadBytes.Median}}
# Peak Memory              {{.PeakMemoryUsage.Min}} {{.PeakMemoryUsage.Max}} {{.PeakMemoryUsage.Avg}} {{.PeakMemoryUsage.Percentile95}} {{.PeakMemoryUsage.StdDev}} {{.PeakMemoryUsage.Median}}
`
