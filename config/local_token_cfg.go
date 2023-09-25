package config

import (
	"github.com/kordar/basetotal/converter/tomap"
	"github.com/kordar/basetotal/converter/toslice"
	"github.com/kordar/goutil"
	"github.com/spf13/viper"
)

// LocalTokenConfig 缓存本地Token独立配置，并返回，需传入prefix进行标识，格式 prexfix.xxxxx
type LocalTokenConfig struct {
	Cfg    *viper.Viper
	cache1 map[string]*tomap.ToStrMapInt
	cache2 map[string]*tomap.ToStrMapStr
	//cache3 map[string]*tools.MealPeriod
	cache4 map[string]*toslice.ToStrSlice
	cache5 map[string]*toslice.ToIntSlice
}

func NewLocalTokenConfig(key string) *LocalTokenConfig {
	sub := goutil.GetCfg().Sub(key)
	if sub == nil {
		sub = &viper.Viper{}
	}
	return &LocalTokenConfig{
		Cfg:    sub,
		cache1: make(map[string]*tomap.ToStrMapInt),
		cache2: make(map[string]*tomap.ToStrMapStr),
		//		cache3: map[string]*tools.MealPeriod{},
		cache4: make(map[string]*toslice.ToStrSlice),
		cache5: make(map[string]*toslice.ToIntSlice),
	}
}

func (c *LocalTokenConfig) SetCache1Value(key string, val *tomap.ToStrMapInt) {
	c.cache1[key] = val
}

func (c *LocalTokenConfig) SetCache2Value(key string, val *tomap.ToStrMapStr) {
	c.cache2[key] = val
}

//func (c *CommonTokenConfig) SetCache3Value(key string, val *tools.MealPeriod) {
//	c.cache3[key] = val
//}

func (c *LocalTokenConfig) SetCache4Value(key string, val *toslice.ToStrSlice) {
	c.cache4[key] = val
}

func (c *LocalTokenConfig) SetCache5Value(key string, val *toslice.ToIntSlice) {
	c.cache5[key] = val
}

func (c *LocalTokenConfig) GetCache1Value(key string) *tomap.ToStrMapInt {
	return c.cache1[key]
}

func (c *LocalTokenConfig) GetCache2Value(key string) *tomap.ToStrMapStr {
	return c.cache2[key]
}

//func (c *CommonTokenConfig) GetCache3Value(key string) *tools.MealPeriod {
//	return c.cache3[key]
//}

func (c *LocalTokenConfig) GetCache4Value(key string) *toslice.ToStrSlice {
	return c.cache4[key]
}

func (c *LocalTokenConfig) GetCache5Value(key string) *toslice.ToIntSlice {
	return c.cache5[key]
}

// GetMapString 获取token下所有配置
func (c *LocalTokenConfig) GetMapString(token string) map[string]string {
	return c.Cfg.GetStringMapString(token)
}

//
//// GetValueMapString 以map形式获取token下某个配置
//func (c *CommonTokenConfig) GetValueMapString(token string, key string) *tools.ToMapStr {
//	cachekey := token + "-" + key
//	if c.cache2[cachekey] == nil {
//		cc := tools.NewToMapStr()
//		sub := c.Cfg.Sub(token)
//		if sub != nil {
//			cc.Init(sub.GetString(key))
//		}
//		c.cache2[cachekey] = cc
//	}
//	return c.cache2[cachekey]
//}

// GetValueMapInt 以map形式获取token下某个配置，值为整型
func (c *LocalTokenConfig) GetValueMapInt(token string, key string) *tomap.ToStrMapInt {
	cachekey := token + "-" + key
	if c.cache1[cachekey] == nil {
		cc := tomap.NewToStrMapInt()
		sub := c.Cfg.Sub(token)
		if sub != nil {
			cc.Init(sub.GetString(key))
		}
		c.cache1[cachekey] = cc
	}
	return c.cache1[cachekey]
}

// GetValueSlice 获取token下某个key的slice配置数据
func (c *LocalTokenConfig) GetValueSlice(token string, key string) *toslice.ToStrSlice {
	cachekey := token + "-" + key
	if c.cache4[cachekey] == nil {
		cc := toslice.NewToStrSlice()
		sub := c.Cfg.Sub(token)
		if sub != nil {
			cc.Init(sub.GetString(key))
		}
		c.cache4[cachekey] = cc
	}
	return c.cache4[cachekey]
}

func (c *LocalTokenConfig) GetValueSliceInt(token string, key string) *toslice.ToIntSlice {
	cachekey := token + "-" + key
	if c.cache5[cachekey] == nil {
		cc := toslice.NewToIntSlice()
		sub := c.Cfg.Sub(token)
		if sub != nil {
			cc.Init(sub.GetString(key))
		}
		c.cache5[cachekey] = cc
	}
	return c.cache5[cachekey]
}

func (c *LocalTokenConfig) GetValueString(token string, key string) string {
	if c.Cfg.Sub(token) == nil {
		return ""
	}
	return c.Cfg.Sub(token).GetString(key)
}

//func (c *CommonTokenConfig) GetValueInt(token string, key string) int {
//	if c.Cfg.Sub(token) == nil {
//		return 0
//	}
//	return c.Cfg.Sub(token).GetInt(key)
//}
//
//func (c *CommonTokenConfig) GetValueInt64(token string, key string) int64 {
//	if c.Cfg.Sub(token) == nil {
//		return 0
//	}
//	return c.Cfg.Sub(token).GetInt64(key)
//}
//
//// GetExpired 获取过期时间配置
//func (c *CommonTokenConfig) GetExpired(token string, duration time.Duration) time.Duration {
//	if c.GetValueInt64(token, "expired") == 0 {
//		return duration
//	}
//	expired := c.GetValueInt64(token, "expired")
//	return time.Duration(expired) * time.Second
//}
//
//func (c *CommonTokenConfig) GetValueMealPeriod(token string, key string) *tools.MealPeriod {
//	cachekey := token + "-" + key
//	if c.cache3[cachekey] == nil {
//		cc := tools.NewMealPeriod()
//		sub := c.Cfg.Sub(token)
//		if sub != nil {
//			cc.Init(sub.GetString(key))
//		}
//		c.cache3[cachekey] = cc
//	}
//	return c.cache3[cachekey]
//}
