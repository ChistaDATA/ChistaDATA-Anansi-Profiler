package types

type IntSet map[int]struct{}

func InitIntSet() IntSet {
	return map[int]struct{}{}
}

func (set IntSet) Add(s int) {
	set[s] = struct{}{}
}

func (set IntSet) IsPresent(s int) bool {
	_, ok := set[s]
	return ok
}

func (set IntSet) ToArray() []int {
	var keys []int
	for i := range set {
		keys = append(keys, i)
	}
	return keys
}
