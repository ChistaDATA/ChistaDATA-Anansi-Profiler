package types

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/report_templates"
	log "github.com/sirupsen/logrus"
	"os"
	"sort"
	"text/template"
)

type QueryStringQueryInfoMap map[string]*QueryInfo

func InitQueryInfoMap() QueryStringQueryInfoMap {
	return map[string]*QueryInfo{}
}

func (queryInfoMap *QueryStringQueryInfoMap) Add(query *Query) {
	singleQueryInfo, ok := (*queryInfoMap)[query.GetTransformedQuery()]
	if !ok {
		singleQueryInfo = InitQueryInfo(query.Query)
		(*queryInfoMap)[query.GetTransformedQuery()] = singleQueryInfo
	}
	singleQueryInfo.Add(query)
}

func (queryInfoMap *QueryStringQueryInfoMap) GenerateReportByDuration(config *CliConfig) {

	recordTemplate := report_templates.TopQueryRecord
	accumulatedTemplate := report_templates.AccumulatedInfoTemplate
	topQueriesTemplate := report_templates.TopQueriesTemplate
	queryInfoTemplate := report_templates.QueryInfoTemplate
	outfile := "output.txt"

	if config.ReportType == "md" {
		outfile = "output.md"
		recordTemplate = report_templates.TopQueryMDRecord
		accumulatedTemplate = report_templates.AccumulatedInfoMDTemplate
		topQueriesTemplate = report_templates.TopQueriesMDTemplate
		queryInfoTemplate = report_templates.QueryInfoMDTemplate
	}

	sortedQueryInfos := queryInfoMap.sortQueryInfoByDuration()

	var queryInfoTemplateInputs []QueryInfoTemplateInput
	for i := 0; i < len(sortedQueryInfos) && i < config.TopQueryCount; i++ {
		queryInfoTemplateInputs = append(queryInfoTemplateInputs, InitQueryInfoTemplateInput(i, sortedQueryInfos[i]))
	}

	accumulatedInfoTemplateInput := InitAccumulatedInfoTemplateInput(sortedQueryInfos, config.Files)

	var topQueriesRecords []TopQueriesTemplateInputRecord

	for i := 0; i < len(queryInfoTemplateInputs); i++ {
		topQueriesRecords = append(topQueriesRecords, InitTopQueriesTemplateInputRecord(sortedQueryInfos[i], &queryInfoTemplateInputs[i], accumulatedInfoTemplateInput.TotalDuration))
	}

	topQueriesTemplateInput := InitTopQueriesTemplateInput(topQueriesRecords, recordTemplate)

	temp, err := template.New("AccumulatedInfoTemplate").Option("missingkey=error").Parse(accumulatedTemplate)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	var bf bytes.Buffer
	err = temp.Execute(&bf, accumulatedInfoTemplateInput)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	//log.Infoln(bf.String())

	temp, err = template.New("TopQueriesTemplate").Option("missingkey=error").Parse(topQueriesTemplate)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	err = temp.Execute(&bf, topQueriesTemplateInput)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	//log.Infoln(bf.String())
	temp, err = template.New("queryInfoTemplate").Option("missingkey=error").Parse(queryInfoTemplate)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
	for i := 0; i < len(queryInfoTemplateInputs); i++ {
		err = temp.Execute(&bf, queryInfoTemplateInputs[i])
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		log.Infoln(bf.String())
	}

	f, err := os.Create(outfile)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(bf.String())
	w.Flush()
}

func (queryInfoMap *QueryStringQueryInfoMap) sortQueryInfoByDuration() []*QueryInfo {
	var sortedQueryInfos []*QueryInfo
	for _, queryInfo := range *queryInfoMap {
		sortedQueryInfos = append(sortedQueryInfos, queryInfo)
	}
	sort.Sort(byDuration(sortedQueryInfos))
	return sortedQueryInfos
}

// byDuration implements sort.Interface based on the GetMaxDuration.
type byDuration []*QueryInfo

func (a byDuration) Len() int           { return len(a) }
func (a byDuration) Less(i, j int) bool { return a[i].GetMaxDuration() > a[j].GetMaxDuration() }
func (a byDuration) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
