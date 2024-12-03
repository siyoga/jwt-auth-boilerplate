package adapter

import "time"

type (
	RandomAdapter interface {
		RandomString(length int) string
		RandomStringWithTimeNanoSeed(length int) string
		RandomIntn(n int) int
		RandomToken(length int) (string, error)
	}

	TimeAdapter interface {
		Now() time.Time
		TodayMidnight() time.Time
		TimeMidnight(t time.Time) time.Time
		MillisecondsToTime(ms int64) time.Time
		Locale() *time.Location
	}
)
