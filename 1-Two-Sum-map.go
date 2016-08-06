func twoSum(nums []int, target int) []int {
	var temp int
	var flag bool
	record := make(map[int]int)
	result := make([]int, 2)
	for i, v := range nums {
		temp, flag = record[target-v]//Retrieve `record[target-v]` first to avoid duplicate
		record[v] = i
		if flag == true {
			result[0] = temp
			result[1] = i
			return result
		}
	}
	return result
}