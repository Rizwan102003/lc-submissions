class Solution {
    public int maximum69Number (int num) {
        String c=String.valueOf(num);
        String ans=c.replaceFirst("6","9");
        return Integer.valueOf(ans);
    }
}