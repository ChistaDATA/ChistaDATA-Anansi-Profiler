package stucts

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/types"
	"sync"
)

// QueryList list of all queries, Query
type QueryList struct {
	list map[string]*Query
	lock sync.RWMutex
}

func InitQueryList() QueryList {
	return QueryList{
		list: map[string]*Query{},
	}
}

func (queryList *QueryList) Add(pq PartialQuery, log ExtractedLog) {
	queryList.lock.RLock()
	queryId := log.QueryId
	if pq.QueryId != nil {
		queryId = *pq.QueryId
	}
	q, ok := queryList.list[queryId]
	queryList.lock.RUnlock()
	if !ok {
		queryList.lock.Lock()
		q, ok = queryList.list[queryId]
		if !ok {
			databases := types.InitStringSet()
			if log.DatabaseName != "" {
				databases.Add(log.DatabaseName)
			}
			q = &Query{QueryId: queryId, Databases: databases, Tables: types.InitStringSet(), ThreadIds: types.InitIntSet(), User: log.UserName, ClientHost: log.RemoteHost}
			queryList.list[queryId] = q
		}
		queryList.lock.Unlock()
	}
	q.Add(pq)
}

func (queryList *QueryList) GetQuery(key string) *Query {
	queryList.lock.RLock()
	q, _ := queryList.list[key]
	queryList.lock.RUnlock()
	return q
}

func (queryList *QueryList) GetList() map[string]*Query {
	return queryList.list
}

func (queryList *QueryList) AppendQueryList(m QueryList) {
	queryList.lock.Lock()
	defer queryList.lock.Unlock()
	for k, v := range m.GetList() {
		queryList.list[k] = v
	}
}
