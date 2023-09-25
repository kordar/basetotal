package test

import (
	"github.com/kordar/basetotal/converter"
	"log"
	"testing"
)

func TestMealPeriod(t *testing.T) {
	mealPeriod := converter.NewMealPeriod("&")
	mealPeriod.Init("123&30,456&56")
	params := mealPeriod.GetParams()
	log.Printf("----%v, %v", params, mealPeriod.HasValue(converter.MealConfig{}))
}

func TestTimePeriod(t *testing.T) {
	mealPeriod := converter.NewTimePeriod()
	mealPeriod.Init("762:30,123:45")
	//params := mealPeriod.Init("123&30,456&56").GetParams()
	log.Printf("----%v", mealPeriod.GetDays(762))
}
