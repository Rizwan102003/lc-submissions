class Solution {
    public long zeroFilledSubarray(int[] nums) {
        long ans=0,streak=0;
        for(int i:nums){
            if(i==0)
                streak++;
            else 
                streak = 0;
            ans+=streak;
        }
        return ans;
    }
}