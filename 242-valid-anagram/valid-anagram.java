class Solution {
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) return false;
        int[] store = new int[26];
        for (int i = 0; i < s.length(); i++) {
            store[s.charAt(i) - 'a']++;
            store[t.charAt(i) - 'a']--;
        }
        for (int n : store) if (n != 0) return false;
        return true;
    }
}
/*
keep an integer array store of size 26 to count characters (assuming only lowercase 'a' to 'z').
store[i] keeps track of the difference in counts between s and t for the character i + 'a'.
For each character in both strings:
Increment count for the character from s.
Decrement count for the character from t.

If both strings have the same characters in same quantity, the net result in store[] will be all zeros.

If any character count is not zero, return false (not anagrams).
Otherwise, return true.
*/