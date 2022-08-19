package bar

import (
	"time"
	"wlog3/ta4go/core/Number"
)

type BaseBarBuilder struct {
	timePeriod time.Duration
	endTime    time.Time
	openPrice  Number.Num
	closePrice Number.Num
	highPrice  Number.Num
	lowPrice   Number.Num
	amount     Number.Num
	volume     Number.Num
	trades     int64
}

func newBaseBarBuilder() *BaseBarBuilder {
	return &BaseBarBuilder{}
}

func (b *BaseBarBuilder) SetTimePeriod(d time.Duration) {
	b.timePeriod = d
}

func (b *BaseBarBuilder) SetEndTime(t time.Time) {
	b.endTime = t
}

func (b *BaseBarBuilder) SetOpenPrice(o Number.Num) {
	b.openPrice = o
}

func (b *BaseBarBuilder) SetClosePrice(c Number.Num) {
	b.closePrice = c
}

func (b *BaseBarBuilder) SetHighPrice(h Number.Num) {
	b.highPrice = h
}

func (b *BaseBarBuilder) SetLowPrice(l Number.Num) {
	b.lowPrice = l
}

func (b *BaseBarBuilder) SetAmount(a Number.Num) {
	b.amount = a
}

func (b *BaseBarBuilder) SetVolume(v Number.Num) {
	b.volume = v
}

func (b *BaseBarBuilder) SetTrades(t int64) {
	b.trades = t
}

func (b *BaseBarBuilder) getHouse() BaseBar {
	return BaseBar{
		openPrice:  b.openPrice,
		closePrice: b.closePrice,
		highPrice:  b.highPrice,
		lowPrice:   b.lowPrice,
		amount:     b.amount,
		volume:     b.volume,
		trades:     b.trades,
		timePeriod: b.timePeriod,
		endTime:    b.endTime,
		beginTime:  b.endTime.Add(-1 * b.timePeriod),
	}
}
