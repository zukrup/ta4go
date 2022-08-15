package bar

import (
	"time"
	"wlog3/ta4go/core/Number"
)

type Bar interface {
	GetOpenPrice() Number.Num
	GetLowPrice() Number.Num
	GetHighPrice() Number.Num
	GetClosePrice() Number.Num
	GetVolume() Number.Num
	GetTrades() int64
	GetAmount() Number.Num
	GetTimePeriod() time.Duration
	GetBeginTime() time.Time
	GetEndTime() time.Time
	InPeriod(time.Time) bool
	GetDateName() string
	GetSimpleDateName() string
	IsBearish() bool
	IsBullish() bool
	AddTrade(float32, float32)
	AddTradeSimple(Number.Num, Number.Num)
	AddPriceWithString(string)
	AddPrice(Number.Num)
	ToString() string
}
