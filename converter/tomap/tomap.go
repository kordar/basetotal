package tomap

import (
	"golang.org/x/exp/constraints"
)

type ToMap[K constraints.Integer | string, V constraints.Integer | constraints.Float | string] struct {
	origin  string
	data    map[K]V
	symbol1 string
	symbol2 string
}

func NewToMap[K constraints.Integer | string, V constraints.Integer | constraints.Float | string]() *ToMap[K, V] {
	return NewToMapWithSymbol[K, V](",", ":")
}

func NewToMapWithSymbol[K constraints.Integer | string, V constraints.Integer | constraints.Float | string](s1 string, s2 string) *ToMap[K, V] {
	return &ToMap[K, V]{origin: "", data: map[K]V{}, symbol1: s1, symbol2: s2}
}

func (p *ToMap[K, V]) SetDataValue(key K, value V) {
	p.data[key] = value
}

func (p *ToMap[K, V]) GetOriginData() string {
	return p.origin
}

func (p *ToMap[K, V]) SetOriginData(originData string) {
	p.origin = originData
}

func (p *ToMap[K, V]) GetSymbol1() string {
	return p.symbol1
}

func (p *ToMap[K, V]) GetSymbol2() string {
	return p.symbol2
}

func (p *ToMap[K, V]) IsEmpty() bool {
	return false
}

func (p *ToMap[K, V]) EmptyValue() interface{} {
	return nil
}

func (p *ToMap[K, V]) GetValue(key K) V {
	if p.IsEmpty() {
		return p.EmptyValue().(V)
	}
	return p.data[key]
}

func (p *ToMap[K, V]) GetParams() map[K]V {
	if p.data == nil {
		return make(map[K]V)
	}
	return p.data
}
