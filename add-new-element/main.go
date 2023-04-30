package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		data = append([]int{newData}, data[:(len(data))]...)
		return data
	}
	if position == "down" {
		data = append(data, newData)
		return data
	}
	return nil
}

func main() {
	fmt.Println(AddElement([]int{1, 2, 3, 4, 5}, 6, "up"))
	fmt.Println(AddElement([]int{1, 2, 3, 4, 5}, 6, "down"))
	fmt.Println(AddElement([]int{}, 5, "down"))
	fmt.Println(AddElement([]int{}, 5, "up"))              //slice kosong
	fmt.Println(AddElement([]int{1, 2, 3, 4}, 5, "rigth")) //position bukan up/down
	fmt.Println(AddElement([]int{1, 2, 3, 4}, 5, "left"))  //position bukan up/down
}
