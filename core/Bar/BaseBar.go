package bar

import (
	"time"
	core "wlog3/ta4go/core"
)

var _ Bar = &BaseBar{}

type BaseBar struct {
	Bar
	openPrice  core.Num
	closePrice core.Num
	highPrice  core.Num
	lowPrice   core.Num
	amount     core.Num
	volume     core.Num
	trades     int64
	timePeriod time.Duration
	endTime    time.Time
	beginTime  time.Time
}

func New(
	open float64,
	close float64,
	high float64,
	low float64,
	a float64,
	v float64,
	trades int64,
	period time.Duration,
	end time.Time,
	begin time.Time,

) Bar {
	return &BaseBar{
		openPrice:  core.New(open),
		closePrice: core.New(close),
		highPrice:  core.New(high),
		lowPrice:   core.New(low),
		amount:     core.New(a),
		volume:     core.New(v),
		trades:     trades,
		timePeriod: period,
		endTime:    end,
		beginTime:  begin,
	}
}

func (b *BaseBar) getOpenPrice() core.Num {
	return b.openPrice
}

// addPrice implements Bar
func (b *BaseBar) addPrice(tradePrice core.Num) {
	if b.openPrice == nil {
		b.openPrice = tradePrice
	}
	b.closePrice = tradePrice
	if b.highPrice == nil || b.highPrice.IsLessThan(tradePrice) {
		b.highPrice = tradePrice
	}
	if b.lowPrice == nil || b.lowPrice.IsGreaterThan(tradePrice) {
		b.lowPrice = tradePrice
	}
}

// addPriceWithString implements Bar
func (*BaseBar) addPriceWithString(s string) {
	panic("unimplemented")
}

// addTrade implements Bar
func (*BaseBar) addTrade(float32, float32) {
	panic("unimplemented")
}

// addTradeSimple implements Bar
func (b *BaseBar) addTradeSimple(tradeVolume core.Num, tradePrice core.Num) {
	b.addPrice(tradePrice)
	b.volume = b.volume.Plus(tradeVolume)
	b.amount = b.amount.Plus(tradeVolume.MultipliedBy(tradePrice))
}

// getAmount implements Bar
func (b *BaseBar) getAmount() core.Num {
	return b.amount
}

// getBeginTime implements Bar
func (b *BaseBar) getBeginTime() time.Time {
	return b.beginTime
}

// getClosePrice implements Bar
func (b *BaseBar) getClosePrice() core.Num {
	return b.closePrice
}

// getDateName implements Bar
func (b *BaseBar) getDateName() string {
	return b.endTime.Format(time.RFC3339)
}

// getEndTime implements Bar
func (b *BaseBar) getEndTime() time.Time {
	return b.endTime
}

// getHighPrice implements Bar
func (b *BaseBar) getHighPrice() core.Num {
	return b.highPrice
}

// getSimpleDateName implements Bar
func (b *BaseBar) getSimpleDateName() string {
	return b.endTime.Local().Format(time.RFC3339)
}

// getTimePeriod implements Bar
func (b *BaseBar) getTimePeriod() time.Duration {
	return b.timePeriod
}

// getTrades implements Bar
func (b *BaseBar) getTrades() int64 {
	return b.trades
}

// getVolume implements Bar
func (b *BaseBar) getVolume() core.Num {
	return b.volume
}

// inPeriod implements Bar
func (b *BaseBar) inPeriod(timestamp time.Time) bool {
	return !timestamp.Before(b.beginTime) && timestamp.Before(b.endTime)
}

// isBearish implements Bar
func (b *BaseBar) isBearish() bool {
	return b.openPrice != nil && b.closePrice != nil && b.closePrice.IsLessThan(b.openPrice)
}

// isBullish implements Bar
func (b *BaseBar) isBullish() bool {
	return b.openPrice != nil && b.closePrice != nil && b.openPrice.IsLessThan(b.closePrice)
}

// getLowPrice implements Bar
func (b *BaseBar) getLowPrice() core.Num {
	return b.lowPrice
}
