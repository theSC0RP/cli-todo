package todo

type PriorityMap struct {
	ToInt map[string]int
	ToStr map[int]string
}

var Priorities *PriorityMap

func init() {
	Priorities = &PriorityMap{
		ToInt: map[string]int{
			"very less": 1,
			"less":      2,
			"medium":    3,
			"high":      4,
			"very high": 5,
		},
		ToStr: make(map[int]string),
	}
	for k, v := range Priorities.ToInt {
		Priorities.ToStr[v] = k
	}
}
