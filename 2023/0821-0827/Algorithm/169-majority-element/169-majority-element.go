package main

func majorityElement(nums []int) int {
	var count, candidate int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	nums := []int{3, 2, 3}
	println(majorityElement(nums))
}