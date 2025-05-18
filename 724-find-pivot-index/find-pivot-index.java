// https://algo.monster/liteproblems/724
class Solution {
    public int pivotIndex(int[] nums) {
        int sum=0,n=nums.length;
        for(int i: nums)
            sum+=i;
        int l=0,r=0;
        for(int i=0;i<n;i++){
            r=sum-l-nums[i];
            if(l==r) return i;
            l+=nums[i]; // add l after check to omit current value
        }
        return -1;
    }
}