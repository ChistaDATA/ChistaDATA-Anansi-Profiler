package stucts

import (
	"sort"
)

type SimilarQueryInfoList map[string]*SimilarQueryInfo

func InitSimilarQueryInfoList() SimilarQueryInfoList {
	return map[string]*SimilarQueryInfo{}
}

func (similarQueryInfoList *SimilarQueryInfoList) GetList() []*SimilarQueryInfo {
	var sortedQueryInfos []*SimilarQueryInfo
	for _, queryInfo := range *similarQueryInfoList {
		sortedQueryInfos = append(sortedQueryInfos, queryInfo)
	}
	return sortedQueryInfos
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

//func (similarQueryInfoList *SimilarQueryInfoList) applySort(sortValues sort.Interface, sortDesc bool) {
//	if sortDesc {
//		sort.Sort(sort.Reverse(sortValues))
//	} else {
//		sort.Sort(sortValues)
//	}
//}
//
//func (similarQueryInfoList *SimilarQueryInfoList) SortQueryInfoByMaxDuration(sortDesc bool) []*SimilarQueryInfo {
//	sortedVals := byMaxDuration(similarQueryInfoList.GetList())
//	similarQueryInfoList.applySort(sortedVals, sortDesc)
//	return sortedVals
//}
//
//// byMaxDuration implements sort.Interface based on the GetMaxDuration.
//type byMaxDuration []*SimilarQueryInfo
//
//func (a byMaxDuration) Len() int           { return len(a) }
//func (a byMaxDuration) Less(i, j int) bool { return a[i].GetMaxDuration() < a[j].GetMaxDuration() }
//func (a byMaxDuration) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func getSortFunc(sortField string, sortFieldOperation string, totalDuration float64) func(info *SimilarQueryInfo) float64 {
	if sortField == SortFieldExecTime {
		return getSortFuncForExecTime(sortFieldOperation)
	} else if sortField == SortFieldBytesRead {
		return getSortFuncForBytesRead(sortFieldOperation)
	} else if sortField == SortFieldRowsRead {
		return getSortFuncForReadRows(sortFieldOperation)
	} else if sortField == SortFieldPeakMemory {
		return getSortFuncForPeakMemory(sortFieldOperation)
	} else if sortField == SortFieldQPS {
		return func(info *SimilarQueryInfo) float64 {
			return GetQPS(info, totalDuration)
		}
	} else if sortField == SortFieldQueryCount {
		return GetCount
	} else {
		return getSortFuncForExecTime(sortFieldOperation)
	}
}

func getSortFuncForExecTime(sortFieldOperation string) func(info *SimilarQueryInfo) float64 {
	if sortFieldOperation == SortFieldOperationSum {
		return GetTotalDuration
	} else if sortFieldOperation == SortFieldOperationMax {
		return GetMaxDuration
	} else if sortFieldOperation == SortFieldOperationMin {
		return GetMinDuration
	} else if sortFieldOperation == SortFieldOperationAvg {
		return GetAvgDuration
	} else if sortFieldOperation == SortFieldOperationPer95 {
		return GetPer95Duration
	} else if sortFieldOperation == SortFieldOperationStdDev {
		return GetStdDevDuration
	} else if sortFieldOperation == SortFieldOperationMedian {
		return GetMedianDuration
	} else {
		return GetMaxDuration
	}
}

func getSortFuncForBytesRead(sortFieldOperation string) func(info *SimilarQueryInfo) float64 {
	if sortFieldOperation == SortFieldOperationSum {
		return GetTotalReadBytes
	} else if sortFieldOperation == SortFieldOperationMax {
		return GetMaxReadBytes
	} else if sortFieldOperation == SortFieldOperationMin {
		return GetMinReadBytes
	} else if sortFieldOperation == SortFieldOperationAvg {
		return GetAvgReadBytes
	} else if sortFieldOperation == SortFieldOperationPer95 {
		return GetPer95ReadBytes
	} else if sortFieldOperation == SortFieldOperationStdDev {
		return GetStdDevReadBytes
	} else if sortFieldOperation == SortFieldOperationMedian {
		return GetMedianReadBytes
	} else {
		return GetMaxReadBytes
	}
}

func getSortFuncForReadRows(sortFieldOperation string) func(info *SimilarQueryInfo) float64 {
	if sortFieldOperation == SortFieldOperationSum {
		return GetTotalReadRows
	} else if sortFieldOperation == SortFieldOperationMax {
		return GetMaxReadRows
	} else if sortFieldOperation == SortFieldOperationMin {
		return GetMinReadRows
	} else if sortFieldOperation == SortFieldOperationAvg {
		return GetAvgReadRows
	} else if sortFieldOperation == SortFieldOperationPer95 {
		return GetPer95ReadRows
	} else if sortFieldOperation == SortFieldOperationStdDev {
		return GetStdDevReadRows
	} else if sortFieldOperation == SortFieldOperationMedian {
		return GetMedianReadRows
	} else {
		return GetMaxReadRows
	}
}

func getSortFuncForPeakMemory(sortFieldOperation string) func(info *SimilarQueryInfo) float64 {
	if sortFieldOperation == SortFieldOperationSum {
		return GetTotalPeakMemory
	} else if sortFieldOperation == SortFieldOperationMax {
		return GetMaxPeakMemory
	} else if sortFieldOperation == SortFieldOperationMin {
		return GetMinPeakMemory
	} else if sortFieldOperation == SortFieldOperationAvg {
		return GetAvgPeakMemory
	} else if sortFieldOperation == SortFieldOperationPer95 {
		return GetPer95PeakMemory
	} else if sortFieldOperation == SortFieldOperationStdDev {
		return GetStdDevPeakMemory
	} else if sortFieldOperation == SortFieldOperationMedian {
		return GetMedianPeakMemory
	} else {
		return GetMaxPeakMemory
	}
}

func (similarQueryInfoList *SimilarQueryInfoList) Sort(sortField string, sortFieldOperation string, sortOrder string, totalDuration float64) []*SimilarQueryInfo {

	list := similarQueryInfoList.GetList()

	sortFunc := getSortFunc(sortField, sortFieldOperation, totalDuration)

	a := func(i int, j int) bool {
		if sortOrder == SortOrderDesc {
			return sortFunc(list[i]) > sortFunc(list[j])
		}
		return sortFunc(list[i]) < sortFunc(list[j])
	}

	sort.Slice(list, a)
	return list
}
