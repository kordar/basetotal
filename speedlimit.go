package basetotal

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"sync/atomic"
)

type SpeedLimit struct {
	locker map[string]*uint32
	rw     sync.RWMutex
}

func NewSpeedLimit() *SpeedLimit {
	return &SpeedLimit{locker: make(map[string]*uint32)}
}

func (t *SpeedLimit) CanTotal(token string, subname string) bool {
	
	if subname != "" {
		token = token + ":" + subname
	}

	if t.locker[token] == nil {
		var n uint32 = 0
		t.locker[token] = &n
	}

	if atomic.CompareAndSwapUint32(t.locker[token], 0, 1) {
		return true
	}
	log.Warnf("任务处理中，token = %s, subname = %s", token, subname)
	return false
}

func (t *SpeedLimit) Clear(token string, subname string) {
	if subname != "" {
		token = token + ":" + subname
	}
	if t.locker[token] == nil {
		return
	}
	atomic.StoreUint32(t.locker[token], 0)
}
