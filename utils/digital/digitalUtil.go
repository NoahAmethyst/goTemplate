package digital

import "strconv"

func TransferStringToFloat(s string) float64 {
	number, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return number
}

func TransferFloatToString(f float64, digit int) string {
	var str string
	if f != 0 {
		str = strconv.FormatFloat(f, 'f', digit, 64)
	} else {
		str = "0.00"
	}

	return str
}
