package postgres_resport_templates

const TopQueriesHTMLTemplate = `
<section>
<h3>Top Queries</h3>
  <table>
	<thead>
		<tr>
		  <td>Rank</td>
		  <td>Response time</td>
		  <td>Response time(%)</td>
		  <td>Calls</td>
		  <td>R/Call</td>
		  <td>Query</td>
		</tr>
	</thead>
	<tbody>
		{{range $record:=.}}
		<tr>
		  <td>{{$record.Pos}}</td>
		  <td>{{$record.TotalDuration}}</td>
		  <td>{{$record.TotalDurationPercentage}}</td>
		  <td>{{$record.Count}}</td>
		  <td>{{$record.ResponseTimePerCall}}</td>
		  <td>{{$record.Query}}</td>
		</tr>
		{{end}}
	</tbody>
  </table>
</section>
`

const TopQueryHTMLRecord = `<p>| {{.Pos}} | {{.TotalDuration}} {{.TotalDurationPercentage}} |  {{.Count}} | {{.ResponseTimePerCall}} | {{.Query}} |</p>`
