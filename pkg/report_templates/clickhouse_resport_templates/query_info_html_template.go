package clickhouse_resport_templates

const QueryInfoHTMLTemplate = `
<section style="background: #ffffff; padding: 10px">
	<div style="border-bottom: 1px solid #e3e3e3; padding-bottom: 15px">
		<div style="display: flex; justify-content: space-between">
			<h3>Query Rank: <span style="color: #0141a1">{{.Pos}}</span></h3>
			<h3>QPS: <span style="color: #0141a1">{{.QPS}}</span></h3>
		</div>
		<div class="query">{{.Query}}</div>
	</div>

<script>
	function changeTab_{{.Pos}}(event, tabId) {
		document.getElementById("Overview_tab_{{.Pos}}").classList.remove('active');
		document.getElementById("Execution_Time_tab_{{.Pos}}").classList.remove('active');
		document.getElementById("Rows_Read_tab_{{.Pos}}").classList.remove('active');
		document.getElementById("Bytes_Read_tab_{{.Pos}}").classList.remove('active');
		document.getElementById("Peak_Memory_tab_{{.Pos}}").classList.remove('active');

		document.getElementById("Overview_{{.Pos}}").classList.remove('active');
		document.getElementById("Execution_Time_{{.Pos}}").classList.remove('active');
		document.getElementById("Rows_Read_{{.Pos}}").classList.remove('active');
		document.getElementById("Bytes_Read_{{.Pos}}").classList.remove('active');
		document.getElementById("Peak_Memory_{{.Pos}}").classList.remove('active');

		event.currentTarget.classList.add('active');
		document.getElementById(tabId).classList.add('active');
	}
</script>
<div class="tabs" style="margin:10px 0px">
	<button class="tab" id="Overview_tab_{{.Pos}}" onclick="changeTab_{{.Pos}}(event, 'Overview_{{.Pos}}')">Overview</button>
	<button class="tab" id="Execution_Time_tab_{{.Pos}}" onclick="changeTab_{{.Pos}}(event, 'Execution_Time_{{.Pos}}')">Execution Time</button>
	<button class="tab" id="Rows_Read_tab_{{.Pos}}" onclick="changeTab_{{.Pos}}(event, 'Rows_Read_{{.Pos}}')">Rows Read</button>
	<button class="tab" id="Bytes_Read_tab_{{.Pos}}" onclick="changeTab_{{.Pos}}(event, 'Bytes_Read_{{.Pos}}')">Bytes Read</button>
	<button class="tab" id="Peak_Memory_tab_{{.Pos}}" onclick="changeTab_{{.Pos}}(event, 'Peak_Memory_{{.Pos}}')">Peak Memory</button>
</div>

<div class="tab-content">
	<div id="Overview_{{.Pos}}" class="tab-pane active">
		<div>
			<div>
				<h3>Query Timeline</h3>
				<div class="query">
					<p>Time range</p>
					<div>From: {{.FromTimestamp}}</div>
					<div>To: {{.ToTimestamp}}</div>
				</div>
			</div>
			<div>
				<h3>Summary Statistics</h3>
				<div>
					<div>Count</div>
					<div>{{.Count}}</div>

					<ul>
						<li>Databases: {{.DatabaseInfo}}</li>
						<li>Hosts: {{.HostInfo}}</li>
						<li>Users: {{.UserInfo}}</li>
						<li>Completion: {{.CompletedInfo}}</li>
						<li>Errors: {{.ErrorInfo}}</li>
					</ul>
				</div>
			</div>
		</div>
	</div>
	<div id="Execution_Time_{{.Pos}}" class="tab-pane">
		<table>
			<thead>
				<tr>
					<th>Metrics</th>
					<th>Value</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<th>Total</th>
					<td>{{.Duration.Total}}</td>
				</tr>
				<tr>
					<th>Min</th>
					<td>{{.Duration.Min}}</td>
				</tr>
				<tr>
					<th>Max</th>
					<td>{{.Duration.Max}}</td>
				</tr>
				<tr>
					<th>Avg</th>
					<td>{{.Duration.Avg}}</td>
				</tr>
				<tr>
					<th>95%</th>
					<td>{{.Duration.Percentile95}}</td>
				</tr>
				<tr>
					<th>stddev</th>
					<td>{{.Duration.StdDev}}</td>
				</tr>
				<tr>
					<th>Median</th>
					<td>{{.Duration.Median}}</td>
				</tr>
			</tbody>
		</table>
	</div>
	<div id="Rows_Read_{{.Pos}}" class="tab-pane">
			<table>
			<thead>
				<tr>
					<th>Metrics</th>
					<th>Value</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<th>Total</th>
					<td>{{.ReadRows.Total}}</td>
				</tr>
				<tr>
					<th>Min</th>
					<td>{{.ReadRows.Min}}</td>
				</tr>
				<tr>
					<th>Max</th>
					<td>{{.ReadRows.Max}}</td>
				</tr>
				<tr>
					<th>Avg</th>
					<td>{{.ReadRows.Avg}}</td>
				</tr>
				<tr>
					<th>95%</th>
					<td>{{.ReadRows.Percentile95}}</td>
				</tr>
				<tr>
					<th>stddev</th>
					<td>{{.ReadRows.StdDev}}</td>
				</tr>
				<tr>
					<th>Median</th>
					<td>{{.ReadRows.Median}}</td>
				</tr>
			</tbody>
		</table>
	</div>
	<div id="Bytes_Read_{{.Pos}}" class="tab-pane">
			<table>
			<thead>
				<tr>
					<th>Metrics</th>
					<th>Value</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<th>Total</th>
					<td>{{.ReadBytes.Total}}</td>
				</tr>
				<tr>
					<th>Min</th>
					<td>{{.ReadBytes.Min}}</td>
				</tr>
				<tr>
					<th>Max</th>
					<td>{{.ReadBytes.Max}}</td>
				</tr>
				<tr>
					<th>Avg</th>
					<td>{{.ReadBytes.Avg}}</td>
				</tr>
				<tr>
					<th>95%</th>
					<td>{{.ReadBytes.Percentile95}}</td>
				</tr>
				<tr>
					<th>stddev</th>
					<td>{{.ReadBytes.StdDev}}</td>
				</tr>
				<tr>
					<th>Median</th>
					<td>{{.ReadBytes.Median}}</td>
				</tr>
			</tbody>
		</table>
	</div>
	<div id="Peak_Memory_{{.Pos}}" class="tab-pane">
		<table>
			<thead>
				<tr>
					<th>Metrics</th>
					<th>Value</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<th>Total</th>
					<td>-</td>
				</tr>
				<tr>
					<th>Min</th>
					<td>{{.PeakMemoryUsage.Min}}</td>
				</tr>
				<tr>
					<th>Max</th>
					<td>{{.PeakMemoryUsage.Max}}</td>
				</tr>
				<tr>
					<th>Avg</th>
					<td>{{.PeakMemoryUsage.Avg}}</td>
				</tr>
				<tr>
					<th>95%</th>
					<td>{{.PeakMemoryUsage.Percentile95}}</td>
				</tr>
				<tr>
					<th>stddev</th>
					<td>{{.PeakMemoryUsage.StdDev}}</td>
				</tr>
				<tr>
					<th>Median</th>
					<td>{{.PeakMemoryUsage.Median}}</td>
				</tr>
			</tbody>
		</table>
	</div>
</div>

	<script type="text/javascript">
		google.charts.load("current", { packages: ["corechart"] });
		google.charts.setOnLoadCallback(drawChart);
		function drawChart() {
			var data = google.visualization.arrayToDataTable([
				["Time", "Count", { role: "style" }],
				["1us", {{.QueryTimeDistribution.TimeDistNumber.Under10us }}, "#1A74A6"],
		["10us", {{.QueryTimeDistribution.TimeDistNumber.Over10usUnder100us }}, "#1A74A6"],
			["100us", {{.QueryTimeDistribution.TimeDistNumber.Over100usUnder1ms }}, "#1A74A6"],
			["1ms", {{.QueryTimeDistribution.TimeDistNumber.Over1msUnder10ms }}, "#1A74A6"],
			["10ms", {{.QueryTimeDistribution.TimeDistNumber.Under10us }}, "#1A74A6"],
			["100ms", {{.QueryTimeDistribution.TimeDistNumber.Over100msUnder1s }}, "#1A74A6"],
			["1s", {{.QueryTimeDistribution.TimeDistNumber.Over1sUnder10s }}, "#1A74A6"],
			["10s+", {{.QueryTimeDistribution.TimeDistNumber.Over10s }}, "#1A74A6"]
      	]);

		var view = new google.visualization.DataView(data);
		view.setColumns([0, 1,
			{
				calc: "stringify",
				sourceColumn: 1,
				type: "string",
				role: "annotation"
			},
			2]);

		var options = {
			title: "Query Time Distribution",
			width: 600,
			height: 400,
			bar: { groupWidth: "50%" },
			legend: { position: "none" },
		};
		var chart = new google.visualization.BarChart(document.getElementById("barchart_values_{{.Pos}}"));
		chart.draw(view, options);
      }
	</script>
	<div id="barchart_values_{{.Pos}}" style="width: 900px; height: 400px"></div>
</section>
`
