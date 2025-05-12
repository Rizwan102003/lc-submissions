// https://algo.monster/liteproblems/1679
class Solution {
    public int maxOperations(int[] nums, int k) {
        int idx=0;
        for(int i :nums)
            if(i<k)
                nums[idx++]=i;
        Arrays.sort(nums,0,idx);
        int l=0,r=idx-1,sum=0,ans=0;
        while(l<r){
            sum=nums[l]+nums[r];
            if(sum==k){
                l++;
                r--;
                ans++;
            }
            else if(sum<k) l++;
            else r--;
        }
        return ans;
    }
}