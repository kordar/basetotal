package converter

import (
	"github.com/kordar/basetotal/converter/toslice"
	"github.com/spf13/cast"
	"strings"
)

type MealConfig struct {
	Start string
	End   string
}

type MealPeriod struct {
	*toslice.ToSlice[MealConfig]
	symbol2 string
}

func NewMealPeriod(symbol2 string) *MealPeriod {
	return &MealPeriod{ToSlice: toslice.NewToSlice[MealConfig](), symbol2: symbol2}
}

func (p *MealPeriod) Init(origin string) *MealPeriod {
	p.SetOriginData(origin)
	split := strings.Split(p.GetOriginData(), p.GetSymbol1())
	for _, s2 := range split {
		ss := strings.Split(s2, p.symbol2)
		if len(ss) == 2 {
			s := cast.ToString(ss[0])
			e := cast.ToString(ss[1])
			p.AddData(MealConfig{s, e})
		}
	}
	return p
}

func (p *MealPeriod) FilterStartAndEnd(key string) (string, string) {
	for _, item := range p.GetData() {
		if item.Start <= key && item.End >= key {
			return item.Start, item.End
		}
	}
	return "", ""
}
