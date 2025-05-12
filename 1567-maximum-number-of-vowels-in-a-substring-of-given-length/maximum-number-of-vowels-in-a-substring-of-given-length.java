// https://algo.monster/liteproblems/1456
class Solution {
    public int maxVowels(String s, int k) {
        int c = 0,n=s.length();
        for (int i = 0; i < k;i++)
            if (isVowel(s.charAt(i)))
                c++;
        int max = c;
        for (int i = k; i <n;i++) {
            if (isVowel(s.charAt(i)))
                c++;
            if (isVowel(s.charAt(i - k)))
                c--;
            max = Math.max(max,c);
        }
        return max;
    }
    boolean isVowel(char character) {
        return character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u';
    }
}