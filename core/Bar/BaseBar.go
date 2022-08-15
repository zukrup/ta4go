package bar

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"wlog3/ta4go/core/Number"
)

var _ Bar = &BaseBar{}

type BaseBar struct {
	openPrice  Number.Num
	closePrice Number.Num
	highPrice  Number.Num
	lowPrice   Number.Num
	amount     Number.Num
	volume     Number.Num
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
		openPrice:  Number.New(open),
		closePrice: Number.New(close),
		highPrice:  Number.New(high),
		lowPrice:   Number.New(low),
		amount:     Number.New(a),
		volume:     Number.New(v),
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
		openPrice:  Number.New(seed * 0.25),
		closePrice: Number.New(seed * 0.9),
		highPrice:  Number.New(seed * 1.1),
		lowPrice:   Number.New(seed * 0.1),
		amount:     Number.New(seed * 5),
		volume:     Number.New(vol),
		trades:     5,
		timePeriod: time.Hour,
		endTime:    time.Now().Add(time.Hour),
		beginTime:  time.Now(),
	}
}

func (b *BaseBar) GetOpenPrice() Number.Num {
	return b.openPrice
}

// addPrice implements Bar
func (b *BaseBar) AddPrice(tradePrice Number.Num) {
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
func (b *BaseBar) AddTradeSimple(tradeVolume Number.Num, tradePrice Number.Num) {
	b.AddPrice(tradePrice)
	b.volume = b.volume.Plus(tradeVolume)
	b.amount = b.amount.Plus(tradeVolume.MultipliedBy(tradePrice))
}

// GetAmount implements Bar
func (b *BaseBar) GetAmount() Number.Num {
	return b.amount
}

// GetBeginTime implements Bar
func (b *BaseBar) GetBeginTime() time.Time {
	return b.beginTime
}

// GetClosePrice implements Bar
func (b *BaseBar) GetClosePrice() Number.Num {
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
func (b *BaseBar) GetHighPrice() Number.Num {
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
func (b *BaseBar) GetVolume() Number.Num {
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
func (b *BaseBar) GetLowPrice() Number.Num {
	return b.lowPrice
}

func (b *BaseBar) ToString() string {
	return fmt.Sprintf("{end time: %d, close price: %s, open price: %s, low price: %s, high price: %s, volume: %s}",
		b.endTime.Local().Nanosecond(), b.closePrice.ToString(), b.openPrice.ToString(), b.lowPrice.ToString(), b.highPrice.ToString(),
		b.volume.ToString())
}
