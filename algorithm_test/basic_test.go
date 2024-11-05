package algorithm

import (
	"fmt"
	"testing"
)

func Test_facto(t *testing.T) {
	n := 5
	fmt.Println("Factorial of", n, "is", facto(n))
}

func facto(n int) int {
	if n == 0 {
		return 1
	}
	return n * facto(n-1)
}

func Test_reverse(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	//reverse(nums)

	reverse2(nums)

	fmt.Println(nums)

	nums3 := []int{1, 2, 3}
	commonReverse(nums3)
	fmt.Println(nums3)

	fmt.Println(reverseString("abc"))
}

func reverseString(s string) string {
	arr := []rune(s)
	len := len(s)
	for i := 0; i < len/2; i++ {
		arr[i], arr[len-i-1] = arr[len-i-1], arr[i]
	}
	return string(arr)
}

func reverse2(nums []int) {
	len := len(nums)
	for i := 0; i < len/2; i++ {
		nums[i], nums[len-i-1] = nums[len-i-1], nums[i]
	}
}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func commonReverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Test_fib(t *testing.T) {
	fmt.Println(fib(10))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
