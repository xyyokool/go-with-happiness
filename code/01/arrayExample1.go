package main

import (
	"fmt"
)

func main() {
	arr1 := [4]int{2,3,4,5}
	printArray4(&arr1) //999 3 4 5
	fmt.Println(arr1) // [2 3 4 5]
}

// 1. 数组声明
func array1() {
	// 定义方式1, 需要声明数组长度, 默认值是0
	var arr1 [3]int
	// 定义方式2, 带值+长度
	arr2 := [3]int{1,2,3}
	// 定义方式3, 带值+不限制长度
	arr3 := [...]int{1,2,3,4,5,56,6,7,8}
	// 定义方式4, 3行2列矩阵
	var arr4 [3][2]int
	
	fmt.Println(arr1,arr2,arr3)  // [0 0 0] [1 2 3] [1 2 3 4 5 56 6 7 8]
	fmt.Println(arr4) // [[0 0] [0 0] [0 0]]  --> 3*2
}

// 2. 数组遍历
func array2() {
	arr := [...]int{1,2,3,4,5,56,6,7,8}
	// 遍历方式1, 常规for
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// 遍历方式2, 使用range简化
	sum := 0
	// 第一位是index, 用不到, 所以我们用_占位, 只要value
	for _, value := range arr {
		sum += value
	}
	fmt.Println(sum)
}

// 3. GO语言数组特色
func array3() {
	// var arr1 = [3]int{1,2,3}
	// var arr2 = [4]int{1,2,3,4}
	// var arr3 = [...]int{1,2,3}
	// var arr4 = [...]int{1,2,3,5,6}
	// fmt.Println(arr1 == arr2) // 报错, 无法比较
	// fmt.Println(arr3 == arr4) // 报错, 无法比较
	// fmt.Println(arr1 == arr3) // 可以比较
}

func printArray4(arr *[4]int) {
	arr[0] = 999
	for _, v := range arr {
		fmt.Println(v)
	}
}

