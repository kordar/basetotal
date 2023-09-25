package config

import (
	"bytes"
	"encoding/json"
	"github.com/kordar/goutil"
	"github.com/robfig/cron/v3"
	"log"
	"testing"
	"time"
)

type MasterConfigItem struct {
	A  string
	BB int
}

func (receiver MasterConfigItem) GetTokenKey() string {
	goutil.GetCfg().ReadRemoteConfig()
	return "ccccvc"
}

func (receiver MasterConfigItem) GetToken() string {
	return "ccccvc"
}

func TestMasterConfig_Scan(t *testing.T) {
	c := cron.New()
	master := NewMasterConfig[MasterConfigItem](func() ([]MasterConfigItem, error) {
		items := make([]MasterConfigItem, 0)
		items = append(items, MasterConfigItem{A: "A", BB: 123})
		items = append(items, MasterConfigItem{A: "A", BB: 123})
		return items, nil
	}, c, "@every 5s")
	go c.Start()
	master.Scan(func(conf MasterConfigItem) {
		log.Printf("scan data = %v, token = %s", conf, conf.GetTokenKey())
	})
	time.Sleep(20 * time.Second)
}

func TestViperCfg(t *testing.T) {
	cfg := NewViperCfg()
	m := map[string]interface{}{
		"name": "xxxx", "age": []string{"AAA", "BBB"},
	}
	var yamlExample, _ = json.Marshal(&m)

	cfg.WriteConfig(bytes.NewBuffer(yamlExample))
	//name := cfg.Cfg.GetString("name")

	m2 := map[string]interface{}{
		"name2": "xxxx", "rr": []string{"TTTT", "CCC"},
	}
	var yamlExample2, _ = json.Marshal(&m2)
	cfg.WriteConfig(bytes.NewBuffer(yamlExample2))
	log.Printf("--------%v, %v, %v", cfg.Cfg.GetString("name"), cfg.Cfg.GetStringSlice("age"), cfg.Cfg.GetStringSlice("rr"))
}
