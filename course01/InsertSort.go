//============================================================
// @Description:
// @Author: henry
// @Date: 2021/8/3 23:49
// @File: InsertSort.go
//
//============================================================

package course01

func InsertSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = SwagNum(arr[j], arr[j+1])
		}
	}
	return arr
}
