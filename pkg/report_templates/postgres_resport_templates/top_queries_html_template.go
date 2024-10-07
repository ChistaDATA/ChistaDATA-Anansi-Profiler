package postgres_resport_templates

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
        <td>{{.Records}}</td>
    </tr>
</table>
`

const TopQueryHTMLRecord = `<p>| {{.Pos}} | {{.TotalDuration}} {{.TotalDurationPercentage}} |  {{.Count}} | {{.ResponseTimePerCall}} | {{.Query}} |</p>`
