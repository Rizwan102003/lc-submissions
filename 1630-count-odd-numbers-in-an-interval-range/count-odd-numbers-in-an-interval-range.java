class Solution {
    public int countOdds(int low, int high) {
        low=low/2;
        high=(high+1)/2;
        return high-low;
    }
}