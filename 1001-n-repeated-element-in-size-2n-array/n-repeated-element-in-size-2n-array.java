class Solution {
    public int repeatedNTimes(int[] nums) {
        Set<Integer> store = new HashSet<>(nums.length / 2 + 1);
        for (int i=0; ;i++)
            if (!store.add(nums[i]))
                return nums[i];
    }
}