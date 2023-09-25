package toslice

import (
	"github.com/spf13/cast"
	"strings"
)

// ToIntSlice 转换为int切片
type ToIntSlice struct {
	*ToSlice[int]
}

func NewToIntSlice() *ToIntSlice {
	return &ToIntSlice{ToSlice: NewToSlice[int]()}
}

func (t *ToIntSlice) Init(origin string) *ToIntSlice {
	t.SetOriginData(origin)
	split := strings.Split(t.origin, t.symbol1)
	for _, s := range split {
		t.AddData(cast.ToInt(s))
	}
	return t
}

// ToInt64Slice 转换为int64切片
type ToInt64Slice struct {
	*ToSlice[int64]
}

func NewToInt64Slice() *ToInt64Slice {
	return &ToInt64Slice{ToSlice: NewToSlice[int64]()}
}

func (t *ToInt64Slice) Init(origin string) *ToInt64Slice {
	t.SetOriginData(origin)
	split := strings.Split(t.origin, t.symbol1)
	for _, s := range split {
		t.AddData(cast.ToInt64(s))
	}
	return t
}
