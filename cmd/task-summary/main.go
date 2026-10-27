package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "new":
		return "Статус new."
	case "in_progress":
		return "Статус in_progress."
	case "done":
		return "Статус done."
	default:
		return status + " - Неизвестный статус."
	}
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	count := 0
	for _, status := range statuses {
		if status == target {
			count++
		}
	}
	return count
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))

	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n\nСтатистика:\n", countStatus(statuses, "done"))

	counts := make(map[string]int)

	for _, status := range statuses {
		counts[status]++
	}

	order := []string{"new", "in_progress", "done"}
	for _, status := range order {
		fmt.Printf("%s: %d\n", status, counts[status])
	}
}
