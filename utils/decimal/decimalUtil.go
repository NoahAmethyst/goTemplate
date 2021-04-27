package decimal

import "github.com/shopspring/decimal"

func Add(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Add(d2)
}

// 减法
func Sub(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Sub(d2)
}

// 乘法
func Mul(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Mul(d2)
}

// 除法
func Div(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Div(d2)
}

// int
func DecimalToInt(d decimal.Decimal) int64 {
	return d.IntPart()
}

// float
func DecimalToFloat(d decimal.Decimal) float64 {
	f, _ := d.Float64()
	return f
}

func DecimalToString(d decimal.Decimal) string {
	return d.String()

}
