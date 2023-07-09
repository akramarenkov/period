package period

func clearInteger(input string) string {
	for id, symbol := range input {
		if id == len(input)-1 {
			return input[id:]
		}

		if symbol == '0' {
			continue
		}

		return input[id:]
	}

	return input[:0]
}

func clearFractional(slice string, fractionalSeparator byte) string {
	if len(slice) == 0 {
		return slice
	}

	high := len(slice)

	for id := high - 1; id >= 0; id-- {
		if slice[id] == '0' || slice[id] == fractionalSeparator || slice[id] == 0 {
			high = id
			continue
		}

		return slice[:high]
	}

	return slice[:high]
}
