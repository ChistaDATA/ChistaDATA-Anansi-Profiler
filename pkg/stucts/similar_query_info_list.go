package stucts

import (
	"sort"
)

type SimilarQueryInfoList map[string]*SimilarQueryInfo

func InitSimilarQueryInfoList() SimilarQueryInfoList {
	return map[string]*SimilarQueryInfo{}
}

func (similarQueryInfoList *SimilarQueryInfoList) Remove(key string) {
	delete(*similarQueryInfoList, key)
}

func (similarQueryInfoList *SimilarQueryInfoList) Add(query *Query) {
	singleQueryInfo, ok := (*similarQueryInfoList)[query.GetTransformedQuery()]
	if !ok {
		singleQueryInfo = InitSimilarQueryInfo(query.Query)
		(*similarQueryInfoList)[query.GetTransformedQuery()] = singleQueryInfo
	}
	singleQueryInfo.Add(query)
}

func (similarQueryInfoList *SimilarQueryInfoList) SortQueryInfoByDuration() []*SimilarQueryInfo {
	var sortedQueryInfos []*SimilarQueryInfo
	for _, queryInfo := range *similarQueryInfoList {
		sortedQueryInfos = append(sortedQueryInfos, queryInfo)
	}
	sort.Sort(byDuration(sortedQueryInfos))
	return sortedQueryInfos
}

// byDuration implements sort.Interface based on the GetMaxDuration.
type byDuration []*SimilarQueryInfo

func (a byDuration) Len() int           { return len(a) }
func (a byDuration) Less(i, j int) bool { return a[i].GetMaxDuration() > a[j].GetMaxDuration() }
func (a byDuration) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
