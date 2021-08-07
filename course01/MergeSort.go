//============================================================
// @Description: 归并排序
// @Author: henry
// @Date: 2021/8/4 07:18
// @File: MergeSort
//
//============================================================

package course01

func MergeSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	sortProcess(arr, 0, len(arr)-1)
	return arr
}

func sortProcess(arr []int, L, R int) {
	if L == R {
		return
	}
	var mid int
	mid = (L + R) / 2
	sortProcess(arr, L, mid)
	sortProcess(arr, mid+1, R)

	merge(arr, L, mid, R)
}

func merge(arr []int, L, mid, R int) {
	help := make([]int, R-L+1)
	var i, p1, p2 int
	p1 = L
	p2 = mid + 1

	for {
		if p1 > mid || p2 > R {
			break
		}
		for i := 0; i < len(help); i++ {
			if p1 > mid || p2 > R {
				break
			}
			if arr[p1] > arr[p2] {
				help[i] = arr[p1]
			} else {
				help[i] = arr[p2]
			}
			p1++
			p2++
		}
	}

	if p1 > mid {
		for ; p2 < R; p2++ {
			i++
			help[i] = arr[p2]
		}
	} else {
		for ; p1 < mid; p1++ {
			i++
			help[i] = arr[p1]
		}
	}

	for i := 0; i < len(help); i++ {
		arr[L+i] = help[i]
	}
}
