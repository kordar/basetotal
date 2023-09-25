package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
)

type ViperCfg struct {
	Cfg *viper.Viper
}

func NewViperCfg() *ViperCfg {
	return &ViperCfg{
		Cfg: viper.New(),
	}
}

func (c *ViperCfg) WriteConfig(in io.Reader) {
	c.Cfg.SetConfigType("json") // 或者 viper.SetConfigType("YAML")
	err := c.Cfg.ReadConfig(in)
	if err != nil {
		log.Printf("write config fail, err = %v", err)
	}
}
