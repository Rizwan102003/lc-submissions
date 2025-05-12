// https://algo.monster/liteproblems/643
class Solution {
    public double findMaxAverage(int[] nums, int k) {
        int currentSum = 0;
        for (int i = 0; i < k; ++i)
            currentSum += nums[i];
        int maxSum = currentSum;
        for (int i = k; i < nums.length; ++i) {
            currentSum += (nums[i] - nums[i - k]);
            maxSum = Math.max(maxSum, currentSum);
        }
        return maxSum * 1.0 / k;
    }
}