package main

import "fmt"

//Функция "Проверка последовательности"
func Solution(A []int) int {
	switch len(A) {
	case 0:
		return 0
	default:
		var m map[int]bool = make(map[int]bool)
		for _, val := range A {
			m[val] = !m[val]
		}
		for i := 1; i <= len(A); i++ {
			if !m[i] {
				return 0
			}
		}
		return 1
	}
}

func main() {
	A := []int{}
	fmt.Println(Solution(A))
}
