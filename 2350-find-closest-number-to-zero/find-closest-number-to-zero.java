class Solution {
    public int findClosestNumber(int[] nums) {
        int m=Integer.MAX_VALUE;
        int n=nums.length;
        for(int i=n-1;i>=0;i--){
            int a=nums[i];
            if (Math.abs(a) < Math.abs(m))
                m=nums[i];
            else if (Math.abs(a) == Math.abs(m) && a>m)
                m=nums[i];
        }
        return m;
    }
}
/*
correct answer below :

import java.util.HashSet;
import java.util.Set;
public class Solution {
    public int findClosestNumber(int[] nums) {
        int closest = nums[0];
        for (int x : nums) {
            if (Math.abs(x) < Math.abs(closest)) {
                closest = x;
            }
        }
        if (closest < 0 && contains(nums, Math.abs(closest))) {
            return Math.abs(closest);
        } else {
            return closest;
        }
    }
    private boolean contains(int[] nums, int value) {
        for (int num : nums) {
            if (num == value) {
                return true;
            }
        }
        return false;
    }
}
*/