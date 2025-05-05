//self
class Solution {
    public int missingNumber(int[] nums) {
        int sum=0,a=0,n=nums.length;
        sum=(n*(n+1))/2;
        for(int i=0;i<n;i++)
            a+=nums[i];
        return sum-a;
    }
}