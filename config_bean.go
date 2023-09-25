package basetotal

import (
	"github.com/kordar/goutil"
	"github.com/spf13/cast"
	"strings"
)

// ConfigBean 管理configs相关配置
type ConfigBean struct {
	tokens []string
	all    bool
	params map[string]string
}

// NewConfigBeanWithLocal 从本地加载tokens
func NewConfigBeanWithLocal(id string) *ConfigBean {
	section := goutil.GetSection(id)
	tokens := strings.Split(section["tokens"], ",")
	return NewConfigBean(tokens, cast.ToBool(section["all"]), section)
}

func NewConfigBean(tokens []string, all bool, params map[string]string) *ConfigBean {
	return &ConfigBean{tokens: tokens, all: all, params: params}
}

func (t *ConfigBean) Tokens() []string {
	return t.tokens
}

func (t *ConfigBean) SetTokens(tokens []string) {
	t.tokens = tokens
}

func (t *ConfigBean) All() bool {
	return t.all
}

func (t *ConfigBean) SetAll(all bool) {
	t.all = all
}

func (t *ConfigBean) Params() map[string]string {
	return t.params
}

func (t *ConfigBean) SetParams(params map[string]string) {
	t.params = params
}

func (t *ConfigBean) GetParam(key string) string {
	return t.params[key]
}
