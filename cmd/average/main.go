package main

import "fmt"

func main() {
	time1 := 20
	time2 := 25
	time3 := 35

	sum := time1 + time2 + time3
	medium := float64(sum) / float64(3)

	fmt.Printf("Сумма: %d\nСреднее арифметическое: %.2f\n", sum, medium)
}
