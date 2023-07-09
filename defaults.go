package period

const (
	defaultFormatFractionalSize uint = 9
	defaultFractionalSeparator  byte = '.'
	defaultMinusSign            byte = '-'
	defaultNumberBase           uint = 10
	defaultPlusSign             byte = '+'
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
