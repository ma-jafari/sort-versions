package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//arr := []string{"10", "2.1", "9.0.0.1", "2.0.0", "1", "1.1", "2.0.0.2.1"}
	//arr := []string{"0.0.1", "10", "2.1", "9.0..0.1", "2.0.0", "1", "0", "1.1....", "2.0.0.2.1"}
	arr := []string{"0.0.1", "10", "2.1", "9.0.0.1", "2.0.0", "1", "0", "1.1", "2.0.0.2.1"}
	for i := 0; i <= len(arr); i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			var valA []string
			var valB []string
			valA = strings.Split(arr[i], ".")
			valB = strings.Split(arr[j], ".")

			var searchLen int
			if len(valA) > len(valB) {
				searchLen = len(valA) - 1
			} else {
				searchLen = len(valB) - 1
			}
			for k := 0; k <= searchLen; k++ {
				if k <= len(valA)-1 && k <= len(valB)-1 {
					num1, _ := strconv.Atoi(valA[k])
					num2, _ := strconv.Atoi(valB[k])
					if num1 > num2 {
						temp := arr[i]
						arr[i] = arr[j]
						arr[j] = temp
						break
					} else if num2 > num1 {
						break
					}
				} else if len(valB) > len(valA) {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
					break
				}
			}
		}
	}
	fmt.Println(arr)
}
