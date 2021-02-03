package main

import "fmt"

func romanToInt(s string) int {
	retValue := 0
	romanMap := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400,
		"C": 100, "XC": 90, "L": 50, "XL": 40,
		"X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}

	size := len(s)
	for i := 0; i < size; {
		if i < size-1 && romanMap[s[i:i+2]] > 0 {
			retValue += romanMap[s[i:i+2]]
			i += 2
		} else {
			retValue += romanMap[s[i:i+1]]
			i += 1
		}
	}

	return retValue
}

func main() {
	testArray := []string{"MCMXCIV", "MCMXCIV", "MMCMLIV", "LVIII", "IX", "XXX", "LXXXIX", "MMMDCCCXIX", "C", "CM", "XC", "VII"}
	for _, item := range testArray {
		fmt.Printf("test result %v is %v\n", item, romanToInt(item))
	}

}
