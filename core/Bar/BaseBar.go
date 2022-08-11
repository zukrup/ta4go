package bar

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	core "wlog3/ta4go/core"
)

var _ Bar = &BaseBar{}

type BaseBar struct {
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

func NewFakeBar() Bar {
	seed := math.Abs(rand.Float64())
	vol := math.Abs(float64(rand.Int()))
	return &BaseBar{
		openPrice:  core.New(seed * 0.25),
		closePrice: core.New(seed * 0.9),
		highPrice:  core.New(seed * 1.1),
		lowPrice:   core.New(seed * 0.1),
		amount:     core.New(seed * 5),
		volume:     core.New(vol),
		trades:     5,
		timePeriod: time.Hour,
		endTime:    time.Now().Add(time.Hour),
		beginTime:  time.Now(),
	}
}

func (b *BaseBar) GetOpenPrice() core.Num {
	return b.openPrice
}

// addPrice implements Bar
func (b *BaseBar) AddPrice(tradePrice core.Num) {
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
func (*BaseBar) AddPriceWithString(s string) {
	panic("unimplemented")
}

// addTrade implements Bar
func (*BaseBar) AddTrade(float32, float32) {
	panic("unimplemented")
}

// addTradeSimple implements Bar
func (b *BaseBar) AddTradeSimple(tradeVolume core.Num, tradePrice core.Num) {
	b.AddPrice(tradePrice)
	b.volume = b.volume.Plus(tradeVolume)
	b.amount = b.amount.Plus(tradeVolume.MultipliedBy(tradePrice))
}

// GetAmount implements Bar
func (b *BaseBar) GetAmount() core.Num {
	return b.amount
}

// GetBeginTime implements Bar
func (b *BaseBar) GetBeginTime() time.Time {
	return b.beginTime
}

// GetClosePrice implements Bar
func (b *BaseBar) GetClosePrice() core.Num {
	return b.closePrice
}

// GetDateName implements Bar
func (b *BaseBar) GetDateName() string {
	return b.endTime.Format(time.RFC3339)
}

// GetEndTime implements Bar
func (b *BaseBar) GetEndTime() time.Time {
	return b.endTime
}

// GetHighPrice implements Bar
func (b *BaseBar) GetHighPrice() core.Num {
	return b.highPrice
}

// GetSimpleDateName implements Bar
func (b *BaseBar) GetSimpleDateName() string {
	return b.endTime.Local().Format(time.RFC3339)
}

// GetTimePeriod implements Bar
func (b *BaseBar) GetTimePeriod() time.Duration {
	return b.timePeriod
}

// GetTrades implements Bar
func (b *BaseBar) GetTrades() int64 {
	return b.trades
}

// GetVolume implements Bar
func (b *BaseBar) GetVolume() core.Num {
	return b.volume
}

// inPeriod implements Bar
func (b *BaseBar) InPeriod(timestamp time.Time) bool {
	return !timestamp.Before(b.beginTime) && timestamp.Before(b.endTime)
}

// isBearish implements Bar
func (b *BaseBar) IsBearish() bool {
	return b.openPrice != nil && b.closePrice != nil && b.closePrice.IsLessThan(b.openPrice)
}

// isBullish implements Bar
func (b *BaseBar) IsBullish() bool {
	return b.openPrice != nil && b.closePrice != nil && b.openPrice.IsLessThan(b.closePrice)
}

// GetLowPrice implements Bar
func (b *BaseBar) GetLowPrice() core.Num {
	return b.lowPrice
}

func (b *BaseBar) ToString() string {
	return fmt.Sprintf("{end time: %d, close price: %s, open price: %s, low price: %s, high price: %s, volume: %s}",
		b.endTime.Local().Nanosecond(), b.closePrice.ToString(), b.openPrice.ToString(), b.lowPrice.ToString(), b.highPrice.ToString(),
		b.volume.ToString())
}
