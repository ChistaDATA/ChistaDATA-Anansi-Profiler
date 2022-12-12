package services

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/parsers/clickhouse"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/report_templates"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"text/template"
)

type ReportGenerator struct {
	Config              *stucts.CliConfig
	InfoCorpus          *stucts.InfoCorpus
	ReportTemplates     ReportTemplates
	OutputFileExtension string
}

type ReportTemplates struct {
	TopQueryRecordTemplate *template.Template
	AccumulatedTemplate    *template.Template
	TopQueriesTemplate     *template.Template
	QueryInfoTemplate      *template.Template
}

func InitReportGenerator(cliConfig *stucts.CliConfig, infoCorpus *stucts.InfoCorpus) ReportGenerator {
	reportGenerator := ReportGenerator{
		Config:     cliConfig,
		InfoCorpus: infoCorpus,
	}

	if cliConfig.ReportType == stucts.ReportTypeText {
		reportGenerator.ReportTemplates = initReportTemplates(report_templates.TopQueryRecord, report_templates.AccumulatedInfoTemplate, report_templates.TopQueriesTemplate, report_templates.QueryInfoTemplate)
		reportGenerator.OutputFileExtension = "txt"
	} else {
		reportGenerator.ReportTemplates = initReportTemplates(report_templates.TopQueryMDRecord, report_templates.AccumulatedInfoMDTemplate, report_templates.TopQueriesMDTemplate, report_templates.QueryInfoMDTemplate)
		reportGenerator.OutputFileExtension = "md"
	}

	return reportGenerator
}

func initReportTemplates(topQueryRecordTemplate string, accumulatedTemplate string, topQueriesTemplate string, queryInfoTemplate string) ReportTemplates {
	var err error
	reportTemplates := ReportTemplates{}
	reportTemplates.TopQueryRecordTemplate, err = template.New("TopQueryRecordTemplate").Option("missingkey=error").Parse(topQueryRecordTemplate)
	if err != nil {
		panic(err)
	}
	reportTemplates.AccumulatedTemplate, err = template.New("AccumulatedTemplate").Option("missingkey=error").Parse(accumulatedTemplate)
	if err != nil {
		panic(err)
	}
	reportTemplates.TopQueriesTemplate, err = template.New("TopQueriesTemplate").Option("missingkey=error").Parse(topQueriesTemplate)
	if err != nil {
		panic(err)
	}
	reportTemplates.QueryInfoTemplate, err = template.New("QueryInfoTemplate").Option("missingkey=error").Parse(queryInfoTemplate)
	if err != nil {
		panic(err)
	}
	return reportTemplates
}

// GenerateReport Generates report using list of queries
func (reportGenerator ReportGenerator) GenerateReport() {
	// There can be queries that are similar, so making a list of similar queries
	simplifiedQueryInfoList := stucts.InitSimilarQueryInfoList()
	for _, query := range reportGenerator.InfoCorpus.Queries.GetList() {
		simplifiedQueryInfoList.Add(query)
	}

	// making accumulated info of all queries
	accumulatedInfoTemplateInput := stucts.InitAccumulatedInfoTemplateInput(simplifiedQueryInfoList, reportGenerator.Config.FilePaths)

	// There are certain queries that are not important for profiler
	// like create, insert queries and queries with count less than minimum count
	var keysToRemove []string
	for s, info := range simplifiedQueryInfoList {
		if shouldDiscardQuery(info.Query) || info.Count < reportGenerator.Config.MinimumQueryCallCount {
			keysToRemove = append(keysToRemove, s)
		}
	}
	for _, s := range keysToRemove {
		simplifiedQueryInfoList.Remove(s)
	}

	// Sorting remaining values in simplifiedQueryInfoList, and converting the map to an actual list
	sortedSimilarQueryInfos := reportGenerator.sortSimplifiedQueryInfoList(simplifiedQueryInfoList, accumulatedInfoTemplateInput.TotalDuration)
	var queryInfoTemplateInputs []stucts.QueryInfoTemplateInput

	// Generating Input structs for all templates

	for i := 0; i < len(sortedSimilarQueryInfos) && i < reportGenerator.Config.TopQueryCount; i++ {
		queryInfoTemplateInputs = append(queryInfoTemplateInputs, stucts.InitQueryInfoTemplateInput(i, sortedSimilarQueryInfos[i], accumulatedInfoTemplateInput.TotalDuration))
	}

	var topQueriesRecords []stucts.TopQueriesTemplateInputRecord
	for i := 0; i < len(queryInfoTemplateInputs); i++ {
		topQueriesRecords = append(topQueriesRecords, stucts.InitTopQueriesTemplateInputRecord(sortedSimilarQueryInfos[i], &queryInfoTemplateInputs[i], accumulatedInfoTemplateInput.TotalDuration))
	}

	topQueriesTemplateInput := stucts.InitTopQueriesTemplateInput(topQueriesRecords, reportGenerator.ReportTemplates.TopQueryRecordTemplate)

	// Generating output from templates, to a buffer

	var err error
	var bf bytes.Buffer
	err = reportGenerator.ReportTemplates.AccumulatedTemplate.Execute(&bf, accumulatedInfoTemplateInput)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	err = reportGenerator.ReportTemplates.TopQueriesTemplate.Execute(&bf, topQueriesTemplateInput)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	for i := 0; i < len(queryInfoTemplateInputs); i++ {
		err = reportGenerator.ReportTemplates.QueryInfoTemplate.Execute(&bf, queryInfoTemplateInputs[i])
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		log.Infoln(bf.String())
	}

	// Persisting the output to a file, from buffer

	f, err := os.Create("output." + reportGenerator.OutputFileExtension)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(bf.String())
	w.Flush()
}

func (reportGenerator ReportGenerator) sortSimplifiedQueryInfoList(list stucts.SimilarQueryInfoList, totalDuration float64) []*stucts.SimilarQueryInfo {
	return list.Sort(reportGenerator.Config.SortField, reportGenerator.Config.SortFieldOperation, reportGenerator.Config.SortOrder, totalDuration)
}

func shouldDiscardQuery(query string) bool {
	if clickhouse.QueriesToDiscardRegEx.MatchString(strings.TrimSpace(query)) {
		return true
	}
	return false
}
