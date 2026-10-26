package main

import "fmt"

func main() {
	var totalMinutes int
	const min_in_hour = 60
	fmt.Print("Введите количество минут: ")

	_, err := fmt.Scan(&totalMinutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}

	fmt.Printf("Полных часов: %d\nОставшихся минут: %d\n", totalMinutes/min_in_hour, totalMinutes%min_in_hour)
}
