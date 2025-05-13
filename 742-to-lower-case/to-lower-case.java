class Solution {
    public String toLowerCase(String s) {
        String ans = "";
        int a,temp;
        char c;
        for(int i=0;i<s.length();i++){
            c=s.charAt(i);
            temp = (int) c;
            if(65<=temp && temp<=90){
                a=temp-65+97;
                ans+= (char) a;
            }
            else ans+=c;
        }
        return ans;
    }
}