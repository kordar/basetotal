package converter

import "github.com/kordar/basetotal/converter/tomap"

// TimePeriod ; 统计周期, 商户Id => 周期(天) 0 所有
// time_period=762:30,123:45
type TimePeriod struct {
	*tomap.ToIntMapInt
}

func NewTimePeriod() *TimePeriod {
	return &TimePeriod{ToIntMapInt: tomap.NewToIntMapInt()}
}

func (p *TimePeriod) GetDays(shopId int) int {
	return p.GetValue(shopId)
}
