class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> indexMap = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int current = nums[i]; 
            int complement = target - current;
            if (indexMap.containsKey(complement))
                return new int[] {indexMap.get(complement), i};
            indexMap.put(current, i);
        }
        return new int[] {0,0};
    }
}