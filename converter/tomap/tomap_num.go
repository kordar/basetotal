package tomap

import (
	"github.com/spf13/cast"
	"strings"
)

// ToStrMapInt 转换为 string->int 字典
type ToStrMapInt struct {
	*ToMap[string, int]
}

func NewToStrMapInt() *ToStrMapInt {
	return &ToStrMapInt{ToMap: NewToMap[string, int]()}
}

func (t *ToStrMapInt) EmptyValue() interface{} {
	return 0
}

func (t *ToStrMapInt) Init(origin string) *ToStrMapInt {
	t.SetOriginData(origin)
	split := strings.Split(t.origin, t.symbol1)
	for _, s2 := range split {
		ss := strings.Split(s2, t.symbol2)
		if len(ss) == 2 {
			t.SetDataValue(ss[0], cast.ToInt(ss[1]))
		}
	}
	return t
}

// ToIntMapInt 转换为 int->int 字典
type ToIntMapInt struct {
	*ToMap[int, int]
}

func NewToIntMapInt() *ToIntMapInt {
	return &ToIntMapInt{ToMap: NewToMap[int, int]()}
}

func (t *ToIntMapInt) EmptyValue() interface{} {
	return 0
}

func (t *ToIntMapInt) Init(origin string) *ToIntMapInt {
	t.SetOriginData(origin)
	split := strings.Split(t.origin, t.symbol1)
	for _, s2 := range split {
		ss := strings.Split(s2, t.symbol2)
		if len(ss) == 2 {
			k := cast.ToInt(ss[0])
			v := cast.ToInt(ss[1])
			t.SetDataValue(k, v)
		}
	}
	return t
}

// ToIntMapStr 转换为 int->int 字典
type ToIntMapStr struct {
	*ToMap[int, string]
}

func NewToIntMapStr() *ToIntMapStr {
	return &ToIntMapStr{ToMap: NewToMap[int, string]()}
}

func (t *ToIntMapStr) EmptyValue() interface{} {
	return 0
}

func (t *ToIntMapStr) Init(origin string) *ToIntMapStr {
	t.SetOriginData(origin)
	split := strings.Split(t.origin, t.symbol1)
	for _, s2 := range split {
		ss := strings.Split(s2, t.symbol2)
		if len(ss) == 2 {
			k := cast.ToInt(ss[0])
			t.SetDataValue(k, ss[1])
		}
	}
	return t
}
