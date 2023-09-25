package config

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"sync"
)

type TaskConfigConfigsWrapper[T TaskConfigDefine] struct {
	Data map[string]T
	Load func() (map[string]T, error)
}

func NewTaskConfigConfigsWrapper[T TaskConfigDefine](f func() (map[string]T, error)) *TaskConfigConfigsWrapper[T] {
	return &TaskConfigConfigsWrapper[T]{
		Data: map[string]T{},
		Load: f,
	}
}

type TaskConfigRefsWrapper struct {
	Data     map[string][]string
	TaskData map[string][]string
	Load     func() (map[string][]string, map[string][]string, error)
}

func NewTaskConfigRefsWrapper(f func() (map[string][]string, map[string][]string, error)) *TaskConfigRefsWrapper {
	return &TaskConfigRefsWrapper{
		Data:     make(map[string][]string),
		TaskData: make(map[string][]string),
		Load:     f,
	}
}

type TaskConfigTokensWrapper[T TaskConfigDefine] struct {
	Data map[string]T
	Load func() (map[string]T, error)
}

func NewTaskConfigTokensWrapper[T TaskConfigDefine](f func() (map[string]T, error)) *TaskConfigTokensWrapper[T] {
	return &TaskConfigTokensWrapper[T]{
		Data: map[string]T{},
		Load: f,
	}
}

type TaskConfigDefine interface {
	ConfigData() map[string]string
	IsEmpty() bool
}

type TaskConfig[C TaskConfigDefine, T TaskConfigDefine] struct {
	rw            *sync.RWMutex
	cronHandle    *cron.Cron
	spec          string
	configWrapper *TaskConfigConfigsWrapper[C]
	refWrapper    *TaskConfigRefsWrapper
	tokensWrapper *TaskConfigTokensWrapper[T]
}

func NewTaskConfig[C TaskConfigDefine, T TaskConfigDefine](
	configWrapper *TaskConfigConfigsWrapper[C],
	refWrapper *TaskConfigRefsWrapper,
	tokensWrapper *TaskConfigTokensWrapper[T],
	cronHandle *cron.Cron,
	spec string,
) *TaskConfig[C, T] {
	c := TaskConfig[C, T]{
		cronHandle:    cronHandle,
		spec:          spec,
		rw:            &sync.RWMutex{},
		configWrapper: configWrapper,
		refWrapper:    refWrapper,
		tokensWrapper: tokensWrapper,
	}
	c.Refresh()
	c.RefreshConfig()
	return &c
}

func (c *TaskConfig[C, T]) Refresh() {
	// "@every 10m" spec
	if c.spec != "" {
		_, _ = c.cronHandle.AddFunc(c.spec, func() {
			c.RefreshConfig()
		})
	}
}

func (c *TaskConfig[C, T]) RefreshConfig() {
	c.rw.Lock()
	defer c.rw.Unlock()

	// 1、refresh configs data
	if configData, err := c.configWrapper.Load(); err == nil {
		c.configWrapper.Data = configData
	} else {
		log.Warnf("refresh task configs fail, err = %v", err)
	}

	// 2、refresh configs ref
	if tokenData, taskData, err := c.refWrapper.Load(); err == nil {
		c.refWrapper.Data = tokenData
		c.refWrapper.TaskData = taskData
	} else {
		log.Warnf("refresh task ref fail, err = %v", err)
	}

	// 3、refresh configs tokens
	if tokenData, err := c.tokensWrapper.Load(); err == nil {
		c.tokensWrapper.Data = tokenData
	} else {
		log.Warnf("refresh task ref fail, err = %v", err)
	}

	log.Infof("refresh task configs finished!")
}

func (c *TaskConfig[C, T]) RefsBySign(sign string) []string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	if c.refWrapper.Data[sign] == nil {
		return make([]string, 0)
	}
	return c.refWrapper.Data[sign]
}

func (c *TaskConfig[C, T]) TaskRefsByToken(token string) []string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	if c.refWrapper.TaskData[token] == nil {
		return make([]string, 0)
	}
	return c.refWrapper.TaskData[token]
}

// MRefTokens 合并配置返回结果
func (c *TaskConfig[C, T]) MRefTokens(sign string, tokens []string) []string {
	reftokens := c.RefsBySign(sign)
	tokens = append(tokens, reftokens...)
	tmp := make(map[string]bool, 0)
	results := make([]string, 0)
	for _, t := range tokens {
		if tmp[t] == true {
			continue
		}
		tmp[t] = true
		results = append(results, t)
	}
	return results
}

// GetTokenById 获取token单独配置
func (c *TaskConfig[C, T]) GetTokenById(token string) T {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.tokensWrapper.Data[token]
}

func (c *TaskConfig[C, T]) GetConfigById(sign string) C {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.configWrapper.Data[sign]
}
