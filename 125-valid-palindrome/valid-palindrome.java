// self
class Solution {
    public boolean isPalindrome(String s) {
        if (s.isEmpty()) 
        	return true;
        int start = 0;
        int last = s.length() - 1;
        while(start <= last) {
        	if (!Character.isLetterOrDigit(s.charAt(start)))
        		start++;
        	else if(!Character.isLetterOrDigit(s.charAt(last)))
        		last--;
        	else {
        		if (Character.toLowerCase(s.charAt(start)) != Character.toLowerCase(s.charAt(last)))
        			return false;
        		start++;
        		last--;
        	}
        }
        return true;
    }
}