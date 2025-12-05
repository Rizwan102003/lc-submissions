class Solution {
    public int countPartitions(int[] nums) {
        int count=0;
        int max=0;
        for(int i=0;i<nums.length;i++){
            max+=nums[i];
        }
        for(int i=0;i<nums.length;i++){
            if(max%2!=0) return 0;
        }
        return nums.length-1;
    }
}