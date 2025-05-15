//self
class Solution {
    public int largestAltitude(int[] gain) {
        int n=gain.length;
        int c=0,max=0;
        for(int i=0;i<n;i++){
            c+=gain[i];
            max=Math.max(c,max);
        }
        return max;
    }
}