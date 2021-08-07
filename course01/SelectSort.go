//============================================================
// @Description: 选择排序:时间复杂度O(N^2) 额外空间复杂度O(1) go里面应该是O(0)
// @Author: henry
// @Date: 2021/8/3 23:13
// @File: SelectSort
//
//============================================================

package course01

func SelectSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	var min int
	for i := len(arr) - 1; i > 0; i-- {
		min = i
		for j := 0; j <= i; j++ {
			if arr[min] < arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = SwagNum(arr[i], arr[min])
	}
	return arr
}
