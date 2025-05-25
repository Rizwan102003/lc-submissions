class Solution {
    public int longestSubarray(int[] nums) {
        int i = 0,j = 0,k=1,n=nums.length; 
        while (i < n){
            if (nums[i++] == 0) k--;
            if (k<0 && nums[j++] == 0) k++;
        }
        return i-j-1;
    }
}