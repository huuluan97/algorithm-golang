package some_sorts

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	q := partition(arr)
	QuickSort(arr[:q])
	QuickSort(arr[q+1:])
	return arr
}

func partition(arr []int) int {
	r := len(arr) - 1
	x := arr[r] // last element of array
	i := -1
	for j := 0; j <= r; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
