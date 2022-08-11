package bar

import (
	"time"
	"wlog3/ta4go/core"
)

type Bar interface {
	GetOpenPrice() core.Num
	GetLowPrice() core.Num
	GetHighPrice() core.Num
	GetClosePrice() core.Num
	GetVolume() core.Num
	GetTrades() int64
	GetAmount() core.Num
	GetTimePeriod() time.Duration
	GetBeginTime() time.Time
	GetEndTime() time.Time
	InPeriod(time.Time) bool
	GetDateName() string
	GetSimpleDateName() string
	IsBearish() bool
	IsBullish() bool
	AddTrade(float32, float32)
	AddTradeSimple(core.Num, core.Num)
	AddPriceWithString(string)
	AddPrice(core.Num)
	ToString() string
}
