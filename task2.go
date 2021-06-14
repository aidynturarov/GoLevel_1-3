package main

import "fmt"

func primeNumbers(numSlice []int) []int {
	// Перебираем по методу Решето Эратосфена
	for i := range numSlice {
		for j := range numSlice {
			if numSlice[j] > 0 && numSlice[j]%numSlice[i] == 0 && j != i {
				numSlice[j] = numSlice[j] * -1
			}
		}
	}
	// Удаляем с слайса отрицательные числа
	for r := 0; r < len(numSlice); r++ {
		if numSlice[r] < 0 {
			numSlice = append(numSlice[:r], numSlice[r+1:]...)
			r--
		}
	}

	return numSlice
}

func main() {
	var n int
	numSlice := []int{2}

	fmt.Print("Все простые числа от 2 до ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}

	if n < 3 {
		fmt.Println("Число N должно быть больше 2")
	} else {
		var i int
		for i = 3; i <= n; i++ {
			numSlice = append(numSlice, i)
		}
		fmt.Println(primeNumbers(numSlice))
	}
}
