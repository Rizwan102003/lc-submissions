/* 
Instead of starting at bottom,think from which steps we can reach n. we can take only 1 or 2 steps at a time so we can reach n from n-1 and n-2 only;
how to reach n-1 and n-2 ? again same logic.
*/
class Solution {
    public int climbStairs(int n) {
        int store[]=new int[n+1];
        store[0]=1;
        store[1]=1;
        for(int i=2;i<=n;i++)
            store[i]=store[i-1]+store[i-2];
        return store[n];
    }
}