package stucts

import "sync"

type QueryList struct {
	List map[string]*Query
	Lock sync.RWMutex
}

func InitQueryList() QueryList {
	return QueryList{
		List: map[string]*Query{},
	}
}
