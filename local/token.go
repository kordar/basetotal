package local

import (
	"github.com/kordar/basetotal/converter/tomap"
	"github.com/kordar/goutil"
	"github.com/spf13/viper"
)

type TokenConfig struct {
	Cfg    *viper.Viper
	cache1 TokenCacheToStrMapIntWrapper
}

func NewLocalTokenConfig(key string) *TokenConfig {
	sub := goutil.GetCfg().Sub(key)
	if sub == nil {
		sub = &viper.Viper{}
	}
	return &TokenConfig{
		Cfg:    sub,
		cache1: TokenCacheToStrMapIntWrapper{data: map[string]*tomap.ToStrMapInt{}},
	}
}

func (t *TokenConfig) SetCacheValue(key string, value any) {
	switch v := value.(type) {
	case *tomap.ToStrMapInt:
		t.cache1.SetCacheValue(key, v)
		break
	}
}

func (t *TokenConfig) GetValueOfStrMapInt(key string) *tomap.ToStrMapInt {
	return t.cache1.GetCacheValue(key)
}

func (t *TokenConfig) GetTokenCacheValue(token string, key string, cfg *viper.Viper) *tomap.ToStrMapInt {
	return t.cache1.GetTokenCacheValue(token, key, cfg)
}

// ------------------------
type TokenCacheToStrMapIntWrapper struct {
	data map[string]*tomap.ToStrMapInt
}

func (w *TokenCacheToStrMapIntWrapper) SetCacheValue(key string, value *tomap.ToStrMapInt) {
	w.data[key] = value
}

func (w *TokenCacheToStrMapIntWrapper) GetCacheValue(key string) *tomap.ToStrMapInt {
	return w.data[key]
}

// GetTokenCacheValue 以map形式获取token下某个配置，值为整型
func (w *TokenCacheToStrMapIntWrapper) GetTokenCacheValue(token string, key string, cfg *viper.Viper) *tomap.ToStrMapInt {
	cachekey := token + "-" + key
	if w.data[cachekey] == nil {
		cc := tomap.NewToStrMapInt()
		sub := cfg.Sub(token)
		if sub != nil {
			cc.Init(sub.GetString(key))
		}
		w.data[cachekey] = cc
	}
	return w.data[cachekey]
}
