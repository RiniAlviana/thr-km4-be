package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	for a, _ := range dataMap {
		if a == key {
			delete(dataMap, key)
		}
	}
	return dataMap
}

func main() {
	dataMap := map[string]any{
		"satu":  "ichi",
		"dua":   "ni",
		"tiga":  "san",
		"empat": "yon",
		"lima":  "go",
	}

	key := "tiga"

	unreleted := removeUnrelated(dataMap, key)

	fmt.Println(unreleted)
}
