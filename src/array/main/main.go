package main

import "fmt"

func ChangeArrayValue(arr [5]int) [5]int {
	arr[3] = 3000
	return arr
}

func main(){

	nums:=[...]int{10, 20, 30, 40, 50}

	for i:=0;i<len(nums);i++{
		fmt.Println(nums[i])
	}

	t:=[...]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	var t_2 [5] float64
	t_2 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	for i, v:=range t{
		fmt.Println(i, v)
	}

	for i, v:=range t_2{
		fmt.Println("t_2", i, v)
	}

	fmt.Println()
	temp := []float64{24.0, 25.9, 27.8, 26.9, 26.2}
	temp_2 := []float64{24.0, 25.9, 27.8, 26.9, 26.2}

	temp_2 = temp
	temp_2[1] = 10

	for _, v := range temp{
		fmt.Println("temp : ", v)
	}

	t_2 = t
	t_2[1] = 10.0
	//복사를 해도 이 경우에 값이 안바뀐다...할당을 다르게 하고 메모리 주소는 복사 안하는게 팩트
	for i, v:=range t{
		fmt.Println(i, v)
	}

	arr:=[...]int{1, 2, 3, 4, 5}
	arr = ChangeArrayValue(arr)
	for _, v:=range arr{
		fmt.Println(v)
	}

}
