package types

type StringCountPair struct {
	String string
	Count  int
}

type StringCountPairArray []StringCountPair

func (a StringCountPairArray) Len() int           { return len(a) }
func (a StringCountPairArray) Less(i, j int) bool { return a[i].Count < a[j].Count }
func (a StringCountPairArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
