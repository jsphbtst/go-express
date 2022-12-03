package express

type StringSet []string

func (set *StringSet) Add(key string) {
	if set.Contains(key) {
		return
	}
	*set = append(*set, key)
}

// O(n) memory
func (set *StringSet) Remove(target string) {
	result := []string{}
	for _, vals := range *set {
		if vals == target {
			continue
		}
		result = append(result, vals)
	}
	*set = result
}

func (set *StringSet) Contains(target string) bool {
	for _, place := range *set {
		if place == target {
			return true
		}
	}
	return false
}
