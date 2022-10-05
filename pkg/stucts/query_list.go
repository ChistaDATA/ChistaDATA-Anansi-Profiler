package stucts

type QueryList map[string]*Query

func InitQueryList() QueryList {
	queryList := map[string]*Query{}
	return queryList
}
