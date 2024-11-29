package clickhouse_resport_templates

const QueryInfoHTMLTemplate = `
    <section>
      <h1>Query {{.Pos}}</h1>
      <p>{{.Query}}</p>
      <div>QPS: {{.QPS}}</div>
      <p>Time range</p>
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
			  <th>Count</th>
			  <td>{{.Count}}</td>
			  <td></td>
			  <td></td>
			  <td></td>
			  <td></td>
			  <td></td>
			  <td></td>
			</tr>
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
			  <td></td>
			  <td>{{.PeakMemoryUsage.Min}}</td>
			  <td>{{.PeakMemoryUsage.Max}}</td>
			  <td>{{.PeakMemoryUsage.Avg}}</td>
			  <td>{{.PeakMemoryUsage.Percentile95}}</td>
			  <td>{{.PeakMemoryUsage.StdDev}}</td>
			  <td>{{.PeakMemoryUsage.Median}}</td>
			</tr>
		</tbody>
      </table>

      <ul>
        <li>Databases: {{.DatabaseInfo}}</li>
        <li>Hosts: {{.HostInfo}}</li>
        <li>Users: {{.UserInfo}}</li>
        <li>Completion: {{.CompletedInfo}}</li>
        <li>Errors: {{.ErrorInfo}}</li>
      </ul>

      <h1>Query_time distribution</h1>
      <table>
        <tr>
          <td>1us</td>
          <td>{{.QueryTimeDistribution.Under10us}}</td>
        </tr>
        <tr>
          <td>10us</td>
          <td>{{.QueryTimeDistribution.Over10usUnder100us}}</td>
        </tr>
        <tr>
          <td>100us</td>
          <td>{{.QueryTimeDistribution.Over100usUnder1ms}}</td>
        </tr>
        <tr>
          <td>1ms</td>
          <td>{{.QueryTimeDistribution.Over1msUnder10ms}}</td>
        </tr>
        <tr>
          <td>10ms</td>
          <td>{{.QueryTimeDistribution.Under10us}}</td>
        </tr>
        <tr>
          <td>100ms</td>
          <td>{{.QueryTimeDistribution.Over100msUnder1s}}</td>
        </tr>
        <tr>
          <td>1s {{.QueryTimeDistribution.Over1sUnder10s}}</td>
          <td>{{.QueryTimeDistribution.Over1sUnder10s}}</td>
        </tr>
        <tr>
          <td>10s+ {{.QueryTimeDistribution.Over10s}}</td>
          <td>{{.QueryTimeDistribution.Over10s}}</td>
        </tr>
      </table>
    </section>
  </body>
</html>
`
