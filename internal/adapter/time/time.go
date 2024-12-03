package time

import (
	"time"

	def "github.com/siyoga/jwt-auth-boilerplate/internal/adapter"
)

var _ def.TimeAdapter = (*adapter)(nil)

type (
	adapter struct {
		locale *time.Location
	}
)

func NewAdapter(locale int64) def.TimeAdapter {
	return &adapter{
		locale: time.FixedZone("MSC", int(locale)*3600),
	}
}

func (a *adapter) Now() time.Time {
	return time.Now().In(a.locale)
}

func (a *adapter) Locale() *time.Location {
	return a.locale
}

func (a *adapter) TodayMidnight() time.Time {
	t := a.Now().In(a.locale)
	res := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return res
}

func (a *adapter) TimeMidnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, a.locale)
}

func (a *adapter) MillisecondsToTime(milliseconds int64) time.Time {
	return time.UnixMilli(milliseconds).In(a.locale)
}
