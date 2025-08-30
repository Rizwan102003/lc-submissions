class Solution {
    public int gcdOfOddEvenSums(int n) {
        int se=n*(n+1);
        int so=n*n;
        int ans=gcd(se,so);
        return ans;
    }
    int gcd(int a,int b){
        if(b==0) return a;
        return gcd(b,a%b);
    }
}