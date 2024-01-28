
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
	const map = {};
	const array = [];
	for(let i=0;i<nums.length;i++) {
		map[nums[i]] = i;
	}
	for(let i=0;i<nums.length;i++) {
		if(map[target-nums[i]]&&map[target-nums[i]]!==i) {
			array.push(i);
			array.push(map[target-nums[i]]);
			break;
		}
	}
	return array;
};
//leetcode submit region end(Prohibit modification and deletion)
