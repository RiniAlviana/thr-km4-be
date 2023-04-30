package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	if arr != nil && position == "right" {
		return arr[:len(arr)-1]
	}
	if arr != nil && position == "left" {
		return arr[1:]
	}

	return nil
}

func main() {
	fmt.Println(removeLeftRight([]any{1, 2, 3, 4, 5}, "left"))
	fmt.Println(removeLeftRight([]any{1, 2, 3, 4, 5}, "right"))
	fmt.Println(removeLeftRight([]any{"Edo", "Budi", "Joko", "Tono"}, "left"))
	fmt.Println(removeLeftRight([]any{"Edo", "Budi", "Joko", "Tono"}, "right"))
}
