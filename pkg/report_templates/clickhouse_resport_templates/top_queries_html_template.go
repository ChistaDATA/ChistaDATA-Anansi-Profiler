package clickhouse_resport_templates

const TopQueriesHTMLTemplate = `
<h1>Top Queries</h1>
<section>
  <table>
    <tr>
      <td>Rank</td>
      <td>Response time</td>
      <td>Response time(%)</td>
      <td>Calls</td>
      <td>R/Call</td>
      <td>Query</td>
      <td></td>
    </tr>
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
  </table>
</section>
`

const TopQueryHTMLRecord = `<p>| {{.Pos}} | {{.TotalDuration}} {{.TotalDurationPercentage}} |  {{.Count}} | {{.ResponseTimePerCall}} | {{.Query}} |</p>`
