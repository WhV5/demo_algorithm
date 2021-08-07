//============================================================
// @Description:
// @Author: henry
// @Date: 2021/8/3 23:06
// @File: BubbleSort_Test
//
//============================================================

package course01

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBubble(m *testing.T) {
	arr := make([]int, 10)

	rand.Seed(time.Now().Unix())

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	time.Sleep(time.Second)

	fmt.Println(BubbleSort(arr))
}
