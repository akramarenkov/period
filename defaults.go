package period

const (
	defaultFormatFractionalSize uint = 9
	defaultFractionalSeparator       = '.'
	defaultMinusSign                 = '-'
	defaultNumberBase           uint = 10
	defaultPlusSign                  = '+'
)

var defaultUnits = UnitsTable{
	UnitYear: {
		"y",
	},
	UnitMonth: {
		"mo",
	},
	UnitDay: {
		"d",
	},
	UnitHour: {
		"h",
	},
	UnitMinute: {
		"m",
	},
	UnitSecond: {
		"s",
	},
	UnitMillisecond: {
		"ms",
	},
	UnitMicrosecond: {
		"µs",
		"μs",
		"us",
	},
	UnitNanosecond: {
		"ns",
	},
}
