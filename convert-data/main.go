package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	var dataUser User
	u := strings.Split(data, ",")
	dataUser.Name = u[0]
	dataUser.Age, _ = strconv.Atoi(u[1])
	dataUser.Address = u[2]
	return dataUser
}

func main() {
	data := "Edo,20,Jakarta"
	dataUser := ConvertData(data)
	fmt.Printf("input:\"%s\"\n", data)
	fmt.Printf("%+v\n", dataUser)
}
