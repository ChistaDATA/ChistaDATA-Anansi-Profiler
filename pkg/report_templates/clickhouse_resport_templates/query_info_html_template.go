package clickhouse_resport_templates

const QueryInfoHTMLTemplate = `
    <section style="background: #FFFFFF;padding: 10px">
      <h3>Query {{.Pos}}</h3>
      <div class="query">{{.Query}}</div>
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
	<script type="text/javascript">
			google.charts.load("current", {packages:["corechart"]});
			google.charts.setOnLoadCallback(drawChart);
			function drawChart() {
				var data = google.visualization.arrayToDataTable([
					["Time", "Count", { role: "style" } ],
					["1us", {{.QueryTimeDistribution.TimeDistNumber.Under10us}}, "#1A74A6"],
					["10us", {{.QueryTimeDistribution.TimeDistNumber.Over10usUnder100us}}, "#1A74A6"],
					["100us", {{.QueryTimeDistribution.TimeDistNumber.Over100usUnder1ms}}, "#1A74A6"],
					["1ms", {{.QueryTimeDistribution.TimeDistNumber.Over1msUnder10ms}}, "#1A74A6"],
					["10ms", {{.QueryTimeDistribution.TimeDistNumber.Under10us}}, "#1A74A6"],
					["100ms", {{.QueryTimeDistribution.TimeDistNumber.Over100msUnder1s}}, "#1A74A6"],
					["1s", {{.QueryTimeDistribution.TimeDistNumber.Over1sUnder10s}}, "#1A74A6"],
					["10s+", {{.QueryTimeDistribution.TimeDistNumber.Over10s}}, "#1A74A6"]
				]);

				var view = new google.visualization.DataView(data);
				view.setColumns([0, 1,
					{ calc: "stringify",
						sourceColumn: 1,
						type: "string",
						role: "annotation" },
					2]);

				var options = {
					title: "Query Time Distribution",
					width: 600,
					height: 400,
					bar: {groupWidth: "50%"},
					legend: { position: "none" },
				};
				var chart = new google.visualization.BarChart(document.getElementById("barchart_values_{{.Pos}}"));
				chart.draw(view, options);
			}
		</script>
		<div id="barchart_values_{{.Pos}}" style="width: 900px; height: 400px;"></div>

    </section>

`
