//============================================================
// @Description: 冒泡排序  时间复杂度O(N^2)  额外空间复杂度O(1) go里面应该是O(0)
// @Author: henry
// @Date: 2021/8/3 23:02
// @File: BubbleSort
//
//============================================================

package course01

func BubbleSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = SwagNum(arr[j], arr[j+1])
			}
		}
	}
	return arr
}

func SwagNum(a, b int) (int, int) {
	return b, a
}
