package clickhouse_resport_templates

const AccumulatedInfoMDTemplate = `
# Profiler output
* Current date: {{.CurrentDate}}
* Hostname: {{.Hostname}}
* Files:
{{.Files}}
* Query
	* Overall: {{.TotalQueryCount}}
	* Unique: {{.TotalUniqueQueryCount}}
	* QPS: {{.TotalQPS}}
* Time range 
	* From: {{.FromTimestamp}}
	* To: {{.ToTimestamp}}

| Attribute | total | min | max | avg | 95% | stddev | median |
|-----------|-------|-----|-----|-----|-----|--------|--------|
| Exec time | {{.Duration.Total}} | {{.Duration.Min}} | {{.Duration.Max}} | {{.Duration.Avg}} | {{.Duration.Percentile95}} | {{.Duration.StdDev}} | {{.Duration.Median}} |
| Rows read | {{.ReadRows.Total}} | {{.ReadRows.Min}} | {{.ReadRows.Max}} | {{.ReadRows.Avg}} | {{.ReadRows.Percentile95}} | {{.ReadRows.StdDev}} | {{.ReadRows.Median}} |
| Bytes read | {{.ReadBytes.Total}} | {{.ReadBytes.Min}} | {{.ReadBytes.Max}} | {{.ReadBytes.Avg}} | {{.ReadBytes.Percentile95}} | {{.ReadBytes.StdDev}} | {{.ReadBytes.Median}} |
| Peak Memory | |{{.PeakMemoryUsage.Min}} | {{.PeakMemoryUsage.Max}} | {{.PeakMemoryUsage.Avg}} | {{.PeakMemoryUsage.Percentile95}} | {{.PeakMemoryUsage.StdDev}} | {{.PeakMemoryUsage.Median}} |

* Query Type
	* Selects:{{.QueryTypeCount.Select}}
	* Inserts:{{.QueryTypeCount.Insert}}
	* Updates:{{.QueryTypeCount.Update}}
	* Deletes:{{.QueryTypeCount.Delete}}
`
