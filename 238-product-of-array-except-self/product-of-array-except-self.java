class Solution {
    public int[] productExceptSelf(int[] nums) {
        int n=nums.length;
        int ans[] = new int[n];
        Arrays.fill(ans,1);
        int mult=1,z=0,pos=0;
        for(int i=0;i<n;i++){
            if(nums[i]!=0) mult*=nums[i];
            else {
                z++;
                pos=i;
            }
        }
        if(z==0){
            for(int i=0;i<n;i++)
                ans[i]=mult/nums[i];
        }
        else if(z==1){
            Arrays.fill(ans,0);
            ans[pos]=mult;
        }
        else 
            Arrays.fill(ans,0);
        return ans;
    }
}
/*
3 cases :
1) no zeros - multiply all and divide each
2) more than one zero - all will be zeros
3) one zero - zero pos=multiply,rest all 0
*/