class Solution {
    public int scoreOfString(String s) {
        int n=s.length();
        int ans=0;
        for(int i=1;i<n;i++)
            ans+=Math.abs(s.charAt(i)-s.charAt(i-1));
        return ans;
    }
}