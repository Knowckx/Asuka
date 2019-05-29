package Array

// 把数组当前一个int,整体+1
// Example 1:

// Input: [1,2,3]  123+1=124
// Output: [1,2,4]
// Explanation: The array represents the integer 123.

// 9 --> 10       99 --> 100
func plusOne(digits []int) []int {
	length := len(digits)
	plus := true                       //是否进位
	for i := length - 1; i >= 0; i-- { //从List最后一个开始倒序
		if digits[i] < 9 {
			digits[i]++
			plus = false
			break
		} else {
			digits[i] = 0
		}
	}
	if plus {
		digits = append([]int{1}, digits...)
	}
	return digits
}
