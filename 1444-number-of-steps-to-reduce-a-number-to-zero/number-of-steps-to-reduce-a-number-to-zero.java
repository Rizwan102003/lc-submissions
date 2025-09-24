public class Solution {
    public int numberOfSteps(int n) {
        if (n == 0) return 0;
        int ones = Integer.bitCount(n);
        int highestBitIndex = 31 - Integer.numberOfLeadingZeros(n);
        return ones + highestBitIndex;
    }
    public int numberOfStepsLong(long n) {
        if (n == 0L) return 0;
        int ones = Long.bitCount(n);
        int highestBitIndex = 63 - Long.numberOfLeadingZeros(n);
        return ones + highestBitIndex;
    }
}
