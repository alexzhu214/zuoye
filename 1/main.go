package main

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	result := DeleteSlice(s, 3)
	for _, v := range result {
		println(v)
	}

}

func DeleteSlice(a []int, elem int) []int {
	for i := 0; i < len(a); i++ {
		if a[i] == elem {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}