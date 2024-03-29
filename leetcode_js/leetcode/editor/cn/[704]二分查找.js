
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
	let left = 0;
	let right = nums.length;
	while(left<right){
		let mid = left + Math.floor((right-left)/2);
		if(nums[mid]===target){
			right = mid;
		}else if(nums[mid]<target){
			left = mid + 1;
		}else if(nums[mid]>target){
			right = mid;
		}
	}
	return nums[right]===target?right:-1;
};
//leetcode submit region end(Prohibit modification and deletion)
