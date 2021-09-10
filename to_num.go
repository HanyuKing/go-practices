package main

import "fmt"

func main() {
	fmt.Println(convertNum(1994))
}

func convertNum(num int) string {
	mapData := make(map[int]string)
	mapData[1] = "I"
	mapData[2] = "II"
	mapData[3] = "III"
	mapData[4] = "IV"
	mapData[5] = "V"
	mapData[9] = "IX"
	mapData[10] = "X"
	mapData[40] = "XL"
	mapData[50] = "L"
	mapData[90] = "XC"
	mapData[100] = "C"
	mapData[400] = "CD"
	mapData[500] = "D"
	mapData[900] = "CM"
	mapData[1000] = "M"

	baseNum := []int{1, 2, 3, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	result := ""

	for num != 0 {
		for i := len(baseNum) - 1; i >= 0; {
			if num-baseNum[i] >= 0 {
				result += mapData[baseNum[i]]
				num = num - baseNum[i]
			} else if num-baseNum[i] < 0 {
				i--
			}
		}
	}

	return result
}
