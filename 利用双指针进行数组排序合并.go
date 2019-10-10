package leetcode

import "fmt"

//两个递增序列A，B 需要合并成递增序列C
func MergeSortIntArr(arr1, arr2 []int) (res []int) {
	//定义两个指针
	arr1Pot, arr2Pot, arr1Len, arr2Len := 0, 0, len(arr1), len(arr2)
	// 两个数组不同的索引界限
	for arr1Pot < arr1Len && arr2Pot < arr2Len {
		if arr1[arr1Pot] < arr2[arr2Pot] {
			res = append(res, arr1[arr1Pot])
			arr1Pot++
		} else {
			res = append(res, arr2[arr2Pot])
			arr2Pot++
		}
	}
	//处理最后尾巴部位剩余的元素
	for arr2Pot < arr2Len {
		res = append(res, arr2[arr2Pot])
		arr2Pot++
	}
	for arr1Pot < arr1Len {
		res = append(res, arr1[arr1Pot])
		arr1Pot++
	}
	return
}

func main() {
	// 提供两个有序的数列
	arr1 := []int{1, 2, 6, 8, 9}
	arr2 := []int{1, 3, 5, 7}
	fmt.Println(MergeSortIntArr(arr1, arr2))
}
