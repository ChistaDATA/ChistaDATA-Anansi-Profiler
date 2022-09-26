package types

type StringSet map[string]struct{}

func InitStringSet() StringSet {
	return map[string]struct{}{}
}

func (set StringSet) Add(s string) {
	set[s] = struct{}{}
}

func (set StringSet) IsPresent(s string) bool {
	_, ok := set[s]
	return ok
}

func (set StringSet) ToArray() []string {
	var keys []string
	for i := range set {
		keys = append(keys, i)
	}
	return keys
}
