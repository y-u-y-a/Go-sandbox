package libs

// 大文字で始める
func Average(list []int) int {
	total := 0
	for _, i := range list {
		total += i
	}
	return int(total / len(list))
}
