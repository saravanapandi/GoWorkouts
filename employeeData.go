package main

import (
	"fmt"
)

func main() {
	var id int
	var name string
	// var empId int
	employeeData := make(map[int]string)
	for {
		n, _ := fmt.Scanf("%d", &id)
		fmt.Println(n)
		// empId, _ = strconv.Atoi(id)
		if id == -999 {
			break
		}
		fmt.Scanf(" %s", &name)
		employeeData[id] = name
	}
	fmt.Println(employeeData)
}
