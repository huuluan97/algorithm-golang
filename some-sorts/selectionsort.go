package some_sorts

func swap(a int, b int) (int,int) {
	a, b = b, a
	return a, b
}

func SelectionSort(arr []int) []int {
	for i := 0; i <= len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[min] {
				min = j
			}
			temp := arr[i]
			arr[i] = arr[min]
			arr[min] = temp
		}
	}
	return arr
}
