// https://youtu.be/97oTiOCuxho?si=gGynbohBqyOZde0L
class Solution {
    public int longestOnes(int[] nums, int k) {
        int left = 0,right = 0,n=nums.length; 
        while (right < n) { 
            if (nums[right] == 0) k--;
            if (k < 0) {
                if (nums[left] == 0) k++;
                left++; 
            }
            right++;
        }
        return right-left;
    }
}