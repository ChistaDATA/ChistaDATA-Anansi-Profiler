package types

type StringSet map[string]struct{}

func InitStringSet(arr ...string) StringSet {
	s := map[string]struct{}{}
	for _, key := range arr {
		s[key] = struct{}{}
	}
	return s
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
