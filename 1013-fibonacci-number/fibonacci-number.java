class Solution {
    public int fib(int n) {
        int previousNumber = 0;
        int currentNumber = 1;
        while (n-- > 0) {
            int nextNumber = previousNumber + currentNumber;
            previousNumber = currentNumber;
            currentNumber = nextNumber;
        }
        return previousNumber;
    }
}