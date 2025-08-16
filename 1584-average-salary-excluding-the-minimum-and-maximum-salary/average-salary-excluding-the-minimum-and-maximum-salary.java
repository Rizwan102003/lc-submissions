class Solution {
    public double average(int[] salary) {
        int n=salary.length;
        int s=0,min=Integer.MAX_VALUE,max=Integer.MIN_VALUE;
        for(int i:salary){
            s+=i;
            if(min>i) min=i;
            if(max<i) max=i;
        }
        s=s-max-min;
        return (double)s/(n-2);
    }
}