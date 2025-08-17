class Solution {
    public int search(int[] nums, int target) {
        int h=nums.length-1,l=0;
        while(l<=h){
            int mid=l+(h-l)/2;
            if(nums[mid]==target){
                return mid;
            }
            else if(nums[mid]>target)
                h=--mid;
            else if(nums[mid]<target)
                l=mid+1;
        }
        return -1;
    }
}