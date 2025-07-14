class Solution {
    public int arrayPairSum(int[] nums) {
        Arrays.sort(nums);
        int a=0;
        int n=nums.length;
        for(int i=0;i<n;i=i+2)
            a+=nums[i];
        return a;
    }
}