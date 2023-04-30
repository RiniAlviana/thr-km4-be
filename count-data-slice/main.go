package main

import "fmt"

func howManyElements(data []any) int {
	if data != nil {
		a := len(data)
		return a
	}
	return -1
}

func main() {
	fmt.Println(howManyElements([]any{1, 2, 3, 4, 5}))
	fmt.Println(howManyElements([]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(howManyElements([]any{1, 1, 1, 5, 5, 5}))
	fmt.Println(howManyElements([]any{"Edo", "Budi", "Joko", "Tono"}))
	fmt.Println(howManyElements([]any{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}))
	fmt.Println(howManyElements([]any{true, false, true, false, true, false}))
}
