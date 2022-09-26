package main

import (
	"bufio"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/parsers"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout) // TODO add file log supportz

	// Only log the warning severity or above.
	log.SetLevel(log.ErrorLevel) // TODO make configurable
}

func main() {

	cliConfig := types.InitializeCliConfig()

	queries := make(map[string]*types.Query)

	for _, file := range cliConfig.Files {
		readLogs(file, &queries)
	}

	queryMap := types.InitQueryInfoMap()
	for _, query := range queries {
		queryMap.Add(query)
	}

	var wg sync.WaitGroup

	for _, query := range queryMap {
		wg.Add(1)
		go func() {
			defer wg.Done()
			query.CompleteProcessing()
		}()
	}
	wg.Wait()
	queryMap.GenerateReportByDuration(&cliConfig)
}

func readLogs(filePath string, queries *map[string]*types.Query) {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		cl, err := parsers.ParseClickHouseLog(scanner.Text())
		if err != nil {
			log.Traceln(err)
		} else {
			q, ok := (*queries)[cl.QueryId]
			if !ok {
				q = &types.Query{QueryId: cl.QueryId, Databases: types.InitStringSet(), Tables: types.InitStringSet(), ThreadIds: types.InitIntSet()}
				(*queries)[cl.QueryId] = q
			}

			q.ThreadIds.Add(cl.ThreadId)

			err = parsers.ParseMessageWithQuery(cl.Message, q)
			if err == nil {
				q.Timestamp = cl.Timestamp //TODO move this else where
				continue
			}
			err = parsers.ParseMessageWithDataInfo(cl.Message, q)
			if err == nil {
				continue
			}
			err = parsers.ParseMessageWithPeakMemory(cl.Message, q)
			if err == nil {
				continue
			}
			err = parsers.ParseMessageWithErrorInfo(cl.Message, q)
			if err == nil {
				continue
			}
			err = parsers.ParseMessageWithAccessInfo(cl.Message, q)
			if err == nil {
				continue
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
