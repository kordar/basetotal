package toslice

import (
	"strings"
)

// ToStrSlice 转换为int64切片
type ToStrSlice struct {
	*ToSlice[string]
}

func NewToStrSlice() *ToStrSlice {
	return &ToStrSlice{ToSlice: NewToSlice[string]()}
}

func (t *ToStrSlice) Init(origin string) *ToStrSlice {
	t.SetOriginData(origin)
	if origin == "" {
		return t
	}
	d := strings.Split(t.origin, t.symbol1)
	t.SetData(d)
	return t
}

func (t *ToStrSlice) IsEmpty() bool {
	return t.origin == ""
}
