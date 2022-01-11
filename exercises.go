package exercises


func Add(numbers ...int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}

	return sum
}
