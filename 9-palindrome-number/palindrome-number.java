// self
class Solution {
    public boolean isPalindrome(int x) {
        int rev = 0, r = 0, c = x;
        while (x > 0) {
            rev = rev * 10 + (x % 10);
            x /= 10;
        }
        if (rev == c) return true; else return false;
    }
}