package algorithm_test

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestPFunc(t *testing.T) {
	fileData := loadFileData("p_data_2.txt")
	ordersTotal := 500000
	femaleTotal := 100
	orderCounts := make([]int, femaleTotal)

	for i := 0; i < ordersTotal; i++ {
		rand.Seed(time.Now().UnixNano())
		randIndex := rand.Intn(4471833)
		realIndex := BoundBinSearch(randIndex, fileData[100])
		if realIndex < 0 {
			fmt.Println(femaleTotal)
			fmt.Println(randIndex)
			fmt.Println(realIndex)
		}
		orderCounts[realIndex]++
	}

	fmt.Println(orderCounts)
}

func TestRandomGetIndex(t *testing.T) {
	start := time.Now().UnixNano()
	fileData := loadFileData("p_data.txt")
	fmt.Println(fmt.Sprintf("cost: %d ns", time.Now().UnixNano()-start))

	start = time.Now().Unix()
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		randIndex := 12200 //rand.Intn(4583333)
		realIndex := BoundBinSearch(randIndex, fileData[332])

		fmt.Println(realIndex)
	}
	fmt.Println(fmt.Sprintf("cost: %d s", time.Now().Unix()-start))
}

func TestLoadFileData(t *testing.T) {
	fileData := loadFileData("p_data.txt")
	fmt.Println(len(fileData))
	fmt.Println(len(fileData[2000]))
	fmt.Println(fileData[2000])

}

func loadFileData(fileName string) map[int][]int {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}

	mapData := make(map[int][]int)

	br := bufio.NewReader(fi)

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		s := string(a)
		lineData := strings.Split(s, "|")
		lineNum, _ := strconv.Atoi(lineData[0])
		strArr := strings.Split(lineData[1], ",")
		var intArr []int
		for _, s := range strArr {
			if len(s) == 0 {
				continue
			}
			i, _ := strconv.Atoi(s)
			intArr = append(intArr, i)
		}

		mapData[lineNum] = append(mapData[lineNum], intArr...)
	}
	fi.Close()
	return mapData
}

func readLine(r *bufio.Reader) (string, error) {
	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err
}

func TestBoundBinSearch(t *testing.T) {
	str := "0,1078491,2073567,2924194,3593750,4070027,4365234,4515991,4583333"
	strArr := strings.Split(str, ",")
	var intArr []int
	for _, s := range strArr {
		i, _ := strconv.Atoi(s)
		intArr = append(intArr, i)
	}
	fmt.Println(BoundBinSearch(124423, intArr))
	//fmt.Println(BoundBinSearch(867083, intArr))
	//fmt.Println(BoundBinSearch(1690000, intArr))
	//fmt.Println(BoundBinSearch(2433750, intArr))
	//fmt.Println(BoundBinSearch(3073333, intArr))
	//fmt.Println(BoundBinSearch(3990000, intArr))
	//fmt.Println(BoundBinSearch(4267083, intArr))
	//fmt.Println(BoundBinSearch(4440000, intArr))
	//fmt.Println(BoundBinSearch(4533750, intArr))
	//fmt.Println(BoundBinSearch(4583333, intArr))
	//
	//fmt.Println("---------------------------")
	//
	//fmt.Println(BoundBinSearch(-1, intArr))
	//fmt.Println(BoundBinSearch(0, intArr))
	//fmt.Println(BoundBinSearch(1, intArr))
	//fmt.Println(BoundBinSearch(867084, intArr))
	//fmt.Println(BoundBinSearch(1690001, intArr))
	//fmt.Println(BoundBinSearch(2433751, intArr))
	//fmt.Println(BoundBinSearch(3073334, intArr))
	//fmt.Println(BoundBinSearch(3990001, intArr))
	//fmt.Println(BoundBinSearch(4267084, intArr))
	//fmt.Println(BoundBinSearch(4440001, intArr))
	//fmt.Println(BoundBinSearch(4533751, intArr))
	//fmt.Println(BoundBinSearch(4583334, intArr))
}

func BoundBinSearch(target int, nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	numsLen := len(nums)
	low := 0
	high := numsLen - 1

	for low <= high {
		mid := (low + high) / 2

		if low == high {
			if low == 0 {
				return low
			}
			return mid - 1
		}

		if target > nums[high] {
			if high == numsLen-1 {
				return -1
			}
			return high
		}
		if target < nums[low] {
			return low - 1
		}

		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			if nums[low] == target {
				return low - 1
			} else if nums[high] == target {
				return high - 1
			} else {
				return mid - 1
			}
		}
	}

	return low
}
