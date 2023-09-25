package tomap

import (
	"strings"
)

// ToStrMapStr 转换为 string->string 字典
type ToStrMapStr struct {
	*ToMap[string, string]
}

func NewToStrMapStr() *ToStrMapStr {
	return &ToStrMapStr{ToMap: NewToMap[string, string]()}
}

func (t *ToStrMapStr) EmptyValue() interface{} {
	return ""
}

func (t *ToStrMapStr) Init(origin string) *ToStrMapStr {
	t.SetOriginData(origin)
	split := strings.Split(t.origin, t.symbol1)
	for _, s2 := range split {
		ss := strings.Split(s2, t.symbol2)
		if len(ss) == 2 {
			t.SetDataValue(ss[0], ss[1])
		}
	}
	return t
}
