package main

import (
	"fmt"
)

// 冒泡排序的精髓在于每次遍历所有的元素。每次遍历都有一个最值出现。
// 例如下列，从小到大排序。 1 6 比较。不变。6 2 然后 6 后移，变成 126 。继续调换 6 4 。变成 1246
// 实际上下面的实现有点多余。在于最后一个元素已经有序了。第二次实际上可以不比对这两个
// 原数组 1 6 2 4 9 0 5 3 7 8
// 内层循环
// 第一次1 6比较，不变 1 6 2 4 9 0 5 3 7 8
// 第二次6 2, 变成 1 2 6 4 9 0 5 3 7 8
// 第三次6 4，变成 1 2 4 6 9 0 5 3 7 8
// 第四次6 9，不变
// 第五次9 0，变成 1 2 4 6 0 9 5 3 7 8
// 第六次9 5，变成 1 2 4 6 0 5 9 3 7 8
// 第七次9 3, 变成 1 2 4 6 0 5 3 9 7 8
// 第八次9 7，变成 1 2 4 6 0 5 3 7 9 8
// 第八次9 8，变成 1 2 4 6 0 5 3 7 8 9
// 这个时候内层循环结束，9冒泡了，成为 最大的最值。
// 继续上面的过程。
// 只要有一个交换动作。那么势必会走到 swappwed = true。会令死循环继续。直到不发生一次交换。这个时候也就成了整体有序。
// 这里面有一个弊端，就是8 9 还会再比较一下。。毫无意义的比较。
// 这里面又有一个好处。如果原本的数据有序度比较高。那么可以在狠少的计算中，排好序
func bubbleSort(arrayzor []int) {

	swapped := true
	for swapped {
		count := 0
		swapped = false
		for i := 0; i < len(arrayzor)-1; i++ {
			if arrayzor[i+1] < arrayzor[i] {
				Swap(arrayzor, i, i+1)
				swapped = true
			}
			count += 1
		}
	}
	fmt.Println("count:", count)
}

func Swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

func main() {

	arrayzor := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", arrayzor)
	bubbleSort(arrayzor)
	fmt.Println("Sorted array: ", arrayzor)
}

// 常规版。这个是比较常规的冒泡算法。
//
func bubbleSort(arrayzor []int) {
	count := 0
	for j := 0; j < len(arrayzor); j++ {
		for i := 0; i < len(arrayzor)-1-j; i++ {
			if arrayzor[i+1] < arrayzor[i] {
				Swap(arrayzor, i, i+1)
			}
			count += 1
		}
	}
	fmt.Println("count:", count)
}

func Swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

func main() {

	arrayzor := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", arrayzor)
	bubbleSort(arrayzor)
	fmt.Println("Sorted array: ", arrayzor)
}

// 我们通过以上两个加入count 这个计数变量之后，得到第一个通过控制位跳出需要要54次循环。第二个都是固定45个。。
