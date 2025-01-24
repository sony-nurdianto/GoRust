package note

import (
	"fmt"
	"reflect"
)

type WEEKDAY int

type DaysInfo struct {
	Name     string
	Sequence WEEKDAY
}

type Days struct {
	Sunday    DaysInfo
	Monday    DaysInfo
	Tuesday   DaysInfo
	Wednesday DaysInfo
	Thursday  DaysInfo
	Friday    DaysInfo
	Saturday  DaysInfo
}

const (
	SundaySquence WEEKDAY = iota
	MondaySquence
	TuesdaySquence
	WednesdaySquence
	ThursdaySquence
	FridaySquence
	SaturdaySquence
)

func StructReflect() {
	days := Days{
		Sunday: DaysInfo{
			Name:     "Sunday",
			Sequence: SundaySquence,
		},
		Monday: DaysInfo{
			Name:     "Monday",
			Sequence: MondaySquence,
		},
		Tuesday: DaysInfo{
			Name:     "Tuesday",
			Sequence: TuesdaySquence,
		},
		Wednesday: DaysInfo{
			Name:     "Wednesday",
			Sequence: WednesdaySquence,
		},
		Thursday: DaysInfo{
			Name:     "Thursday",
			Sequence: ThursdaySquence,
		},
		Friday: DaysInfo{
			Name:     "Friday",
			Sequence: FridaySquence,
		},
		Saturday: DaysInfo{
			Name:     "Saturday",
			Sequence: SaturdaySquence,
		},
	}

	v := reflect.ValueOf(days)

	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Field(i))
	}
}
