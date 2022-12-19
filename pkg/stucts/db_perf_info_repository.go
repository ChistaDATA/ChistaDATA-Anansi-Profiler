package stucts

// DBPerfInfoRepository this is a canonical model which holds all information about database, queries etc
// now it only has queries, more info will be added later
type DBPerfInfoRepository struct {
	Queries QueryList
}

func InitDBPerfInfoRepository() *DBPerfInfoRepository {
	return &DBPerfInfoRepository{
		Queries: InitQueryList(),
	}
}
