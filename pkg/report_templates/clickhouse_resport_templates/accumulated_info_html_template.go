package clickhouse_resport_templates

const AccumulatedInfoHTMLTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <title>Anansi-report</title>
    <style>
        *, *::before, *::after {
            box-sizing: border-box;
        }

        * {
            margin: 0;
        }

        body {
            line-height: 1.5;
            -webkit-font-smoothing: antialiased;
        }

        table {
            background: #FFFFFF;
            width: 100%;
            border-collapse: collapse;
        }

        thead {
            background: #F5F6FA;
            font-weight: bold;
            border-bottom: 1px solid black;
            height: 40px;
            top: 0;
            margin: 0;
            text-align: left;
        }

        tr {
            height: 14px;
            border-bottom: 1px solid black;
            text-align: left;
        }

        td {
            padding: 10px;
            text-align: left;
        }

        html {
            display: flex;
        }

        header {
            font-family: monospace;
            color: #0141A1;
        }

        section {
            margin: 20px;
            background: #F7F7F7;
        }

        body {
            height: 100vh;
            width: 100vw;
            display: flex;
            flex-direction: column;
            font-family: sans-serif;
            background: #F1F1F1;
            overflow-x: auto;
        }

        .card {
            height: 109px;
            width: 353px;
            background: #FFFFFF;
            padding: 20px 24px;
            border-radius: 8px;
            display: flex;
            justify-content: space-between;
        }
    </style>
    <script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
    <script type="text/javascript">
        google.charts.load('current', {'packages': ['corechart']});
        google.charts.setOnLoadCallback(drawChart);

        function drawChart() {

            var data = google.visualization.arrayToDataTable([
                ['Query Type', 'Count'],
                ['Select', 4943],
                ['Insert', 1],
                ['Delete', 0],
                ['Update', 0],
            ]);

            var options = {
                title: 'Query types'
            };

            var chart = new google.visualization.PieChart(document.getElementById('piechart'));

            chart.draw(data, options);
        }
    </script>

</head>

<body>
<header style="background: #FFFFFF;padding-left: 20px;"><h1>Anansi Profiler</h1></header>
<section style="margin: 0px;padding: 20px;font-weight: bold">Log analysis and optimization report</section>

<section
        style="display: flex;justify-content:space-between;flex-wrap: wrap; font-size: small; gap: 20px; padding: 20px">
    <div style="display: flex;flex-direction: column;gap:10px">
        <div>Execution Timestamp<br>{{.CurrentDate}}</div>
        <div>Hostname<br>{{.Hostname}}</div>
    </div>
    <div style="display: flex;flex-direction: column;gap:10px">
        <div>Files<br>{{.Files}}</div>
        <div>Analysis Period<br>From: {{.FromTimestamp}} - To: {{.ToTimestamp}}</div>
    </div>
</section>

<section style="display: flex; flex: 1;background: none">
    <div style="display: flex;flex-wrap: wrap; gap: 32px">
        <div class="card">
            <div>Total Queries</div>
            <div>
                <div style="font-size: x-large;color: #0141A1;font-weight: bold">{{.TotalQueryCount}}</div>
                <div style="font-size: small">Unique: {{.TotalUniqueQueryCount}}</div>
            </div>
        </div>
        <div class="card">
            <div>Execution Time</div>
            <div>
                <div style="font-size: x-large;color: #0141A1;font-weight: bold">{{.Duration.Total}}</div>
                <div style="font-size: small">QPS: {{.TotalQPS}}</div>
            </div>
        </div>
        <div class="card">
            <div>Data Read</div>
            <div>
                <div style="font-size: x-large;color: #0141A1;font-weight: bold">{{.ReadBytes.Total}}</div>
                <div style="font-size: small">{{.ReadRows.Total}} Rows</div>
            </div>
        </div>
    </div>
</section>
<section>
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
    </section>
    <section style="display: flex; justify-content: space-around">
      <div id="piechart" style="width:400px; height: 400px;"></div>
       <ul>
        <li>Selects:{{.QueryTypeCount.Select}}</li>
        <li>Inserts:{{.QueryTypeCount.Insert}}</li>
        <li>Updates:{{.QueryTypeCount.Update}}</li>
        <li>Deletes:{{.QueryTypeCount.Delete}}</li>
      </ul>
    </section>
`
