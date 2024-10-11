package clickhouse_resport_templates

const TopQueriesHTMLTemplate = `
<h1>Top Queries</h1>
<table>
    <tr>
        <td>Rank</td>
        <td>Response time</td>
        <td>Calls</td>
        <td>R/Call</td>
        <td>Query</td>
        <td></td>
    </tr>
    <tr>
    {{range $record:=.}}
    <td>{{$record.Pos}}</td>
    <td>{{$record.TotalDuration}}</td>
    <td>{{$record.TotalDurationPercentage}}</td>
    <td>{{$record.Count}}</td>
    <td>{{$record.ResponseTimePerCall}}</td>
    <td>{{$record.Query}}</td>
    {{end}}     
    </tr>
</table>
`

const TopQueryHTMLRecord = `<p>| {{.Pos}} | {{.TotalDuration}} {{.TotalDurationPercentage}} |  {{.Count}} | {{.ResponseTimePerCall}} | {{.Query}} |</p>`
