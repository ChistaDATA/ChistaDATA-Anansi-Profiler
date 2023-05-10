package stucts

// DBPerfInfoRepository this is a canonical model which holds all information about database, queries etc
// now it only has queries, more info will be added later
type DBPerfInfoRepository struct {
	Queries      QueryList
	CurrentQuery *Query
}

func InitDBPerfInfoRepository() *DBPerfInfoRepository {
	return &DBPerfInfoRepository{
		Queries:      InitQueryList(),
		CurrentQuery: &Query{},
	}
}

func CombineDBPerfInfoRepository(d ...*DBPerfInfoRepository) *DBPerfInfoRepository {
	r := InitDBPerfInfoRepository()
	for _, repository := range d {
		r.Queries.AppendQueryList(repository.Queries)
	}
	return r
}
