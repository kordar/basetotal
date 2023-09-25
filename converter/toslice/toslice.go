package toslice

// TODO 原始数据：value1,value2,value3,.... 通过英文逗号进行切片处理

// ToSlice 转换为slice
type ToSlice[T comparable] struct {
	origin  string
	data    []T
	symbol1 string
}

func NewToSlice[T comparable]() *ToSlice[T] {
	return NewToSliceWithSymbol[T](",")
}

func NewToSliceWithSymbol[T comparable](s string) *ToSlice[T] {
	return &ToSlice[T]{origin: "", data: make([]T, 0), symbol1: s}
}

func (p *ToSlice[T]) GetData() []T {
	return p.data
}

func (p *ToSlice[T]) SetData(items []T) {
	p.data = items
}

func (p *ToSlice[T]) AddData(item T) {
	p.data = append(p.data, item)
}

func (p *ToSlice[T]) GetOriginData() string {
	return p.origin
}

func (p *ToSlice[T]) SetOriginData(originData string) {
	p.origin = originData
}

func (p *ToSlice[T]) GetSymbol1() string {
	return p.symbol1
}

func (p *ToSlice[T]) HasValue(value T) bool {
	if p.data == nil {
		return false
	}
	for _, v := range p.data {
		if v == value {
			return true
		}
	}
	return false
}

func (p *ToSlice[T]) GetParams() []T {
	if p.data == nil {
		return []T{}
	}
	return p.data
}
