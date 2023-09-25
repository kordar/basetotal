package config

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"sync"
)

type MasterConfigDefine interface {
	GetTokenKey() string
	GetToken() string
}

type MasterConfig[T MasterConfigDefine] struct {
	rw         *sync.RWMutex
	cronHandle *cron.Cron
	spec       string
	load       func() ([]T, error)
	configs    []T
}

func NewMasterConfig[T MasterConfigDefine](load func() ([]T, error), cronHandle *cron.Cron, spec string) *MasterConfig[T] {
	c := MasterConfig[T]{
		cronHandle: cronHandle,
		spec:       spec,
		rw:         &sync.RWMutex{},
		configs:    make([]T, 0),
		load:       load,
	}
	c.Refresh()
	c.RefreshConfig()
	return &c
}

func (c *MasterConfig[T]) Refresh() {
	// "@every 10m" spec
	if c.spec != "" {
		_, _ = c.cronHandle.AddFunc(c.spec, func() {
			c.RefreshConfig()
		})
	}
}

func (c *MasterConfig[T]) RefreshConfig() {
	c.rw.Lock()
	defer c.rw.Unlock()
	if cfg, err := c.load(); err == nil {
		c.configs = cfg
		log.Infof("load master config success")
	} else {
		log.Warnf("load master config fail，err = %v", err)
	}
}

func (c *MasterConfig[T]) Scan(f func(conf T)) {
	c.rw.RLock()
	defer c.rw.RUnlock()
	for _, config := range c.configs {
		f(config)
	}
}
