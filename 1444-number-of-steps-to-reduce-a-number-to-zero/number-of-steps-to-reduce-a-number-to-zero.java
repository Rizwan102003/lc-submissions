class Solution {
    public int numberOfSteps(int num) {
        if (num == 0) return 0;
        int ones = Integer.bitCount(num);
        int highestBitIndex = 31 - Integer.numberOfLeadingZeros(num);
        return ones + highestBitIndex;
    }
}