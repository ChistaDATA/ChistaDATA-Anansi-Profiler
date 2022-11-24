package stucts

type InfoCorpus struct {
	Queries QueryList
}

func InitInfoCorpus() *InfoCorpus {
	return &InfoCorpus{
		Queries: InitQueryList(),
	}
}
