package main

import "fmt"

func main() {
	const name_course = "Бэкенд-разработка на языке Go"
	name_student := "Алексей"
	var name_group string = "Учебная-01"
	complete_works := 1
	var interes = true

	fmt.Println("Участник:", name_student)
	fmt.Printf("Группа: %s\nКурс: %s\nЗавершено работ: %d\n", name_group, name_course, complete_works)
	fmt.Println("Интерес к бэкенду:", interes)
}
