package tools

func Keys(m map[string]interface{}) []string {
	r := make([]string, len(m))
	index := 0
	for k, _ := range m {
		r[index] = k
		index++
	}
	return r
}
