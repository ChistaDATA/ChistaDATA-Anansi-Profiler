package clickhouse_resport_templates

const QueryInfoHTMLTemplate = `
<h1>Query {{.Pos}}</h1>
<h2>{{.Query}}</h2>
<div> QPS: {{.QPS}} </div>
<p>Time range </p>
	<li>From: {{.FromTimestamp}}</li>
	<li>To: {{.ToTimestamp}}</li>

<table>
    <tr>
        <td>Attribute</td>
        <td>total</td>
        <td>min</td>
        <td>max</td>
        <td>avg</td>
        <td>95%</td>
        <td>stddev</td>
        <td>median</td>
    </tr>
    <tr>
        <td>Count</td>
        <td>{{.Count}}</td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>Exec time</td>
        <td>{{.Duration.Total}}</td>
        <td>{{.Duration.Min}}</td>
        <td>{{.Duration.Max}}</td>
        <td>{{.Duration.Avg}}</td>
        <td>{{.Duration.Percentile95}}</td>
        <td>{{.Duration.StdDev}}</td>
        <td>{{.Duration.Median}}</td>
    </tr>
    <tr>
        <td>Rows read</td>
        <td>{{.ReadRows.Total}}</td>
        <td>{{.ReadRows.Min}}</td>
        <td>{{.ReadRows.Max}}</td>
        <td>{{.ReadRows.Avg}}</td>
        <td>{{.ReadRows.Percentile95}}</td>
        <td>{{.ReadRows.StdDev}}</td>
        <td>{{.ReadRows.Median}}</td>
    </tr>
    <tr>
        <td>Bytes read</td>
        <td>{{.ReadBytes.Total}}</td>
        <td>{{.ReadBytes.Min}}</td>
        <td>{{.ReadBytes.Max}}</td>
        <td>{{.ReadBytes.Avg}}</td>
        <td>{{.ReadBytes.Percentile95}}</td>
        <td>{{.ReadBytes.StdDev}}</td>
        <td>{{.ReadBytes.Median}}</td>
    </tr>
    <tr>
        <td>Peak Memory</td>
        <td></td>
        <td>{{.PeakMemoryUsage.Min}}</td>
        <td>{{.PeakMemoryUsage.Max}}</td>
        <td>{{.PeakMemoryUsage.Avg}}</td>
        <td>{{.PeakMemoryUsage.Percentile95}}</td>
        <td>{{.PeakMemoryUsage.StdDev}}</td>
        <td>{{.PeakMemoryUsage.Median}}</td>
    </tr>
</table>

<ul>
	<li>Databases:    {{.DatabaseInfo}}</li>
	<li>Hosts:        {{.HostInfo}}</li>
	<li>Users:        {{.UserInfo}}</li>
	<li>Completion:   {{.CompletedInfo}}</li>
	<li>Errors:       {{.ErrorInfo}}</li>
</ul>

<h1>Query_time distribution</h1>
<table>
    <tr>
        <td>1us  </td>
		<td>{{.QueryTimeDistribution.Under10us}}</td>
    </tr>
    <tr>
        <td>10us </td>
		<td>{{.QueryTimeDistribution.Over10usUnder100us}}</td>
    </tr>
    <tr>
        <td>100us</td>
		<td>{{.QueryTimeDistribution.Over100usUnder1ms}}</td>
    </tr>
    <tr>
        <td>1ms  </td>
		<td>{{.QueryTimeDistribution.Over1msUnder10ms}}</td>
    </tr>
    <tr>
        <td>10ms </td>
		<td>{{.QueryTimeDistribution.Under10us}}</td>
    </tr>
    <tr>
        <td>100ms</td>
		<td>{{.QueryTimeDistribution.Over100msUnder1s}}</td>
    </tr>
    <tr>
        <td>1s  {{.QueryTimeDistribution.Over1sUnder10s}}</td>
		<td>{{.QueryTimeDistribution.Over1sUnder10s}}</td>
    </tr>
    <tr>
        <td>10s+  {{.QueryTimeDistribution.Over10s}}</td>
		<td>{{.QueryTimeDistribution.Over10s}}</td>
    </tr>
</table>
`
