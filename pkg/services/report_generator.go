package services

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/report_templates/clickhouse_resport_templates"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/report_templates/postgres_resport_templates"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	log "github.com/sirupsen/logrus"
)

// ReportGenerator it is used to generate a report using the given *stucts.Config and *stucts.DBPerfInfoRepository
type ReportGenerator struct {
	Config               *stucts.Config
	DBPerfInfoRepository *stucts.DBPerfInfoRepository
	ReportTemplates      ReportTemplates
	OutputFileExtension  string
	discardQueryRegex    *regexp.Regexp
}

type ReportTemplates struct {
	TopQueryRecordTemplate *template.Template
	AccumulatedTemplate    *template.Template
	TopQueriesTemplate     *template.Template
	QueryInfoTemplate      *template.Template
}

func InitReportGenerator(config *stucts.Config, dBPerfInfoRepository *stucts.DBPerfInfoRepository) ReportGenerator {
	reportGenerator := ReportGenerator{
		Config:               config,
		DBPerfInfoRepository: dBPerfInfoRepository,
		discardQueryRegex:    getDiscardQueryRegex(config.DiscardQueries),
	}

	if config.DatabaseType == stucts.ClickHouseDatabase {
		if config.ReportType == stucts.ReportTypeText {
			reportGenerator.ReportTemplates = initReportTemplates(clickhouse_resport_templates.TopQueryRecord, clickhouse_resport_templates.AccumulatedInfoTemplate, clickhouse_resport_templates.TopQueriesTemplate, clickhouse_resport_templates.QueryInfoTemplate)
			reportGenerator.OutputFileExtension = "txt"
		} else {
			reportGenerator.ReportTemplates = initReportTemplates(clickhouse_resport_templates.TopQueryMDRecord, clickhouse_resport_templates.AccumulatedInfoMDTemplate, clickhouse_resport_templates.TopQueriesMDTemplate, clickhouse_resport_templates.QueryInfoMDTemplate)
			reportGenerator.OutputFileExtension = "md"
		}
	} else {
		if config.ReportType == stucts.ReportTypeText {
			reportGenerator.ReportTemplates = initReportTemplates(postgres_resport_templates.TopQueryRecord, postgres_resport_templates.AccumulatedInfoTemplate, postgres_resport_templates.TopQueriesTemplate, postgres_resport_templates.QueryInfoTemplate)
			reportGenerator.OutputFileExtension = "txt"
		} else {
			reportGenerator.ReportTemplates = initReportTemplates(postgres_resport_templates.TopQueryMDRecord, postgres_resport_templates.AccumulatedInfoMDTemplate, postgres_resport_templates.TopQueriesMDTemplate, postgres_resport_templates.QueryInfoMDTemplate)
			reportGenerator.OutputFileExtension = "md"
		}
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
	for _, query := range reportGenerator.DBPerfInfoRepository.Queries.GetList() {
		simplifiedQueryInfoList.Add(query)
	}

	var filePaths []string
	filePaths = append(filePaths, reportGenerator.Config.FilePaths...)
	filePaths = append(filePaths, reportGenerator.Config.S3Config.FileLocations...)

	// making accumulated info of all queries
	accumulatedInfoTemplateInput := stucts.InitAccumulatedInfoTemplateInput(simplifiedQueryInfoList, filePaths)

	// There are certain queries that are not important for profiler
	// like create, insert queries and queries with count less than minimum count
	var keysToRemove []string
	for s, info := range simplifiedQueryInfoList {
		if reportGenerator.shouldDiscardQuery(info.Query) || info.Count < reportGenerator.Config.MinimumQueryCallCount {
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

	reportString, err := reportGenerator.executeTemplates(accumulatedInfoTemplateInput, topQueriesTemplateInput, queryInfoTemplateInputs)
	if err != nil {
		log.Fatalln(err)
	}

	log.Debug(reportString)

	// Persisting the output to a file, from buffer

	// TODO externalise output file location
	f, err := os.Create("output." + reportGenerator.OutputFileExtension)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(reportString)
	w.Flush()
	log.Info("Report generated successfully, output." + reportGenerator.OutputFileExtension + " created.")
}

func (reportGenerator ReportGenerator) executeTemplates(accumulatedInfoTemplateInput stucts.AccumulatedInfoTemplateInput, topQueriesTemplateInput stucts.TopQueriesTemplateInput, queryInfoTemplateInputs []stucts.QueryInfoTemplateInput) (string, error) {
	var err error
	var bf bytes.Buffer
	err = reportGenerator.ReportTemplates.AccumulatedTemplate.Execute(&bf, accumulatedInfoTemplateInput)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}

	err = reportGenerator.ReportTemplates.TopQueriesTemplate.Execute(&bf, topQueriesTemplateInput)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}

	for i := 0; i < len(queryInfoTemplateInputs); i++ {
		err = reportGenerator.ReportTemplates.QueryInfoTemplate.Execute(&bf, queryInfoTemplateInputs[i])
		if err != nil {
			log.Fatalln(err.Error())
			return "", err
		}
	}
	return bf.String(), nil
}

func (reportGenerator ReportGenerator) sortSimplifiedQueryInfoList(list stucts.SimilarQueryInfoList, totalDuration float64) []*stucts.SimilarQueryInfo {
	return list.Sort(reportGenerator.Config.SortField, reportGenerator.Config.SortFieldOperation, reportGenerator.Config.SortOrder, totalDuration)
}

func (reportGenerator ReportGenerator) shouldDiscardQuery(query string) bool {
	if reportGenerator.discardQueryRegex != nil && reportGenerator.discardQueryRegex.MatchString(strings.TrimSpace(query)) {
		return true
	}
	return false
}

func getDiscardQueryRegex(discardQueries []string) *regexp.Regexp {
	if len(discardQueries) == 0 {
		return nil
	}
	regexText := "^((?i)" + discardQueries[0]
	for _, query := range discardQueries[1:] {
		regexText += "|(?i)" + query
	}
	regexText += ").*$"
	return regexp.MustCompile(regexText)
}
