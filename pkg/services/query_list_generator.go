package services

import (
	"bufio"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	log "github.com/sirupsen/logrus"
	"os"
)

func GenerateQueryList(config *stucts.CliConfig) stucts.QueryList {
	queries := stucts.InitQueryList()
	addQueriesFromFiles(config.FilePaths, &queries)
	return queries
}

func addQueriesFromFiles(paths []string, queryList *stucts.QueryList) {
	for _, file := range paths {
		readFileAndParseLogs(file, queryList)
	}
}

// readFileAndParseLogs Reads a file, extracts queries line by line
func readFileAndParseLogs(filePath string, queryList *stucts.QueryList) {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		parseLogAndAdd(scanner.Text(), queryList)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// parseLogAndAdd Parses a log line, checks it with multiple regexes, if matched adds extracted data to query list
func parseLogAndAdd(logLine string, queryList *stucts.QueryList) {
	cl, err := ParseClickHouseLog(logLine)
	if err != nil {
		log.Traceln(err)
	} else {
		q, ok := (*queryList)[cl.QueryId]
		if !ok {
			q = &stucts.Query{QueryId: cl.QueryId, Databases: types.InitStringSet(), Tables: types.InitStringSet(), ThreadIds: types.InitIntSet()}
			(*queryList)[cl.QueryId] = q
		}

		q.ThreadIds.Add(cl.ThreadId)

		err = ParseMessageWithQuery(cl.Message, q)
		if err == nil {
			q.Timestamp = cl.Timestamp //TODO move this else where
			return
		}
		err = ParseMessageWithDataInfo(cl.Message, q)
		if err == nil {
			return
		}
		err = ParseMessageWithPeakMemory(cl.Message, q)
		if err == nil {
			return
		}
		err = ParseMessageWithErrorInfo(cl.Message, q)
		if err == nil {
			q.Timestamp = cl.Timestamp //TODO move this else where
			return
		}
		err = ParseMessageWithAccessInfo(cl.Message, q)
		if err == nil {
			return
		}
	}
}
