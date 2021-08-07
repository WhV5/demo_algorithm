//============================================================
// @Description:
// @Author: henry
// @Date: 2021/8/4 03:24
// @File: InsertSort_test
//
//============================================================

package course01

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInsert(m *testing.T) {
	arr := make([]int, 10)

	rand.Seed(time.Now().Unix())

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	time.Sleep(time.Second)

	fmt.Println(InsertSort(arr))
}
