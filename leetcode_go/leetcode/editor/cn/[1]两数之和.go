
//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums{
		m[num] = index+1
	}
	array := make([]int, 2)
	for i,num := range nums{
		index := m[target-num]
		if m[target-num]!=0&&i!=index-1{
			array[0] = i
			array[1] = index-1
			break
		}
	}
	return array
}
//leetcode submit region end(Prohibit modification and deletion)
