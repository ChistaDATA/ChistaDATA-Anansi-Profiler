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
	q, ok := queryList.list[log.QueryId]
	queryList.lock.RUnlock()
	if !ok {
		queryList.lock.Lock()
		q, ok = queryList.list[log.QueryId]
		if !ok {
			q = &Query{QueryId: log.QueryId, Databases: types.InitStringSet(), Tables: types.InitStringSet(), ThreadIds: types.InitIntSet()}
			queryList.list[log.QueryId] = q
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
