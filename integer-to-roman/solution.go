package integer_to_roman

import (
	"fmt"
	"math"
)

type symbolValue struct {
	symbol string
	value int
}

func intToRoman(num int) string {
	symbols := []symbolValue{
		{ symbol: "M",  value: 1000 },
		{ symbol: "CM", value: 900},
		{ symbol: "D",  value: 500 },
		{ symbol: "CD", value: 400 },
		{ symbol: "C",  value: 100 },
		{ symbol: "XC", value: 90 },
		{ symbol: "L",  value: 50 },
		{ symbol: "XL", value: 40 },
		{ symbol: "X",  value: 10 },
		{ symbol: "IX", value: 9 },
		{ symbol: "V",  value: 5 },
		{ symbol: "IV", value: 4 },
		{ symbol: "I",  value: 1 },
	}
	result := ""
	for i := 0; i < len(symbols); i++ {
		v := int(math.Floor(float64(num / symbols[i].value)))

		for vi := 0; vi < v; vi++ {
			result = result + symbols[i].symbol
		}
		num = num - v*symbols[i].value
	}

	return result
}

func main() {
	fmt.Println(intToRoman(58))
}

