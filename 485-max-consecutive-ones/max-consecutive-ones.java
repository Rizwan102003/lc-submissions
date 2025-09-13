class Solution {
    public int findMaxConsecutiveOnes(int[] nums) {
        int ans=0,c=0;
        int n=nums.length;
        for(int i=0;i<n;i++){
            if(nums[i]==1)c++;
            else{ 
                ans=Math.max(ans,c);
                c=0;
            }
        }
        return Math.max(ans,c);
    }
}