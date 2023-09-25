package test

import (
	"fmt"
	"github.com/kordar/basetotal"
	"testing"
	"time"
)

func TestSpeedLimit(t *testing.T) {
	limit := basetotal.NewSpeedLimit()
	go func() {
		for i := 0; i < 50; i++ {
			limit.CanTotal(fmt.Sprintf("aaa-%d", i), "bbb")
		}
	}()

	time.Sleep(20 * time.Second)

}

type Aes struct {
	Age int
}

func (a Aes) compare(other Aes) int {
	return a.Age - other.Age
}

type Esc struct {
	Age int
}

func TestCompare(t *testing.T) {
	aes := Aes{Age: 12}
	esc := Aes{Age: 12}
	if aes == esc {
		println("aes > esc")
	} else {
		println("aes < esc")
	}
}
