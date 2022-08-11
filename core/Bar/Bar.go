package bar

import (
	"time"
	"wlog3/ta4go/core"
)

type Bar interface {
	getOpenPrice() core.Num

	getLowPrice() core.Num
	getHighPrice() core.Num
	getClosePrice() core.Num
	getVolume() core.Num
	getTrades() int64
	getAmount() core.Num
	getTimePeriod() time.Duration
	getBeginTime() time.Time
	getEndTime() time.Time
	inPeriod(time.Time) bool
	getDateName() string
	getSimpleDateName() string
	isBearish() bool
	isBullish() bool
	addTrade(float32, float32)
	addTradeSimple(core.Num, core.Num)
	addPriceWithString(string)
	addPrice(core.Num)
}
