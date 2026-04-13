package date

import (
	"encoding/gob"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) String() string {
	return d.Time.Format("2006-01-02")
}
func (d Date) AsTime() time.Time {
	return d.Time
}
func (d Date) FromTime(tm time.Time) Date {
	d.Time = tm
	return d
}
func (d Date) Before(t time.Time) bool {
	return d.Time.Before(t)
}
func (d Date) After(t time.Time) bool {
	return d.Time.After(t)
}
func (d Date) Add(value time.Duration) Date {
	return d.FromTime(d.Time.Add(value).Truncate(time.Hour * 24))
}
func (d Date) Sub(t Date) time.Duration {
	return d.Time.Sub(t.Time)
}
func (d Date) Format(layout string) string {
	return d.Time.Format(layout)
}
func (d Date) Equal(other Date) bool {
	return d.Year() == other.Year() && d.Month() == other.Month() && d.Day() == other.Day()
}

func New(year int, month time.Month, day int) Date {
	return Date{time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)}
}
func Now() Date {
	now := time.Now().Truncate(time.Hour * 24)
	return New(now.Year(), now.Month(), now.Day())
}

func Parse(input, format string) (Date, error) {
	t, err := time.Parse(format, input)
	if err != nil {
		return Date{}, err
	}
	return Date{t}, nil
}

func init() {
	gob.RegisterName("Date", Date{})
}
