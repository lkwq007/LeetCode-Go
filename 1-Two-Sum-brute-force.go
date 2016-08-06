func twoSum(nums []int, target int) []int {
	result:=make([]int,2)
	var i,j int
	n:=len(nums)
	for i=0;i<n-1;i++ {
		for j=i+1;j<n;j++ {
			if nums[i]+nums[j]==target {
				result[0]=i
				result[1]=j
				return result
			}
		}
	}
	return result
}