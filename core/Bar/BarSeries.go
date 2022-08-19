package bar

import "wlog3/ta4go/core/Number"

type BarSeries interface {
	GetName() string
	GetBar(int) Bar
	GetFirstBar() Bar
	GetLastBar() Bar
	GetBarCount() int
	IsEmpty() bool
	GetBarData() []Bar
	GetBeginIndex() int
	GetEndIndex() int
	GetMaximumBarCount() int
	SetMaximumBarCount(int)
	GetRemovedBarsCount() int
	AddBar(b Bar)
	AddBarReplaceIfExists(b Bar)
	AddTrade(volume Number.Num, price Number.Num)
	AddPrice(p Number.Num)
	GetSubSeries(startIndex int, endIndex int) BarSeries
	Function() Number.Function[float64, Number.Num]
}
