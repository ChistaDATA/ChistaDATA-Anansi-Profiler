package postgres_resport_templates

const AccumulatedInfoHTMLTemplate = `
<style>
table {
  border-collapse: collapse;
}
</style>
<h1>Profiler output</h1>
<div>Current date: {{.CurrentDate}}</div>
<div>Hostname: {{.Hostname}}</div>
<div>Files:{{.Files}}</div>
<p>Query</p>
<ul>
	<li>Overall: {{.TotalQueryCount}}</li>
	<li>Unique: {{.TotalUniqueQueryCount}}</li>
	<li>QPS: {{.TotalQPS}}</li>
</ul>

<p>Time range </p>
	<li>From: {{.FromTimestamp}}</li>
	<li>To: {{.ToTimestamp}}</li>

<table>
	<thead>
		<tr>
			<th>Attribute</th>
			<th>total</th>
			<th>min</th>
			<th>max</th>
			<th>avg</th>
			<th>95%</th>
			<th>stddev</th>
			<th>median</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<th>Exec time</th>
			<td>{{.Duration.Total}}</td>
			<td>{{.Duration.Min}}</td>
			<td>{{.Duration.Max}}</td>
			<td>{{.Duration.Avg}}</td>
			<td>{{.Duration.Percentile95}}</td>
			<td>{{.Duration.StdDev}}</td>
			<td>{{.Duration.Median}}</td>
		</tr>
		<tr>
			<th>Rows read</th>
			<td>{{.ReadRows.Total}}</td>
			<td>{{.ReadRows.Min}}</td>
			<td>{{.ReadRows.Max}}</td>
			<td>{{.ReadRows.Avg}}</td>
			<td>{{.ReadRows.Percentile95}}</td>
			<td>{{.ReadRows.StdDev}}</td>
			<td>{{.ReadRows.Median}}</td>
		</tr>
		<tr>
			<th>Bytes read</th>
			<td>{{.ReadBytes.Total}}</td>
			<td>{{.ReadBytes.Min}}</td>
			<td>{{.ReadBytes.Max}}</td>
			<td>{{.ReadBytes.Avg}}</td>
			<td>{{.ReadBytes.Percentile95}}</td>
			<td>{{.ReadBytes.StdDev}}</td>
			<td>{{.ReadBytes.Median}}</td>
		</tr>
		<tr>
			<th>Peak Memory</th>
			<td>-</td>
			<td>{{.PeakMemoryUsage.Min}}</td>
			<td>{{.PeakMemoryUsage.Max}}</td>
			<td>{{.PeakMemoryUsage.Avg}}</td>
			<td>{{.PeakMemoryUsage.Percentile95}}</td>
			<td>{{.PeakMemoryUsage.StdDev}}</td>
			<td>{{.PeakMemoryUsage.Median}}</td>
		</tr>
	</tbody>
</table>
`
