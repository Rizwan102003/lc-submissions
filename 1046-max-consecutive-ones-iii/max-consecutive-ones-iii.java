// https://youtu.be/97oTiOCuxho?si=gGynbohBqyOZde0L
class Solution {
    public int longestOnes(int[] nums, int k) {
        int i = 0,j = 0,n=nums.length; 
        while (i < n) { 
            if (nums[i] == 0) k--;
            if (k<0 && nums[j++] == 0) k++;
            i++;
        }
        return i-j;
    }
}