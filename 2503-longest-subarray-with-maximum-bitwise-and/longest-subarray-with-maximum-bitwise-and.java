class Solution {
    public int longestSubarray(int[] nums) {
        int maxNum=nums[0];
        int maxLength = 0;
        int currentLength = 0;
        for (int num : nums) 
            maxNum = Math.max(maxNum, num);
        for (int num : nums) {
            if (num == maxNum)
                currentLength++;
            else {
                maxLength = Math.max(maxLength, currentLength);
                currentLength = 0;
            }
        }
        return Math.max(maxLength,currentLength);
    }
}