// https://www.youtube.com/watch?v=9kmUaXrjizQ&t=2s
class Solution {
    public boolean isValid(String s) {
        if (s.length()%2!=0) return false;
        Stack<Character> store=new Stack();
        for(char c: s.toCharArray()){
            if(c=='(' || c=='{' || c=='[')
                store.push(c);
            else if(c==')' && !store.isEmpty() && store.peek()=='(')
                store.pop();
            else if(c=='}' && !store.isEmpty() && store.peek()=='{')
                store.pop();
            else if(c==']' && !store.isEmpty() && store.peek()=='[')
                store.pop();
            else return false;
        }
        return store.isEmpty();
    }
}