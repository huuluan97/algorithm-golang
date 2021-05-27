package some_sorts

func bubbleSort(arr []int) []int {
	isSwapValue := true
	for isSwapValue {
		isSwapValue = false
		for i := 0; i < len(arr) - 1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				isSwapValue = true
			}
		}
	}

	return arr
}