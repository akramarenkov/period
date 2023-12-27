package period

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	period, found, err := Parse(" 3mo 10d 2y 23h 59m 58s 10ms 30us 10ns")
	require.NoError(t, err)
	require.True(t, found)

	require.Equal(t, 2, period.Years())
	require.Equal(t, 3, period.Months())
	require.Equal(t, 10, period.Days())
	require.False(t, period.IsNegative())
	require.Equal(t, time.Duration(86398010030010), period.Duration())

	period, found, err = Parse(" - 3mo10d2y23h59m58s10ms30us10ns")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, -2, period.Years())
	require.Equal(t, -3, period.Months())
	require.Equal(t, -10, period.Days())
	require.True(t, period.IsNegative())
	require.Equal(t, time.Duration(-86398010030010), period.Duration())
}

func TestParseEmpty(t *testing.T) {
	period, found, err := Parse("")
	require.NoError(t, err)
	require.Equal(t, Period{opts: Opts{Units: defaultUnits}}, period)
	require.False(t, found)

	period, found, err = Parse("   ")
	require.NoError(t, err)
	require.Equal(t, Period{opts: Opts{Units: defaultUnits}}, period)
	require.False(t, found)
}

func TestParseNotUniqueUnit(t *testing.T) {
	period, found, err := Parse("1y1y1mo1mo1d1d1h1h1m1m1s1s1ms1ms1us1us1ns1ns")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, 2, period.Years())
	require.Equal(t, 2, period.Months())
	require.Equal(t, 2, period.Days())
	require.False(t, period.IsNegative())
	require.Equal(
		t,
		2*time.Hour+
			2*time.Minute+
			2*time.Second+
			2*time.Millisecond+
			2*time.Microsecond+
			2*time.Nanosecond,
		period.Duration(),
	)
}

func TestParseCustom(t *testing.T) {
	input := " 3mo 10d 2y 23h 59m 58s 10ms 30us 10ns"

	regular, found, err := Parse(input)
	require.NoError(t, err)
	require.True(t, found)

	custom, found, err := ParseCustom(input, defaultUnits)
	require.NoError(t, err)
	require.True(t, found)

	unsafe, found, err := ParseCustomUnsafe(input, defaultUnits)
	require.NoError(t, err)
	require.True(t, found)

	withOpts, found, err := ParseWithOpts(input, Opts{Units: defaultUnits})
	require.NoError(t, err)
	require.True(t, found)

	require.Equal(t, regular, custom)
	require.Equal(t, regular, unsafe)
	require.Equal(t, regular, withOpts)
}

func TestParseExtraZerosResistance(t *testing.T) {
	dataSet := []struct {
		input    string
		expected time.Duration
	}{
		{
			".00000007000000h",
			252 * time.Microsecond,
		},
		{
			".00050000000000000000s",
			500 * time.Microsecond,
		},
		{
			".03333300000000000000000000000000s",
			33*time.Millisecond + 333*time.Microsecond,
		},
		{
			".00051000000000000000s",
			510 * time.Microsecond,
		},
		{
			".00000800000000000000s",
			8 * time.Microsecond,
		},
		{
			".00000100000000000000s",
			time.Microsecond,
		},
		{
			".100000000000000000m1s",
			7 * time.Second,
		},
		{
			".00000000200000000000s",
			2 * time.Nanosecond,
		},
	}

	opts := Opts{
		ExtraZerosResistance: true,
		Units:                defaultUnits,
	}

	for _, item := range dataSet {
		t.Run(
			item.input,
			func(t *testing.T) {
				regular, found, err := Parse(item.input)
				require.NoError(t, err)
				require.True(t, found)

				resistance, found, err := ParseWithOpts(item.input, opts)
				require.NoError(t, err)
				require.True(t, found)

				duration, err := time.ParseDuration(item.input)
				require.NoError(t, err)

				require.Equal(t, item.expected, resistance.Duration())
				require.NotEqual(t, item.expected, duration)
				require.Equal(t, duration, regular.Duration())
			},
		)
	}
}

func TestParseCustomInvalidUnitsTable(t *testing.T) {
	input := " 3mo 10d 2y 23h 59m 58s 10ms 30us 10ns"

	units := UnitsTable{
		UnitYear: {
			"y",
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
			"us",
			"µs",
		},
		UnitNanosecond: {
			"ns",
		},
	}

	_, _, err := ParseCustom(input, units)
	require.Error(t, err)

	_, _, err = ParseWithOpts(input, Opts{Units: units})
	require.Error(t, err)
}

func TestParseRequireError(t *testing.T) {
	inputs := []string{
		"d",
		" - 3mo 10d 2y 23h 59m 58s 10ms 30us 20µs 10ns 1",
		" - 3mo 10d 2y 23h 59m 58s 10ms 30us 10zs",
		" - ৩mo 10d 2y 23h 59m 58s 10ms 30us 10ns",
		" - 3mo 10d 2y 23h 59m ৩s 10ms 30us 10ns",
		" - 3mo 10d 2y 23h 59m 1.৩s 10ms 30us 10ns",
	}

	opts := Opts{
		UnitsMustBeUnique: true,
		Units:             defaultUnits,
	}

	for _, input := range inputs {
		period, found, err := Parse(input)
		require.Error(t, err)
		require.Equal(t, Period{}, period)
		require.False(t, found)

		period, found, err = ParseWithOpts(input, opts)
		require.Error(t, err)
		require.Equal(t, Period{}, period)
		require.False(t, found)
	}

	period, found, err := ParseWithOpts(
		" - 3mo 10d 2y 23h 59m 58s 10ms 30us 10ns 1us",
		opts,
	)
	require.Error(t, err)
	require.Equal(t, Period{}, period)
	require.False(t, found)
}

func TestParseOverflow(t *testing.T) {
	normal := []string{
		"9223372036854775807ns",
		"9223372036854775us",
		"9223372036854ms",
		"9223372036s",
		"153722867m",
		"2562047h",
		"2562047h2836s",
		"0.9223372036854775808s",
	}

	overflows := []string{
		"9223372036854775808ns",
		"9223372036854776us",
		"9223372036855ms",
		"9223372037s",
		"153722868m",
		"2562048h",
		"2562047h2837s",
		"2562046.5h30.5m2837.5s",
		"9223372036854775808y",
		"9223372036854775808mo",
		"9223372036854775808d",
		"9223372036854775807y1y",
		"9223372036854775807mo1mo",
		"9223372036854775807d1d",
		"10000000000s",
		"9223372036854775807000ns",
		"9223372036.9s",
	}

	for _, input := range normal {
		_, _, err := Parse(input)
		require.NoError(t, err)

		opts := Opts{
			Units: defaultUnits,
		}

		_, _, err = ParseWithOpts(input, opts)
		require.NoError(t, err)

		opts = Opts{
			Units:                defaultUnits,
			ExtraZerosResistance: true,
		}

		_, _, err = ParseWithOpts(input, opts)
		require.NoError(t, err)
	}

	for _, input := range overflows {
		_, _, err := Parse(input)
		require.Error(t, err)

		opts := Opts{
			Units: defaultUnits,
		}

		_, _, err = ParseWithOpts(input, opts)
		require.Error(t, err)

		opts = Opts{
			Units:                defaultUnits,
			ExtraZerosResistance: true,
		}

		_, _, err = ParseWithOpts(input, opts)
		require.Error(t, err)
	}
}

func TestShiftTime(t *testing.T) {
	period, found, err := Parse("1y")
	require.NoError(t, err)
	require.True(t, found)

	date := time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)

	expectedDate := time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC)
	expectedDuration := date.AddDate(1, 0, 0).Sub(date)

	unexpectedDuration := 365 * 24 * time.Hour
	unexpectedDate := date.Add(unexpectedDuration)

	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.NotEqual(t, unexpectedDate, period.ShiftTime(date))

	require.Equal(t, expectedDuration, period.RelativeDuration(date))
	require.NotEqual(t, unexpectedDuration, period.RelativeDuration(date))

	period, found, err = Parse("365d24h")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.Equal(t, expectedDuration, period.RelativeDuration(date))

	period, found, err = Parse("8760h1d")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.Equal(t, expectedDuration, period.RelativeDuration(date))
}

func TestShiftTimeNegative(t *testing.T) {
	period, found, err := Parse("-1y")
	require.NoError(t, err)
	require.True(t, found)

	date := time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC)

	expectedDate := time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)
	expectedDuration := date.AddDate(-1, 0, 0).Sub(date)

	unexpectedDuration := -365 * 24 * time.Hour
	unexpectedDate := date.Add(unexpectedDuration)

	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.NotEqual(t, unexpectedDate, period.ShiftTime(date))

	require.Equal(t, expectedDuration, period.RelativeDuration(date))
	require.NotEqual(t, unexpectedDuration, period.RelativeDuration(date))

	period, found, err = Parse("-365d24h")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.Equal(t, expectedDuration, period.RelativeDuration(date))

	period, found, err = Parse("-8760h1d")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.Equal(t, expectedDuration, period.RelativeDuration(date))
}

func TestString(t *testing.T) {
	period, found, err := Parse("-2y3mo10d")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, "-2y3mo10d0h0m0s", period.String())

	input := "-2y3mo10d23h59m58.01003001s"

	period, found, err = Parse(input)
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, input, period.String())
}

func TestStdLibraryCompatibility(t *testing.T) {
	inputs := []string{
		"0",
		".0s",
		"+1s",
		"1s",
		"1.1s",
		"1.000000001s",
		"1.0000000001s",
		"7.01001001s",
		"8.001s",
		"0201000s",
		"001m",
		"1h1h",
		".1s",
		"00h.0m",
		"+0",
		".0001s",
		".0101s",
		"0000000000000.000100100000000000s",
		".2700000h",
		"0.0000000001m",
		"0.0000000010000000000s",
		".00000000010000000000s",
		".18700000000000000000s",
		".00000000001m",
		"00000.00000000007m",
		"0.20000000000000000001s",
		".021000017h",
		"000000000000000000000h",
		".0000000000012h",
		".00000007000000h",
		"-.00000007000000h",
		".5000000000005555h",
		".0000000017s",
		"1.9007199279999999s",
		"-0.0000000007s",
		".0000000007s",
		"0μs",
		".9227000002799999700h",
		"-23.1h59.1m58.01003001s10.1ms10.1us1.1ns",
		"0.9223372036854775808s",
	}

	for _, input := range inputs {
		t.Run(
			input,
			func(t *testing.T) {
				period, found, err := Parse(input)
				require.NoError(t, err)
				require.True(t, found)

				duration, err := time.ParseDuration(input)
				require.NoError(t, err)
				require.Equal(t, duration, period.Duration())
				require.Equal(t, duration.String(), period.String())
			},
		)
	}
}

func TestAddDate(t *testing.T) {
	period, found, err := Parse("2y3mo10d23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.True(t, found)

	years := period.Years()
	months := period.Months()
	days := period.Days()

	require.NoError(t, period.AddDate(1, 10, 100))
	require.Equal(t, years+1, period.Years())
	require.Equal(t, months+10, period.Months())
	require.Equal(t, days+100, period.Days())

	require.NoError(t, period.AddDate(-1, -10, -100))
	require.Equal(t, years, period.Years())
	require.Equal(t, months, period.Months())
	require.Equal(t, days, period.Days())
}

func TestAddDateNegative(t *testing.T) {
	period, found, err := Parse("-2y3mo10d23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.True(t, found)

	years := period.Years()
	months := period.Months()
	days := period.Days()

	require.NoError(t, period.AddDate(1, 10, 100))
	require.Equal(t, years+1, period.Years())
	require.Equal(t, months+10, period.Months())
	require.Equal(t, days+100, period.Days())

	require.NoError(t, period.AddDate(-1, -10, -100))
	require.Equal(t, years, period.Years())
	require.Equal(t, months, period.Months())
	require.Equal(t, days, period.Days())
}

func TestAddDateRequireError(t *testing.T) {
	period, found, err := Parse("-2y3mo10d23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.True(t, found)

	require.Error(t, period.AddDate(math.MinInt, 0, 0))
	require.Error(t, period.AddDate(math.MinInt+1, 0, 0))
	require.Error(t, period.AddDate(0, math.MinInt, 0))
	require.Error(t, period.AddDate(0, math.MinInt+1, 0))
	require.Error(t, period.AddDate(0, 0, math.MinInt))
	require.Error(t, period.AddDate(0, 0, math.MinInt+1))

	period, found, err = Parse("2y3mo10d23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.True(t, found)

	require.Error(t, period.AddDate(math.MaxInt, 0, 0))
	require.Error(t, period.AddDate(math.MaxInt-1, 0, 0))
	require.Error(t, period.AddDate(0, math.MaxInt, 0))
	require.Error(t, period.AddDate(0, math.MaxInt-1, 0))
	require.Error(t, period.AddDate(0, 0, math.MaxInt))
	require.Error(t, period.AddDate(0, 0, math.MaxInt-1))
}

func TestAddDuration(t *testing.T) {
	period, found, err := Parse("2y3mo10d23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.True(t, found)

	duration := period.Duration()

	require.NoError(t, period.AddDuration(1*time.Hour))
	require.Equal(t, duration+1*time.Hour, period.Duration())

	require.NoError(t, period.AddDuration(-1*time.Hour))
	require.Equal(t, duration, period.Duration())
}

func TestAddDurationNegative(t *testing.T) {
	period, found, err := Parse("-2y3mo10d23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.True(t, found)

	duration := period.Duration()

	require.NoError(t, period.AddDuration(1*time.Hour))
	require.Equal(t, duration+1*time.Hour, period.Duration())

	require.NoError(t, period.AddDuration(-1*time.Hour))
	require.Equal(t, duration, period.Duration())
}

func TestAddDurationRequireError(t *testing.T) {
	period, found, err := Parse("-2y3mo10d23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.True(t, found)

	require.Error(t, period.AddDuration(math.MinInt64))
	require.Error(t, period.AddDuration(math.MinInt64+1))

	period, found, err = Parse("2y3mo10d23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.True(t, found)

	require.Error(t, period.AddDuration(math.MaxInt64))
	require.Error(t, period.AddDuration(math.MaxInt64-1))
}

func TestSetNegative(t *testing.T) {
	period, found, err := Parse("2y10ns")
	require.NoError(t, err)
	require.True(t, found)
	require.False(t, period.IsNegative())
	require.Equal(t, "2y0mo0d0h0m0.00000001s", period.String())

	period.SetNegative(true)
	require.True(t, period.IsNegative())
	require.Equal(t, "-2y0mo0d0h0m0.00000001s", period.String())
}

func TestSetYears(t *testing.T) {
	period, found, err := Parse("2y10ns")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, 2, period.Years())
	require.Equal(t, "2y0mo0d0h0m0.00000001s", period.String())

	require.NoError(t, period.SetYears(5))
	require.Equal(t, 5, period.Years())
	require.Equal(t, "5y0mo0d0h0m0.00000001s", period.String())

	require.Error(t, period.SetYears(-5))
}

func TestSetYearsNegative(t *testing.T) {
	period, found, err := Parse("-2y10ns")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, -2, period.Years())
	require.Equal(t, "-2y0mo0d0h0m0.00000001s", period.String())

	require.NoError(t, period.SetYears(-5))
	require.Equal(t, -5, period.Years())
	require.Equal(t, "-5y0mo0d0h0m0.00000001s", period.String())

	require.Error(t, period.SetYears(5))
	require.Error(t, period.SetYears(math.MinInt))

	require.NoError(t, period.SetYears(math.MinInt+1))
	require.Equal(t, math.MinInt+1, period.Years())
}

func TestSetMonths(t *testing.T) {
	period, found, err := Parse("2mo10ns")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, 2, period.Months())
	require.Equal(t, "2mo0d0h0m0.00000001s", period.String())

	require.NoError(t, period.SetMonths(5))
	require.Equal(t, 5, period.Months())
	require.Equal(t, "5mo0d0h0m0.00000001s", period.String())

	require.Error(t, period.SetMonths(-5))
}

func TestSetMonthsNegative(t *testing.T) {
	period, found, err := Parse("-2mo10ns")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, -2, period.Months())
	require.Equal(t, "-2mo0d0h0m0.00000001s", period.String())

	require.NoError(t, period.SetMonths(-5))
	require.Equal(t, -5, period.Months())
	require.Equal(t, "-5mo0d0h0m0.00000001s", period.String())

	require.Error(t, period.SetMonths(5))
	require.Error(t, period.SetMonths(math.MinInt))

	require.NoError(t, period.SetMonths(math.MinInt+1))
	require.Equal(t, math.MinInt+1, period.Months())
}

func TestSetDays(t *testing.T) {
	period, found, err := Parse("2d10ns")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, 2, period.Days())
	require.Equal(t, "2d0h0m0.00000001s", period.String())

	require.NoError(t, period.SetDays(5))
	require.Equal(t, 5, period.Days())
	require.Equal(t, "5d0h0m0.00000001s", period.String())

	require.Error(t, period.SetDays(-5))
}

func TestSetDaysNegative(t *testing.T) {
	period, found, err := Parse("-2d10ns")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, -2, period.Days())
	require.Equal(t, "-2d0h0m0.00000001s", period.String())

	require.NoError(t, period.SetDays(-5))
	require.Equal(t, -5, period.Days())
	require.Equal(t, "-5d0h0m0.00000001s", period.String())

	require.Error(t, period.SetDays(5))
	require.Error(t, period.SetDays(math.MinInt))

	require.NoError(t, period.SetDays(math.MinInt+1))
	require.Equal(t, math.MinInt+1, period.Days())
}

func TestSetDuration(t *testing.T) {
	period, found, err := Parse("2d10ns")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, time.Duration(10), period.Duration())
	require.Equal(t, "2d0h0m0.00000001s", period.String())

	require.NoError(t, period.SetDuration(5))
	require.Equal(t, time.Duration(5), period.Duration())
	require.Equal(t, "2d0h0m0.000000005s", period.String())

	require.Error(t, period.SetDuration(-5))
}

func TestSetDurationNegative(t *testing.T) {
	period, found, err := Parse("-2d10ns")
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, time.Duration(-10), period.Duration())
	require.Equal(t, "-2d0h0m0.00000001s", period.String())

	require.NoError(t, period.SetDuration(-5))
	require.Equal(t, time.Duration(-5), period.Duration())
	require.Equal(t, "-2d0h0m0.000000005s", period.String())

	require.Error(t, period.SetDuration(5))
	require.Error(t, period.SetDuration(math.MinInt64))

	require.NoError(t, period.SetDuration(math.MinInt64+1))
	require.Equal(t, time.Duration(math.MinInt64+1), period.Duration())
}

func TestNewPeriod(t *testing.T) {
	period := New()
	require.Equal(t, "0s", period.String())

	period, err := NewCustom(defaultUnits)
	require.NoError(t, err)
	require.Equal(t, "0s", period.String())

	period = NewCustomUnsafe(defaultUnits)
	require.Equal(t, "0s", period.String())

	period, err = NewWithOpts(Opts{Units: defaultUnits})
	require.NoError(t, err)
	require.Equal(t, "0s", period.String())
}

func TestMethodParse(t *testing.T) {
	period := New()

	found, err := period.Parse(" 3mo 10d 2y 23h 59m 58s 10ms 30us 10ns")
	require.NoError(t, err)
	require.True(t, found)

	found, err = period.Parse(" 3mo 10d 2y 23h 59m ৩s 10ms 30us 10ns")
	require.Error(t, err)
	require.False(t, found)
}

func TestNewPeriodCustomInvalidUnitsTable(t *testing.T) {
	units := UnitsTable{
		UnitYear: {
			"y",
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
			"us",
			"µs",
		},
		UnitNanosecond: {
			"ns",
		},
	}

	_, err := NewCustom(units)
	require.Error(t, err)

	_, err = NewWithOpts(Opts{Units: units})
	require.Error(t, err)
}

func benchmarkParseString(b *testing.B, custom bool) {
	input := " - 3mo 10d 2y 23.5h59m58s10ms30µs10ns"
	output := "-2y3mo10d24h29m58.01003001s"

	for attempt := 0; attempt < 100000; attempt++ {
		if custom {
			period, _, err := ParseCustom(input, defaultUnits)
			require.NoError(b, err)
			require.Equal(b, output, period.String())

			continue
		}

		period, _, err := Parse(input)
		require.NoError(b, err)
		require.Equal(b, output, period.String())
	}
}

func BenchmarkParseString(b *testing.B) {
	benchmarkParseString(b, false)
}

func BenchmarkParseCustomString(b *testing.B) {
	benchmarkParseString(b, true)
}

func FuzzFindPanic(f *testing.F) {
	f.Add("-2y3mo10d23h59m58.01003001s")
	f.Fuzz(
		func(t *testing.T, input string) {
			_, _, _ = Parse(input)
		},
	)
}

func FuzzSelfReparse(f *testing.F) {
	f.Add("-2y3mo10d23h59m58.01003001s")
	f.Fuzz(
		func(t *testing.T, input string) {
			period, _, err := Parse(input)
			if err != nil {
				return
			}

			stage1 := period.String()

			parsed, _, err := Parse(stage1)
			require.NoError(t, err)

			stage2 := parsed.String()
			require.Equal(t, stage1, stage2)

			reparsed, _, err := Parse(stage2)
			require.NoError(t, err)

			stage3 := reparsed.String()
			require.Equal(t, stage2, stage3)
		},
	)
}

func FuzzStdLibraryCompatibility(f *testing.F) {
	f.Add("-23h59m58.01003001s")
	f.Fuzz(
		func(t *testing.T, input string) {
			duration, err := time.ParseDuration(input)
			if err != nil {
				return
			}

			period, _, err := Parse(input)
			require.NoError(t, err)
			require.Equal(t, duration, period.Duration())
			require.Equal(t, duration.String(), period.String())
		},
	)
}
