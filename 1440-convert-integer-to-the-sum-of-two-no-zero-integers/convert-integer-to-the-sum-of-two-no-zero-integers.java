class Solution {
    public int[] getNoZeroIntegers(int n) {
        int i=1;
        while (true){
            String check = i + "" + (n-i);
            if (!check.contains("0"))
                return new int[]{i,n-i};
            i++;
        }
    }
}