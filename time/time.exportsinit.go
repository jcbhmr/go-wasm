package time

import (
	"sync"
	"time"

	"runtime/cgo"

	"github.com/jcbhmr/go-wasm/builtin"
	"go.bytecodealliance.org/cm"
)

func init() {
	// Constants

	Exports.Layout = func() (result string) {
		return time.Layout
	}
	Exports.ANSIC = func() (result string) {
		return time.ANSIC
	}
	Exports.UnixDate = func() (result string) {
		return time.UnixDate
	}
	Exports.RubyDate = func() (result string) {
		return time.RubyDate
	}
	Exports.RFC822 = func() (result string) {
		return time.RFC822
	}
	Exports.RFC822Z = func() (result string) {
		return time.RFC822Z
	}
	Exports.RFC850 = func() (result string) {
		return time.RFC850
	}
	Exports.RFC1123 = func() (result string) {
		return time.RFC1123
	}
	Exports.RFC1123Z = func() (result string) {
		return time.RFC1123Z
	}
	Exports.RFC3339 = func() (result string) {
		return time.RFC3339
	}
	Exports.RFC3339Nano = func() (result string) {
		return time.RFC3339Nano
	}
	Exports.Kitchen = func() (result string) {
		return time.Kitchen
	}
	Exports.Stamp = func() (result string) {
		return time.Stamp
	}
	Exports.StampMilli = func() (result string) {
		return time.StampMilli
	}
	Exports.StampMicro = func() (result string) {
		return time.StampMicro
	}
	Exports.StampNano = func() (result string) {
		return time.StampNano
	}
	Exports.DateTime = func() (result string) {
		return time.DateTime
	}
	Exports.DateOnly = func() (result string) {
		return time.DateOnly
	}
	Exports.TimeOnly = func() (result string) {
		return time.TimeOnly
	}

	Exports.Nanosecond = func() (result Duration) {
		return Duration(time.Nanosecond)
	}
	Exports.Microsecond = func() (result Duration) {
		return Duration(time.Microsecond)
	}
	Exports.Millisecond = func() (result Duration) {
		return Duration(time.Millisecond)
	}
	Exports.Second = func() (result Duration) {
		return Duration(time.Second)
	}
	Exports.Minute = func() (result Duration) {
		return Duration(time.Minute)
	}
	Exports.Hour = func() (result Duration) {
		return Duration(time.Hour)
	}

	// Functions

	Exports.Sleep = func(d Duration) {
		time.Sleep(time.Duration(d))
	}

	// Types

	Exports.ParseDuration = func(s string) (result cm.Result[int64, Duration, Error]) {
		d, err := time.ParseDuration(s)
		if err != nil {
			result.SetErr(builtin.ErrorResourceNew(cm.Rep(cgo.NewHandle(err))))
			return
		}
		result.SetOK(Duration(d))
		return
	}
	Exports.Since = func(t Time) (result Duration) {
		return Duration(time.Since(cgo.Handle(t).Value().(time.Time)))
	}
	Exports.Until = func(t Time) (result Duration) {
		return Duration(time.Until(cgo.Handle(t).Value().(time.Time)))
	}
	Exports.DurationAbs = func(self Duration) (result Duration) {
		return Duration(time.Duration(self).Abs())
	}
	Exports.DurationHours = func(self Duration) (result float64) {
		return time.Duration(self).Hours()
	}
	Exports.DurationMicroseconds = func(self Duration) (result int64) {
		return time.Duration(self).Microseconds()
	}
	Exports.DurationMilliseconds = func(self Duration) (result int64) {
		return time.Duration(self).Milliseconds()
	}
	Exports.DurationNanoseconds = func(self Duration) (result int64) {
		return time.Duration(self).Nanoseconds()
	}
	Exports.DurationRound = func(self, m Duration) (result Duration) {
		return Duration(time.Duration(self).Round(time.Duration(m)))
	}
	Exports.DurationSeconds = func(self Duration) (result float64) {
		return time.Duration(self).Seconds()
	}
	Exports.DurationString = func(self Duration) (result string) {
		return time.Duration(self).String()
	}
	Exports.DurationTruncate = func(self, m Duration) (result Duration) {
		return Duration(time.Duration(self).Truncate(time.Duration(m)))
	}

	Exports.Location.Constructor = func() (result Location) {
		return Location(cgo.NewHandle(&time.Location{}))
	}
	Exports.Location.Destructor = func(self cm.Rep) {
		cgo.Handle(self).Delete()
	}
	Exports.Local = sync.OnceValue(func() (result Location) {
		return Location(cgo.NewHandle(time.Local))
	})
	Exports.UTC = sync.OnceValue(func() (result Location) {
		return Location(cgo.NewHandle(time.UTC))
	})
	Exports.FixedZone = func(name string, offset int32) (result Location) {
		return Location(cgo.NewHandle(time.FixedZone(name, int(offset))))
	}
	Exports.LoadLocation = func(name string) (result cm.Result[Location, Location, Error]) {
		loc, err := time.LoadLocation(name)
		if err != nil {
			result.SetErr(builtin.ErrorResourceNew(cm.Rep(cgo.NewHandle(err))))
			return
		}
		result.SetOK(Location(cgo.NewHandle(loc)))
		return
	}
	Exports.LoadLocationFromTZData = func(name string, data cm.List[uint8]) (result cm.Result[Location, Location, Error]) {
		loc, err := time.LoadLocationFromTZData(name, data.Slice())
		if err != nil {
			result.SetErr(builtin.ErrorResourceNew(cm.Rep(cgo.NewHandle(err))))
			return
		}
		result.SetOK(Location(cgo.NewHandle(loc)))
		return
	}
	Exports.Location.String = func(self cm.Rep) (result string) {
		return cgo.Handle(self).Value().(*time.Location).String()
	}

	Exports.January = func() (result Month) {
		return Month(time.January)
	}
	Exports.February = func() (result Month) {
		return Month(time.February)
	}
	Exports.March = func() (result Month) {
		return Month(time.March)
	}
	Exports.April = func() (result Month) {
		return Month(time.April)
	}
	Exports.May = func() (result Month) {
		return Month(time.May)
	}
	Exports.June = func() (result Month) {
		return Month(time.June)
	}
	Exports.July = func() (result Month) {
		return Month(time.July)
	}
	Exports.August = func() (result Month) {
		return Month(time.August)
	}
	Exports.September = func() (result Month) {
		return Month(time.September)
	}
	Exports.October = func() (result Month) {
		return Month(time.October)
	}
	Exports.November = func() (result Month) {
		return Month(time.November)
	}
	Exports.December = func() (result Month) {
		return Month(time.December)
	}
	Exports.MonthString = func(self Month) (result string) {
		return time.Month(self).String()
	}

	Exports.Time.ExportConstructor = func() (result Time) {
		return Time(cgo.NewHandle(time.Time{}))
	}
	Exports.Time.Destructor = func(self cm.Rep) {
		cgo.Handle(self).Delete()
	}
	Exports.Date = func(year int32, month Month, day, hour, min_, sec, nsec int32, loc Location) (result Time) {
		return Time(cgo.NewHandle(time.Date(int(year), time.Month(month), int(day), int(hour), int(min_), int(sec), int(nsec), cgo.Handle(loc).Value().(*time.Location))))
	}
	Exports.Now = func() (result Time) {
		return Time(cgo.NewHandle(time.Now()))
	}
	Exports.Parse = func(layout, value string) (result cm.Result[Time, Time, Error]) {
		t, err := time.Parse(layout, value)
		if err != nil {
			result.SetErr(builtin.ErrorResourceNew(cm.Rep(cgo.NewHandle(err))))
			return
		}
		result.SetOK(Time(cgo.NewHandle(t)))
		return
	}
	Exports.ParseInLocation = func(layout, value string, loc Location) (result cm.Result[Time, Time, Error]) {
		t, err := time.ParseInLocation(layout, value, cgo.Handle(loc).Value().(*time.Location))
		if err != nil {
			result.SetErr(builtin.ErrorResourceNew(cm.Rep(cgo.NewHandle(err))))
			return
		}
		result.SetOK(Time(cgo.NewHandle(t)))
		return
	}
}
