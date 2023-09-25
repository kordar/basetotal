package toslice

import "golang.org/x/exp/constraints"

func GetEmptyMap[K constraints.Integer | string, V constraints.Integer | string](params []K, emptyValue V) map[K]V {
	m := make(map[K]V, 0)
	for _, v := range params {
		m[v] = emptyValue
	}
	return m
}
